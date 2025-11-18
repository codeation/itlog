package impresslink

import (
	"github.com/codeation/impress/event"
	gtk "github.com/codeation/itlog/gtk4"
)

const (
	destroyEventID      = 1
	textClipboardFormat = 1
)

func (a *Application) onDelete() {
	a.callbacks.EventGeneral(destroyEventID)
}

func (a *Application) onKeyPress(ev *gtk.GdkEventKey) {
	a.callbacks.EventKeyboard(gtk.GdkKey(ev))
}

func (a *Application) onButtonPress(ev *gtk.GdkEventButton) {
	a.callbacks.EventButton(gtk.GdkButton(ev))
	if ev.Type == event.ButtonActionRelease {
		a.layout.GrabFocus()
	}
}

func (a *Application) onMotionNotify(ev *gtk.GdkEventMotion) {
	a.callbacks.EventMotion(gtk.GdkMotion(ev))
}

func (a *Application) onScroll(ev *gtk.GdkEventScroll) {
	a.callbacks.EventScroll(gtk.GdkScroll(ev))
}

var prevWidth, prevHeight, prevInnerWidth, prevInnerHeight int

func (a *Application) onSizeAllocate(innerWidth int, innerHeight int) {
	width, height := gtk.GdkConfigure(a.top)
	if width == prevWidth && height == prevHeight &&
		innerWidth == prevInnerWidth && innerHeight == prevInnerHeight {
		return
	}
	prevWidth, prevHeight, prevInnerWidth, prevInnerHeight = width, height, innerWidth, innerHeight
	a.callbacks.EventConfigure(width, height, innerWidth, innerHeight)
}

func (a *Application) onClipboardText(text string) {
	a.callbacks.EventClipboard(textClipboardFormat, []byte(text))
}
