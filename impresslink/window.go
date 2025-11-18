package impresslink

import (
	"image"
	"image/color"
	"runtime/cgo"
	"sync"

	"github.com/codeation/impress/driver"
	gtk "github.com/codeation/itlog/gtk4"
)

func colors(c color.Color) (uint16, uint16, uint16, uint16) {
	r, g, b, a := c.RGBA()
	return uint16(r), uint16(g), uint16(b), uint16(a)
}

func onDraw(cr *gtk.Cairo, h *cgo.Handle) {
	w := h.Value().(*Window)
	w.mutex.RLock()
	painters := w.painters
	w.mutex.RUnlock()
	for _, paint := range painters {
		paint.Paint(cr)
	}
}

type painter interface{ Paint(cr *gtk.Cairo) }
type destroyer interface{ Destroy() }

type Window struct {
	a          *Application
	drawing    *gtk.Drawing
	size       image.Point
	background color.Color
	cgoHandle  *cgo.Handle
	painters   []painter
	paintqueue []painter
	destroyers []destroyer
	mutex      sync.RWMutex
}

func (f *Frame) NewWindow(rect image.Rectangle, background color.Color) driver.Painter {
	w := &Window{
		a:          f.a,
		size:       rect.Size(),
		background: background,
	}
	h := cgo.NewHandle(w)
	w.cgoHandle = &h
	w.a.commands <- func() {
		w.drawing = f.gtkFrame.NewDrawingArea()
		w.drawing.Move(rect.Min.X, rect.Min.Y)
		w.drawing.Size(rect.Dx(), rect.Dy())
		w.drawing.Show()
		w.drawing.SignalDraw(w.cgoHandle)
	}
	return w
}

func (w *Window) Drop() {
	w.a.commands <- func() {
		w.drawing.Destroy()
		w.cgoHandle.Delete()
		for _, paint := range w.destroyers {
			paint.Destroy()
		}
		w.destroyers = nil
	}
}

func (w *Window) Size(rect image.Rectangle) {
	w.size = rect.Size()
	w.a.commands <- func() {
		w.drawing.Move(rect.Min.X, rect.Min.Y)
		w.drawing.Size(rect.Dx(), rect.Dy())
	}
}

func (w *Window) Raise() {
	w.a.commands <- func() {
		w.drawing.Raise()
	}
}

func (w *Window) Clear() {
	for _, paint := range w.destroyers {
		paint.Destroy()
	}
	w.destroyers = nil
	w.paintqueue = nil
	w.Fill(image.Rectangle{Max: w.size}, w.background)
}

func (w *Window) Show() {
	w.mutex.Lock()
	w.painters = w.paintqueue
	w.mutex.Unlock()
	w.a.commands <- func() { w.drawing.QueueDraw() }
}

func (w *Window) Fill(rect image.Rectangle, foreground color.Color) {
	r, g, b, a := colors(foreground)
	paint := gtk.NewCairoFillPaint(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy(), r, g, b, a)
	w.paintqueue = append(w.paintqueue, paint)
}

func (w *Window) Line(from image.Point, to image.Point, foreground color.Color) {
	r, g, b, a := colors(foreground)
	paint := gtk.NewCairoLinePaint(from.X, from.Y, to.X, to.Y, r, g, b, a)
	w.paintqueue = append(w.paintqueue, paint)
}

func (w *Window) Image(rect image.Rectangle, img driver.Imager) {
	paint := gtk.NewCairoImagePaint(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy(), getImageBitmap(img))
	w.paintqueue = append(w.paintqueue, paint)
}

func (w *Window) Text(text string, font driver.Fonter, from image.Point, foreground color.Color) {
	r, g, b, a := colors(foreground)
	paint := gtk.NewCairoTextPaint(from.X, from.Y, r, g, b, a, getFontSelection(font), text)
	w.paintqueue = append(w.paintqueue, paint)
	w.destroyers = append(w.destroyers, paint)
}
