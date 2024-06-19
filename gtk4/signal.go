package gtk4

// #cgo pkg-config: gtk4
// #include <gtk/gtk.h>
// #include "macro.h"
// #include "signalconnect.h"
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

var (
	startupFn       func()
	activateFn      func()
	shutdownFn      func()
	menuActionFn    func(h *cgo.Handle)
	drawFn          func(cr *Cairo, h *cgo.Handle)
	deleteFn        func()
	sizeAllocateFn  func(allocation *GtkAllocation)
	keyPressFn      func(event *GdkEventKey)
	buttonPressFn   func(event *GdkEventButton)
	buttonReleaseFn func(event *GdkEventButton)
	motionNotifyFn  func(event *GdkEventMotion)
	scrollFn        func(event *GdkEventScroll)
	clipboardFn     func(text string)
	weakRefFn       func(data uintptr)
)

func gSignalConnect(instance *C.GObject, signal string, handler C.GCallback, data C.gpointer) C.gulong {
	name := C.CString(signal)
	defer C.free(unsafe.Pointer(name))
	return C.GSignalConnect(instance, name, handler, data)
}

//export appStartup
func appStartup(application *C.GApplication, data C.gpointer) { startupFn() }

// SignalStartup connects a callback for "startup" event
func (app *Application) SignalStartup(callback func()) {
	startupFn = callback
	gSignalConnect(app.GObject(), "startup", C.GCallback(C.app_startup), nil)
}

//export appActivate
func appActivate(application *C.GApplication, data C.gpointer) { activateFn() }

// SignalActivate connects a callback for "activate" event
func (app *Application) SignalActivate(callback func()) {
	activateFn = callback
	gSignalConnect(app.GObject(), "activate", C.GCallback(C.app_activate), nil)
}

//export appShutdown
func appShutdown(application *C.GApplication, data C.gpointer) { shutdownFn() }

// SignalShutdown connects a callback for "shutdown" event
func (app *Application) SignalShutdown(callback func()) {
	shutdownFn = callback
	gSignalConnect(app.GObject(), "shutdown", C.GCallback(C.app_shutdown), nil)
}

//export windowDraw
func windowDraw(widget *C.GtkWidget, cr *C.cairo_t, data C.gpointer) {
	drawFn((*Cairo)(cr), (*cgo.Handle)(data))
}

// SignalDrawCallback connects a callback for "draw" event
func SignalDrawCallback(callback func(cr *Cairo, h *cgo.Handle)) {
	drawFn = callback
}

// SignalDraw connects a callback parameters for "draw" event
func (drawing *Drawing) SignalDraw(h *cgo.Handle) {
	C.gtk_drawing_area_set_draw_func(C.widgetToGtkDrawingArea(drawing.Widget().GtkWidget()), C.GCallback(C.window_draw), C.gpointer(h), nil)
}

//export widgetDelete
func widgetDelete(self *C.GtkWindow, user_data C.gpointer) { deleteFn() }

// SignalDelete connects a callback for "delete-event" event
func (widget *Widget) SignalDelete(callback func()) {
	deleteFn = callback
	gSignalConnect(widget.GObject(), "close-request", C.GCallback(C.close_request), nil)
}

var sizeWidget *C.GtkWidget

//export widgetIdle
func widgetIdle(user_data C.gpointer) {
	var allocation C.GtkAllocation
	C.gtk_widget_get_allocation(sizeWidget, &allocation)
	if allocation.width == 0 && allocation.height == 0 {
		return
	}
	sizeAllocateFn((*GtkAllocation)(&allocation))
}

// SignalSizeAllocate connects a callback for "size-allocate" event
func (layout *Layout) SignalSizeAllocate(callback func(allocation *GtkAllocation)) {
	sizeWidget = layout.scrolled
	sizeAllocateFn = callback
	gSignalConnect(layout.parent.Widget().GObject(), "notify::default-width", C.GCallback(C.size_notify), nil)
	gSignalConnect(layout.parent.Widget().GObject(), "notify::default-height", C.GCallback(C.size_notify), nil)
	gSignalConnect(C.adjustmentToGObject(layout.adjustment), "changed", C.GCallback(C.adjustment_notify), nil)
	gSignalConnect(C.adjustmentToGObject(layout.adjustment), "value-changed", C.GCallback(C.adjustment_notify), nil)
	C.size_notify_init()
	keyEventController := C.gtk_event_controller_key_new()
	gSignalConnect(C.controllerToGObject(keyEventController), "key-pressed", C.GCallback(C.key_pressed), nil)
	C.gtk_widget_add_controller(layout.scrolled, keyEventController)
}

//export widgetKeyPress
func widgetKeyPress(self *C.GtkEventControllerKey, keyval C.guint, keycode C.guint, state C.GdkModifierType, user_data C.gpointer) {
	event := C.gtk_event_controller_get_current_event(C.keyToEventController(self))
	keyEvent := &GdkEventKey{
		Event:   event,
		Keyval:  keyval,
		Keycode: keycode,
		State:   state,
	}
	keyPressFn(keyEvent)
}

// SignalKeyPress connects a callback for "key_press_event" event
func (widget *Widget) SignalKeyPress(callback func(event *GdkEventKey)) {
	keyPressFn = callback
	keyEventController := C.gtk_event_controller_key_new()
	gSignalConnect(C.controllerToGObject(keyEventController), "key-pressed", C.GCallback(C.key_pressed), nil)
	C.gtk_widget_add_controller(widget.gtkWidget, keyEventController)
}

var prevButtonTime C.guint32
var prevButtonType C.GdkEventType

func buttonType(n_press int) int {
	switch n_press {
	case 1:
		return 4
	case 2:
		return 5
	case 3:
		return 6
	default:
		return 7
	}
}

//export widgetButtonPress
func widgetButtonPress(self *C.GtkGestureClick, n_press C.gint, x C.gdouble, y C.gdouble, user_data C.gpointer) {
	event := C.gtk_event_controller_get_current_event(C.clickToEventController(self))
	if C.gdk_event_get_time(event) == prevButtonTime && C.gdk_event_get_event_type(event) == prevButtonType {
		return
	}
	prevButtonTime = C.gdk_event_get_time(event)
	prevButtonType = C.gdk_event_get_event_type(event)
	buttonEvent := &GdkEventButton{
		Event: event,
		Type:  buttonType(int(n_press)),
		X:     x,
		Y:     y,
	}
	buttonPressFn(buttonEvent)
}

// SignalButtonPress connects a callback for "button_press_event" event
func (widget *Widget) SignalButtonPress(callback func(event *GdkEventButton)) {
	buttonPressFn = callback
	getstureConroller := C.gtk_gesture_click_new()
	C.gtk_gesture_single_set_button(C.gestureToGestureSingle(getstureConroller), 0)
	C.gtk_gesture_single_set_touch_only(C.gestureToGestureSingle(getstureConroller), 0) // false
	gSignalConnect(C.gestureToGObject(getstureConroller), "pressed", C.GCallback(C.button_pressed), nil)
	C.gtk_widget_add_controller(widget.gtkWidget, C.gestureToEventController(getstureConroller))
}

//export widgetButtonRelease
func widgetButtonRelease(self *C.GtkGestureClick, n_press C.gint, x C.gdouble, y C.gdouble, user_data C.gpointer) {
	event := C.gtk_event_controller_get_current_event(C.clickToEventController(self))
	if C.gdk_event_get_time(event) == prevButtonTime && C.gdk_event_get_event_type(event) == prevButtonType {
		return
	}
	prevButtonTime = C.gdk_event_get_time(event)
	prevButtonType = C.gdk_event_get_event_type(event)
	buttonEvent := &GdkEventButton{
		Event: event,
		Type:  7,
		X:     x,
		Y:     y,
	}
	buttonReleaseFn(buttonEvent)
}

// SignalButtonRelease connects a callback for "button_release_event" event
func (widget *Widget) SignalButtonRelease(callback func(event *GdkEventButton)) {
	buttonReleaseFn = callback
	getstureConroller := C.gtk_gesture_click_new()
	C.gtk_gesture_single_set_button(C.gestureToGestureSingle(getstureConroller), 0)
	C.gtk_gesture_single_set_touch_only(C.gestureToGestureSingle(getstureConroller), 0) // false
	gSignalConnect(C.gestureToGObject(getstureConroller), "released", C.GCallback(C.button_released), nil)
	C.gtk_widget_add_controller(widget.gtkWidget, C.gestureToEventController(getstureConroller))
}

//export widgetMotionNotify
func widgetMotionNotify(self *C.GtkEventControllerMotion, x C.gdouble, y C.gdouble, user_data C.gpointer) {
	event := C.gtk_event_controller_get_current_event(C.motionToEventController(self))
	if C.gdk_event_get_time(event) == prevButtonTime && C.gdk_event_get_event_type(event) == prevButtonType {
		return
	}
	prevButtonTime = C.gdk_event_get_time(event)
	prevButtonType = C.gdk_event_get_event_type(event)
	motionEvent := &GdkEventMotion{
		Event: event,
		X:     x,
		Y:     y,
	}
	motionNotifyFn(motionEvent)
}

// SignalMotionNotify connects a callback for "motion_notify_event" event
func (widget *Widget) SignalMotionNotify(callback func(event *GdkEventMotion)) {
	motionNotifyFn = callback
	motionEventController := C.gtk_event_controller_motion_new()
	gSignalConnect(C.controllerToGObject(motionEventController), "motion", C.GCallback(C.motion_notify), nil)
	C.gtk_widget_add_controller(widget.gtkWidget, motionEventController)
}

//export widgetScroll
func widgetScroll(self *C.GtkEventControllerScroll, dx C.gdouble, dy C.gdouble, user_data C.gpointer) {
	event := C.gtk_event_controller_get_current_event(C.scrollToEventController(self))
	scrollEvent := &GdkEventScroll{
		Event:     event,
		Direction: 4,
		DX:        dx,
		DY:        dy,
	}

	scrollFn(scrollEvent)
}

// SignalScroll connects a callback for "scroll-event" event
func (widget *Widget) SignalScroll(callback func(event *GdkEventScroll)) {
	scrollFn = callback
	scrollEventController := C.gtk_event_controller_scroll_new(C.GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES)
	gSignalConnect(C.controllerToGObject(scrollEventController), "scroll", C.GCallback(C.scroll_notify), nil)
	C.gtk_widget_add_controller(widget.gtkWidget, scrollEventController)
}

//export menuItemActivate
func menuItemActivate(action *C.GSimpleAction, parameter *C.GVariant, data C.gpointer) {
	menuActionFn((*cgo.Handle)(data))
}

// SignalMenuItemActivateCallback connects a callback for "activate" event
func SignalMenuItemActivateCallback(callback func(h *cgo.Handle)) {
	menuActionFn = callback
}

// SignalMenuItemActivate connects a callback  parameters for "activate" event
func (action *MenuAction) SignalMenuItemActivate(h *cgo.Handle) {
	gSignalConnect(action.GObject(), "activate", C.GCallback(C.menu_item_activate), C.gpointer(h))
}

//export objectWeakRef
func objectWeakRef(data C.gpointer, where_the_object_was *C.GObject) {
	weakRefFn(uintptr(data))
}

// NotifyWeakRef connects a callback for weak notify
func NotifyWeakRef(object *C.GObject, callback func(data uintptr), data uintptr) {
	weakRefFn = callback
	C.g_object_weak_ref(object, C.GCallback(C.object_weak_notify), C.gpointer(data))
}
