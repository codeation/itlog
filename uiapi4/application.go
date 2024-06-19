// Package implements an internal mechanism to communicate with an impress terminal.
package uiapi4

import (
	"github.com/codeation/impress/joint/iface"
	"github.com/codeation/itlog/gtk4"
)

type uiAPI struct {
	app       *gtk4.Application
	top       *gtk4.TopWindow
	menu      *gtk4.Menu
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
		app:       gtk4.NewApplication(),
		frames:    map[int]*frame{},
		windows:   map[int]*window{},
		fonts:     map[int]*font{},
		images:    map[int]*image{},
		menuNodes: map[int]*menuNode{},
		menuItems: map[int]*menuItem{},
		callbacks: callbacks,
	}

	gtk4.SignalMenuItemActivateCallback(u.onItemActivate)
	gtk4.SignalDrawCallback(onWindowDraw)
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
	text := gtk4.RequestClipboardText(u.top)
	u.callbacks.EventClipboard(typeID, []byte(text))
}

func (u *uiAPI) ClipboardPut(typeID int, data []byte) {
	gtk4.SetClipboardText(u.top, string(data))
}

func (u *uiAPI) Sync() {}

func (u *uiAPI) onStartup() {
	u.menu = u.app.NewMenu()
}

func (u *uiAPI) onActivate() {
	u.top = u.app.NewTopWindow()
	u.top.Widget().SignalDelete(u.onDelete)
	u.top.Widget().SignalKeyPress(u.onKeyPress)
}

func (u *uiAPI) onShutdown() {
	u.top.Destroy()
	u.app.Quit()
}
