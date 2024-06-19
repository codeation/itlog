package gtk4

// #cgo pkg-config: gtk4
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"

// GtkAllocation ia a GtkAllocation wrapper
type GtkAllocation C.GtkAllocation

// GdkEventKey is a keyboard event parameters
type GdkEventKey struct {
	Event   *C.GdkEvent
	Keyval  C.guint
	Keycode C.guint
	State   C.GdkModifierType
}

// GdkEventButton is a button event parameters
type GdkEventButton struct {
	Event *C.GdkEvent
	Type  int
	X, Y  C.double
}

// GdkEventMotion is a motion event parameters
type GdkEventMotion struct {
	Event *C.GdkEvent
	X, Y  C.double
}

// GdkEventScroll is a scroll event parameters
type GdkEventScroll struct {
	Event     *C.GdkEvent
	Direction int
	DX, DY    C.double
}

// GdkConfigure returns application width, height and inner rectange width, height
func GdkConfigure(allocation *GtkAllocation, top *TopWindow) (int, int, int, int) {
	//layoutOffsetX, layoutOffsetY = allocation.x, allocation.y
	var width, height C.gint
	C.gtk_window_get_default_size(C.widgetToGtkWindow(top.Widget().GtkWidget()), &width, &height)
	return int(width), int(height), int(allocation.width), int(allocation.height)
}

// GdkKey returns key rune, shift, control, alt, meta statuses and key name
func GdkKey(event *GdkEventKey) (rune, bool, bool, bool, bool, string) {
	r := rune(C.gdk_keyval_to_unicode(event.Keyval))
	name := C.GoString(C.gdk_keyval_name(event.Keyval))
	shift := event.State&C.GDK_SHIFT_MASK != 0
	control := event.State&C.GDK_CONTROL_MASK != 0
	alt := event.State&C.GDK_ALT_MASK != 0
	meta := event.State&C.GDK_META_MASK != 0
	return r, shift, control, alt, meta, name
}

// GdkButton returns mouse button type, button id, x, y of mouse event
func GdkButton(event *GdkEventButton) (int, int, int, int) {
	button := int(C.gdk_button_event_get_button(event.Event))
	return event.Type, button, int(event.X), int(event.Y)
}

// GdkMotion returns x, y of mouse motion event, shift, control, alt, meta statuses
func GdkMotion(event *GdkEventMotion) (int, int, bool, bool, bool, bool) {
	return int(event.X), int(event.Y), false, false, false, false
}

// GdkScroll returns scroll direction, deltaX, deltaY of scroll event
func GdkScroll(event *GdkEventScroll) (int, int, int) {
	return event.Direction, int(event.DX), int(event.DY)
}
