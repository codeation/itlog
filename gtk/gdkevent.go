package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"

// GtkAllocation ia a GtkAllocation wrapper
type GtkAllocation C.GtkAllocation

// GdkEventKey is a GdkEventKey wrapper
type GdkEventKey C.GdkEventKey

// GdkEventButton is a GdkEventButton wrapper
type GdkEventButton C.GdkEventButton

// GdkEventMotion is a GdkEventMotion wrapper
type GdkEventMotion C.GdkEventMotion

// GdkEventScroll is a GdkEventScroll wrapper
type GdkEventScroll C.GdkEventScroll

var layoutOffsetX, layoutOffsetY C.int

// GdkConfigure returns application width, height and inner rectange width, height
func GdkConfigure(allocation *GtkAllocation, top *TopWindow) (int, int, int, int) {
	layoutOffsetX, layoutOffsetY = allocation.x, allocation.y
	var width, height C.gint
	C.gtk_window_get_size(C.widgetToGtkWindow(top.Widget().GtkWidget()), &width, &height)
	return int(width), int(height), int(allocation.width), int(allocation.height)
}

// GdkKey returns key rune, shift, control, alt, meta statuses and key name
func GdkKey(event *GdkEventKey) (rune, bool, bool, bool, bool, string) {
	r := rune(C.gdk_keyval_to_unicode(event.keyval))
	shift := event.state&C.GDK_SHIFT_MASK != 0
	control := event.state&C.GDK_CONTROL_MASK != 0
	alt := event.state&C.GDK_MOD1_MASK != 0
	meta := event.state&C.GDK_META_MASK != 0 // TODO GDK_SUPER_MASK https://docs.gtk.org/gdk3/flags.ModifierType.html
	name := C.GoString(C.gdk_keyval_name(event.keyval))
	return r, shift, control, alt, meta, name
}

// GdkButton returns mouse button type, button id, x, y of mouse event
func GdkButton(event *GdkEventButton) (int, int, int, int) {
	eventType := int(event._type)
	button := int(event.button)
	x := int(event.x) - int(layoutOffsetX)
	y := int(event.y) - int(layoutOffsetY)
	return eventType, button, x, y
}

// GdkMotion returns x, y of mouse motion event, shift, control, alt, meta statuses
func GdkMotion(event *GdkEventMotion) (int, int, bool, bool, bool, bool) {
	x := int(event.x) - int(layoutOffsetX)
	y := int(event.y) - int(layoutOffsetY)
	shift := event.state&C.GDK_SHIFT_MASK != 0
	control := event.state&C.GDK_CONTROL_MASK != 0
	alt := event.state&C.GDK_MOD1_MASK != 0
	meta := event.state&C.GDK_META_MASK != 0
	return x, y, shift, control, alt, meta
}

// GdkScroll returns scroll direction, deltaX, deltaY of scroll event
func GdkScroll(event *GdkEventScroll) (int, int, int) {
	direction := int(event.direction)
	deltaX := int(event.delta_x)
	deltaY := int(event.delta_y)
	return direction, deltaX, deltaY
}
