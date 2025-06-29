package uiapi

import (
	"log"

	"github.com/codeation/itlog/gtk"
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

type frame struct {
	gtkFrame gtkFramer
	id       int
	parentID int
	x, y     int
}

func (f *frame) isTop() bool { return f.id == 1 && f.parentID == 0 }

func (u *uiAPI) FrameNew(frameID int, parentFrameID int, x, y, width, height int) {
	f := &frame{
		id:       frameID,
		parentID: parentFrameID,
		x:        x,
		y:        y,
	}
	u.frames[frameID] = f

	if f.isTop() {
		layout := u.top.NewLayout()
		layout.LayoutSignalConnect()
		f.gtkFrame = layout
		u.top.ShowAll()
	} else {
		parent, ok := u.frames[parentFrameID]
		if !ok {
			log.Printf("FrameNew: parent frame not found: %d", parentFrameID)
			return
		}
		f.gtkFrame = parent.gtkFrame.NewFixed()
		f.gtkFrame.Move(f.x, f.y)
		f.gtkFrame.Size(width, height)
		f.gtkFrame.Show()
	}
}

func (u *uiAPI) FrameDrop(frameID int) {
	f, ok := u.frames[frameID]
	if !ok {
		log.Printf("FrameDrop: frame not found: %d", frameID)
		return
	}

	f.gtkFrame.Destroy()
	delete(u.frames, frameID)
}

func (u *uiAPI) FrameSize(frameID int, x, y, width, height int) {
	f, ok := u.frames[frameID]
	if !ok {
		log.Printf("FrameSize: frame not found: %d", frameID)
		return
	}

	if f.isTop() {
		log.Printf("FrameSize: top frame cannot be resized")
		return
	}

	f.x = x
	f.y = y
	f.gtkFrame.Move(f.x, f.y)
	f.gtkFrame.Size(width, height)
}

func (u *uiAPI) FrameRaise(frameID int) {
	f, ok := u.frames[frameID]
	if !ok {
		log.Printf("FrameRaise: frame not found: %d", frameID)
		return
	}

	if f.isTop() {
		log.Printf("FrameRaise: top frame cannot be raised")
		return
	}

	f.gtkFrame.Raise()
	f.gtkFrame.Move(f.x, f.y)
}
