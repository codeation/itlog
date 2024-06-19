package uiapi4

import (
	"log"
	"runtime/cgo"

	"github.com/codeation/itlog/gtk4"
)

type painter interface {
	Paint(cr *gtk4.Cairo)
}

type destroyer interface {
	Destroy()
}

type window struct {
	gtkDrawing *gtk4.Drawing
	cgoHandle  *cgo.Handle
	id         int
	frameID    int
	x, y       int
	painters   []painter
	destroyers []destroyer
}

func onWindowDraw(cr *gtk4.Cairo, h *cgo.Handle) {
	w := h.Value().(*window)
	for _, paint := range w.painters {
		paint.Paint(cr)
	}
}

func (u *uiAPI) WindowNew(windowID int, frameID int, x, y, width, height int) {
	f, ok := u.frames[frameID]
	if !ok {
		log.Printf("WindowNew: frame not found: %d", frameID)
		return
	}

	w := &window{
		gtkDrawing: f.gtkFrame.NewDrawingArea(),
		id:         windowID,
		frameID:    frameID,
		x:          x,
		y:          y,
	}

	h := cgo.NewHandle(w)
	w.cgoHandle = &h
	u.windows[windowID] = w

	w.gtkDrawing.Move(w.x, w.y)
	w.gtkDrawing.Size(width, height)
	w.gtkDrawing.Show()
	w.gtkDrawing.SignalDraw(w.cgoHandle)

	//gtk4.NotifyWeakRef(w.gtkDrawing.Widget().GObject(), u.onWeakRef, uintptr(windowID))
}

func (u *uiAPI) WindowDrop(windowID int) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowDrop: window not found: %d", windowID)
		return
	}
	w.gtkDrawing.Destroy()
	w.cgoHandle.Delete()
	for _, paint := range w.destroyers {
		paint.Destroy()
	}
	w.destroyers = nil
	delete(u.windows, windowID)
}

func (u *uiAPI) WindowRaise(windowID int) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowRaise: window not found: %d", windowID)
		return
	}
	w.gtkDrawing.Raise()
	w.gtkDrawing.Move(w.x, w.y)
}

func (u *uiAPI) WindowClear(windowID int) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowClear: window not found: %d", windowID)
		return
	}
	for _, paint := range w.destroyers {
		paint.Destroy()
	}
	w.destroyers = nil
	w.painters = nil
}

func (u *uiAPI) WindowShow(windowID int) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowShow: window not found: %d", windowID)
		return
	}
	w.gtkDrawing.QueueDraw()
}

func (u *uiAPI) WindowSize(windowID int, x, y, width, height int) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowSize: window not found: %d", windowID)
		return
	}
	w.x = x
	w.y = y
	w.gtkDrawing.Move(w.x, w.y)
	w.gtkDrawing.Size(width, height)
}

func (u *uiAPI) WindowFill(windowID int, x, y, width, height int, r, g, b, a uint16) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowFill: window not found: %d", windowID)
		return
	}
	paint := gtk4.NewCairoFillPaint(x, y, width, height, r, g, b, a)
	w.painters = append(w.painters, paint)
}

func (u *uiAPI) WindowLine(windowID int, x0, y0, x1, y1 int, r, g, b, a uint16) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowLine: window not found: %d", windowID)
		return
	}
	paint := gtk4.NewCairoLinePaint(x0, y0, x1, y1, r, g, b, a)
	w.painters = append(w.painters, paint)
}

func (u *uiAPI) WindowText(windowID int, x, y int, r, g, b, a uint16, fontID int, text string) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowText: window not found: %d", windowID)
		return
	}

	f, ok := u.fonts[fontID]
	if !ok {
		log.Printf("WindowText: font not found: %d", fontID)
		return
	}

	paint := gtk4.NewCairoTextPaint(x, y, r, g, b, a, f.selection, text)
	w.painters = append(w.painters, paint)
	w.destroyers = append(w.destroyers, paint)
}

func (u *uiAPI) WindowImage(windowID int, x, y, width, height int, imageID int) {
	w, ok := u.windows[windowID]
	if !ok {
		log.Printf("WindowImage: window not found: %d", windowID)
		return
	}

	i, ok := u.images[imageID]
	if !ok {
		log.Printf("WindowImage: image not found: %d", imageID)
		return
	}

	paint := gtk4.NewCairoImagePaint(x, y, width, height, i.bitmap)
	w.painters = append(w.painters, paint)
}
