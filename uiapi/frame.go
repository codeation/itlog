package uiapi

import (
	"log"

	"github.com/codeation/itlog/gtk"
)

type frame struct {
	gtkFrame *gtk.FrameWidget
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
	if f.isTop() {
		f.gtkFrame = u.top.NewLayout()
	} else {
		parent, ok := u.frames[parentFrameID]
		if !ok {
			log.Printf("FrameNew: parent frame not found: %d", parentFrameID)
			return
		}
		f.gtkFrame = parent.gtkFrame.NewFixed()
	}

	u.frames[frameID] = f

	if f.isTop() {
		u.top.ShowAll()
		f.gtkFrame.SignalSizeAllocate(u.onSizeAllocate)
	} else {
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
