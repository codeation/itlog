package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"

type Menu struct {
	menu *C.GMenu
}

type MenuNode struct {
	*Menu
	label *C.char
}

type MenuItem struct {
	item   *C.GMenuItem
	label  *C.char
	action *C.char
}

type MenuAction struct {
	simpleAction   *C.GSimpleAction
	action         *C.char
	alias          *C.char
	signalHandlers []C.gulong
}

func (app *Application) NewMenu() *Menu {
	node := C.g_menu_new()
	C.gtk_application_set_menubar(app.a, C.menuToGMenuModel(node))
	C.g_object_unref(C.menuToGPointer(node))
	return &Menu{
		menu: node,
	}
}

func (menu *Menu) Destroy() { /* TODO */ }

func (parent *Menu) NewNode(label string) *MenuNode {
	node := C.g_menu_new()
	cLabel := C.CString(label)
	C.g_menu_append_submenu(parent.menu, cLabel, C.menuToGMenuModel(node))
	C.g_object_unref(C.menuToGPointer(node))
	return &MenuNode{
		Menu: &Menu{
			menu: node,
		},
		label: cLabel,
	}
}

func (node *MenuNode) NewItem(label string, action string) *MenuItem {
	cLabel := C.CString(label)
	cAction := C.CString(action)
	item := C.g_menu_item_new(cLabel, cAction)
	C.g_menu_append_item(node.menu, item)
	C.g_object_unref(C.menuItemToGPointer(item))
	return &MenuItem{
		item:   item,
		label:  cLabel,
		action: cAction,
	}
}

func (item *MenuItem) Destroy() { /* TODO */ }

func (app *Application) NewMenuAction(action string) *MenuAction {
	cAction := C.CString(action)
	alias := C.CString(action[4:])
	simpleAction := C.g_simple_action_new(alias, nil)
	C.g_action_map_add_action(app.GActionMap(), C.simpleToGAction(simpleAction))
	return &MenuAction{
		simpleAction: simpleAction,
		action:       cAction,
		alias:        alias,
	}
}

func (action *MenuAction) GObject() *C.GObject { return C.simpleToGObject(action.simpleAction) }

func (action *MenuAction) Destroy() { /* TODO */ }
