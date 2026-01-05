// Package implements an internal mechanism to communicate with an impress terminal.
package uiapi4

import (
	"github.com/codeation/impress/joint/iface"
	gtk "github.com/codeation/itlog/gtk4"
)

type uiAPI struct {
	app         *gtk.Application
	top         *gtk.TopWindow
	layout      *gtk.Layout
	menu        *gtk.Menu
	frames      map[int]*frame
	windows     map[int]*window
	fonts       map[int]*font
	fontMetrics map[int]*font
	images      map[int]*image
	menuNodes   map[int]*menuNode
	menuItems   map[int]*menuItem
	callbacks   iface.CallbackSet
	exitCount   int
}

func New(callbacks iface.CallbackSet) *uiAPI {
	u := &uiAPI{
		app:         gtk.NewApplication(),
		frames:      map[int]*frame{},
		windows:     map[int]*window{},
		fonts:       map[int]*font{},
		fontMetrics: map[int]*font{},
		images:      map[int]*image{},
		menuNodes:   map[int]*menuNode{},
		menuItems:   map[int]*menuItem{},
		callbacks:   callbacks,
	}
	gtk.SignalActivate(u.onActivate)
	gtk.SignalShutdown(u.onShutdown)
	gtk.SignalMenuItemActivateCallback(u.onItemActivate)
	gtk.SignalDrawCallback(onWindowDraw)
	gtk.SignalDelete(u.onDelete)
	gtk.SignalSizeAllocate(u.onSizeAllocate)
	gtk.SignalKeyPress(u.onKeyPress)
	gtk.SignalButtonPress(u.onButtonPress)
	gtk.SignalButtonRelease(u.onButtonPress)
	gtk.SignalMotionNotify(u.onMotionNotify)
	gtk.SignalScroll(u.onScroll)
	u.app.AppSignalConnect()
	return u
}

func (u *uiAPI) Run() {
	u.app.Run()
}

func (u *uiAPI) ApplicationTitle(title string) {
	u.app.SetName(title)
}

func (u *uiAPI) ApplicationSize(x, y, width, height int) {
	u.top.Size(x, y, width, height)
}

func (u *uiAPI) ApplicationExit() {
	u.exitCount++
	switch u.exitCount {
	case 1:
		u.top.TopSignalDisconnect()
	case 2:
		u.top.Destroy()
		u.app.Quit()
	}
}

func (u *uiAPI) ApplicationVersion() string { return logAPIVersion }

func (u *uiAPI) ClipboardGet(typeID int) {
	gtk.RequestClipboardText(u.top, u.onClipboardText)
}

func (u *uiAPI) ClipboardPut(typeID int, data []byte) {
	gtk.SetClipboardText(u.top, string(data))
}

func (u *uiAPI) Sync() {}

func (u *uiAPI) onActivate() {
	u.menu = u.app.NewMenu()
	u.top = u.app.NewTopWindow()
	u.top.TopSignalConnect()

	//gtk.NotifyWeakRef(u.top.Widget().GObject(), u.onWeakRef, uintptr(9001))
}

func (u *uiAPI) onShutdown() {
	u.app.AppSignalDisconnect()
}
