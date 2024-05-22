// Package implements an internal mechanism to communicate with an impress terminal.
package uiapi

import (
	"github.com/codeation/impress/joint/iface"
	"github.com/codeation/itlog/gtk"
)

type uiAPI struct {
	app       *gtk.Application
	top       *gtk.WindowWidget
	menu      *gtk.Menu
	frames    map[int]*frame
	windows   map[int]*window
	fonts     map[int]*font
	images    map[int]*image
	menuNodes map[int]*menuNode
	menuItems map[int]*menuItem
	callbacks iface.CallbackSet
	exitCount int
}

func New(callbacks iface.CallbackSet) *uiAPI {
	w := &uiAPI{
		app:       gtk.NewApplication(),
		frames:    map[int]*frame{},
		windows:   map[int]*window{},
		fonts:     map[int]*font{},
		images:    map[int]*image{},
		menuNodes: map[int]*menuNode{},
		menuItems: map[int]*menuItem{},
		callbacks: callbacks,
	}

	gtk.SignalMenuItemActivateCallback(w.onItemActivate)
	gtk.SignalDrawCallback(onWindowDraw)
	w.app.SignalActivate(w.onActivate)
	w.app.SignalShutdown(w.onShutdown)
	return w
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
	if u.exitCount > 1 {
		u.app.Quit()
	}
}

func (u *uiAPI) ApplicationVersion() string { return itlogVersion }

func (u *uiAPI) ClipboardGet(typeID int) {
	gtk.RequestClipboardText(u.onClipboardText)
}

func (u *uiAPI) ClipboardPut(typeID int, data []byte) {
	gtk.SetClipboardText(string(data))
}

func (u *uiAPI) Sync() {}

func (u *uiAPI) onActivate() {
	u.menu = u.app.NewMenu()
	u.top = u.app.NewWindow()
	u.top.SignalDelete(u.onDelete)
	u.top.SignalKeyPress(u.onKeyPress)
	u.top.SignalButtonPress(u.onButtonPress)
	u.top.SignalButtonRelease(u.onButtonPress)
	u.top.SignalMotionNotify(u.onMotionNotify)
	u.top.SignalScroll(u.onScroll)
}

func (u *uiAPI) onShutdown() {
	u.top.Destroy()
	u.app.Quit()
}
