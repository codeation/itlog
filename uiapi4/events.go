package uiapi4

import (
	"github.com/codeation/itlog/gtk4"
)

const (
	destroyEventID      = 1
	textClipboardFormat = 1
)

func (u *uiAPI) onDelete() {
	u.callbacks.EventGeneral(destroyEventID)
}

func (u *uiAPI) onKeyPress(event *gtk4.GdkEventKey) {
	u.callbacks.EventKeyboard(gtk4.GdkKey(event))
}

func (u *uiAPI) onButtonPress(event *gtk4.GdkEventButton) {
	u.callbacks.EventButton(gtk4.GdkButton(event))
}

func (u *uiAPI) onMotionNotify(event *gtk4.GdkEventMotion) {
	u.callbacks.EventMotion(gtk4.GdkMotion(event))
}

func (u *uiAPI) onScroll(event *gtk4.GdkEventScroll) {
	u.callbacks.EventScroll(gtk4.GdkScroll(event))
}

var prevWidth, prevHeight, prevInnerWidth, prevInnerHeight int

func (u *uiAPI) onSizeAllocate(allocation *gtk4.GtkAllocation) {
	width, height, innerWidth, innerHeight := gtk4.GdkConfigure(allocation, u.top)
	if width == prevWidth && height == prevHeight &&
		innerWidth == prevInnerWidth && innerHeight == prevInnerHeight {
		return
	}
	prevWidth, prevHeight, prevInnerWidth, prevInnerHeight = width, height, innerWidth, innerHeight
	u.callbacks.EventConfigure(width, height, innerWidth, innerHeight)
}

func (u *uiAPI) onClipboardText(text string) {
	u.callbacks.EventClipboard(textClipboardFormat, []byte(text))
}

//func (u *uiAPI) onWeakRef(data uintptr) { fmt.Println("WeakRef", int(data)) }
