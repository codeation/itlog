// Package implements an internal mechanism to communicate with an impress terminal.
package uiapi

import (
	"github.com/codeation/impress/joint/iface"
	"github.com/codeation/itlog/gtk"
)

type uiAPI struct {
	app       *gtk.Application
	top       *gtk.TopWindow
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
	u := &uiAPI{
		app:       gtk.NewApplication(),
		frames:    map[int]*frame{},
		windows:   map[int]*window{},
		fonts:     map[int]*font{},
		images:    map[int]*image{},
		menuNodes: map[int]*menuNode{},
		menuItems: map[int]*menuItem{},
		callbacks: callbacks,
	}

	gtk.SignalMenuItemActivateCallback(u.onItemActivate)
	gtk.SignalDrawCallback(onWindowDraw)
	u.app.SignalStartup(u.onStartup)
	u.app.SignalActivate(u.onActivate)
	u.app.SignalShutdown(u.onShutdown)
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

func (u *uiAPI) onStartup() {
	u.menu = u.app.NewMenu()
}

func (u *uiAPI) onActivate() {
	u.top = u.app.NewTopWindow()
	u.top.Widget().SignalDelete(u.onDelete)
	u.top.Widget().SignalKeyPress(u.onKeyPress)
	u.top.Widget().SignalButtonPress(u.onButtonPress)
	u.top.Widget().SignalButtonRelease(u.onButtonPress)
	u.top.Widget().SignalMotionNotify(u.onMotionNotify)
	u.top.Widget().SignalScroll(u.onScroll)
}

func (u *uiAPI) onShutdown() {
	u.top.Destroy()
	u.app.Quit()
}
