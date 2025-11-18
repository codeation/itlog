package uiapi4

import (
	"github.com/codeation/impress/event"
	gtk "github.com/codeation/itlog/gtk4"
)

const (
	destroyEventID      = 1
	textClipboardFormat = 1
)

func (u *uiAPI) onDelete() {
	u.callbacks.EventGeneral(destroyEventID)
}

func (u *uiAPI) onKeyPress(ev *gtk.GdkEventKey) {
	u.callbacks.EventKeyboard(gtk.GdkKey(ev))
}

func (u *uiAPI) onButtonPress(ev *gtk.GdkEventButton) {
	u.callbacks.EventButton(gtk.GdkButton(ev))
	if ev.Type == event.ButtonActionRelease {
		u.layout.GrabFocus()
	}
}

func (u *uiAPI) onMotionNotify(ev *gtk.GdkEventMotion) {
	u.callbacks.EventMotion(gtk.GdkMotion(ev))
}

func (u *uiAPI) onScroll(event *gtk.GdkEventScroll) {
	u.callbacks.EventScroll(gtk.GdkScroll(event))
}

var prevWidth, prevHeight, prevInnerWidth, prevInnerHeight int

func (u *uiAPI) onSizeAllocate(innerWidth int, innerHeight int) {
	width, height := gtk.GdkConfigure(u.top)
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
