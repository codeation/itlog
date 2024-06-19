package uiapi4

import (
	"log"
	"runtime/cgo"

	"github.com/codeation/itlog/gtk4"
)

type menuNode struct {
	id       int
	parentID int
	node     *gtk4.MenuNode
}

type menuItem struct {
	id        int
	parentID  int
	action    string
	item      *gtk4.MenuItem
	handler   *gtk4.MenuAction
	cgoHandle *cgo.Handle
}

func (u *uiAPI) onItemActivate(h *cgo.Handle) {
	item := h.Value().(*menuItem)
	u.callbacks.EventMenu(item.action)
}

func (u *uiAPI) MenuNew(menuID int, parentMenuID int, label string) {
	var node *gtk4.MenuNode
	if parentMenuID == 0 {
		node = u.menu.NewNode(label)
	} else {
		parent, ok := u.menuNodes[parentMenuID]
		if !ok {
			log.Printf("MenuNew: parent menu not found: %d", parentMenuID)
			return
		}
		node = parent.node.NewNode(label)
	}
	u.menuNodes[menuID] = &menuNode{
		id:       menuID,
		parentID: parentMenuID,
		node:     node,
	}
}

func (u *uiAPI) MenuItem(menuID int, parentMenuID int, label string, action string) {
	parent, ok := u.menuNodes[parentMenuID]
	if !ok {
		log.Printf("MenuNew: parent menu not found: %d", parentMenuID)
		return
	}
	item := &menuItem{
		id:       menuID,
		parentID: parentMenuID,
		action:   action,
		item:     parent.node.NewItem(label, action),
		handler:  u.app.NewMenuAction(action),
	}
	h := cgo.NewHandle(item)
	item.cgoHandle = &h
	u.menuItems[menuID] = item

	item.handler.SignalMenuItemActivate(item.cgoHandle)
}
