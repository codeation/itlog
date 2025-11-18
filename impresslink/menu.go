package impresslink

import (
	"runtime/cgo"

	"github.com/codeation/impress/driver"
	gtk "github.com/codeation/itlog/gtk4"
)

type Menu struct {
	a    *Application
	node *gtk.MenuNode
}

type MenuItem struct {
	action    string
	item      *gtk.MenuItem
	handler   *gtk.MenuAction
	cgoHandle *cgo.Handle
}

func (a *Application) onItemActivate(h *cgo.Handle) {
	item := h.Value().(*MenuItem)
	a.callbacks.EventMenu(item.action)
	a.layout.GrabFocus()
}

func (a *Application) NewMenu(label string) driver.Menuer {
	m := &Menu{
		a: a,
	}
	ready := make(chan struct{})
	m.a.commands <- func() {
		m.node = a.menu.NewNode(label)
		ready <- struct{}{}
	}
	<-ready
	return m
}

func (parent *Menu) NewMenu(label string) driver.Menuer {
	m := &Menu{
		a: parent.a,
	}
	ready := make(chan struct{})
	m.a.commands <- func() {
		m.node = parent.node.NewNode(label)
		ready <- struct{}{}
	}
	<-ready
	return m
}

func (m *Menu) NewItem(label string, action string) {
	i := &MenuItem{
		action: action,
	}
	h := cgo.NewHandle(i)
	i.cgoHandle = &h
	ready := make(chan struct{})
	m.a.commands <- func() {
		i.item = m.node.NewItem(label, action)
		i.handler = m.a.app.NewMenuAction(action)
		i.handler.SignalMenuItemActivate(i.cgoHandle)
		ready <- struct{}{}
	}
	<-ready
}
