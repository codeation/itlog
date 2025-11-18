package impresslink

import (
	"image"

	"github.com/codeation/impress/driver"
	gtk "github.com/codeation/itlog/gtk4"
)

type gtkFramer interface {
	NewFixed() *gtk.Fixed
	NewDrawingArea() *gtk.Drawing
	Destroy()
	Show()
	Move(x, y int)
	Size(width, height int)
	Raise()
}

type Frame struct {
	a        *Application
	gtkFrame gtkFramer
}

func (a *Application) NewFrame(rect image.Rectangle) driver.Framer {
	f := &Frame{
		a: a,
	}
	f.a.commands <- func() {
		a.layout = f.a.top.NewLayout()
		a.layout.LayoutSignalConnect()
		f.gtkFrame = a.layout
		f.a.top.ShowAll()
	}
	return f
}

func (parent *Frame) NewFrame(rect image.Rectangle) driver.Framer {
	f := &Frame{
		a: parent.a,
	}
	f.a.commands <- func() {
		f.gtkFrame = parent.gtkFrame.NewFixed()
		f.gtkFrame.Move(rect.Min.X, rect.Min.Y)
		f.gtkFrame.Size(rect.Dx(), rect.Dy())
		f.gtkFrame.Show()
	}
	return f
}

func (f *Frame) Drop() {
	f.a.commands <- func() {
		if _, ok := f.gtkFrame.(*gtk.Layout); ok {
			f.a.layout.LayoutSignalDisconnect()
			f.a.layout = nil
		}
		f.gtkFrame.Destroy()
	}
}

func (f *Frame) Size(rect image.Rectangle) {
	f.a.commands <- func() {
		f.gtkFrame.Move(rect.Min.X, rect.Min.Y)
		f.gtkFrame.Size(rect.Dx(), rect.Dy())
	}
}
func (f *Frame) Raise() {
	f.a.commands <- func() {
		f.gtkFrame.Raise()
	}
}
