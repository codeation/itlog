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
	activateFn      = func() {}
	shutdownFn      = func() {}
	menuActionFn    = func(h *cgo.Handle) {}
	drawFn          = func(cr *Cairo, h *cgo.Handle) {}
	deleteFn        = func() {}
	sizeAllocateFn  = func(width int, height int) {}
	keyPressFn      = func(event *GdkEventKey) {}
	buttonPressFn   = func(event *GdkEventButton) {}
	buttonReleaseFn = func(event *GdkEventButton) {}
	motionNotifyFn  = func(event *GdkEventMotion) {}
	scrollFn        = func(event *GdkEventScroll) {}
	clipboardFn     = func(text string) {}
	weakRefFn       = func(data uintptr) {}
)

//export appActivate
func appActivate(application *C.GApplication, data C.gpointer) { activateFn() }

// SignalActivate connects a callback for "activate" event
func SignalActivate(callback func()) { activateFn = callback }

//export appShutdown
func appShutdown(application *C.GApplication, data C.gpointer) { shutdownFn() }

// SignalShutdown connects a callback for "shutdown" event
func SignalShutdown(callback func()) { shutdownFn = callback }

//export windowDraw
func windowDraw(widget *C.GtkWidget, cr *C.cairo_t, data C.gpointer) {
	drawFn((*Cairo)(cr), (*cgo.Handle)(data))
}

// SignalDrawCallback connects a callback for "draw" event
func SignalDrawCallback(callback func(cr *Cairo, h *cgo.Handle)) { drawFn = callback }

//export widgetDelete
func widgetDelete(self *C.GtkWindow, data C.gpointer) { deleteFn() }

// SignalDelete connects a callback for "delete-event" event
func SignalDelete(callback func()) { deleteFn = callback }

var sizeWidget *C.GtkWidget

//export widgetIdle
func widgetIdle(data C.gpointer) {
	width := int(C.gtk_widget_get_width(sizeWidget))
	height := int(C.gtk_widget_get_height(sizeWidget))
	if width == 0 && height == 0 {
		return
	}
	sizeAllocateFn(width, height)
}

// SignalSizeAllocate connects a callback for "size-allocate" event
func SignalSizeAllocate(callback func(width int, height int)) { sizeAllocateFn = callback }

//export widgetKeyPress
func widgetKeyPress(self *C.GtkEventControllerKey, keyval C.guint, keycode C.guint, state C.GdkModifierType, data C.gpointer) {
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
func SignalKeyPress(callback func(event *GdkEventKey)) { keyPressFn = callback }

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
func widgetButtonPress(self *C.GtkGestureClick, n_press C.gint, x C.gdouble, y C.gdouble, data C.gpointer) {
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
func SignalButtonPress(callback func(event *GdkEventButton)) { buttonPressFn = callback }

//export widgetButtonRelease
func widgetButtonRelease(self *C.GtkGestureClick, n_press C.gint, x C.gdouble, y C.gdouble, data C.gpointer) {
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
func SignalButtonRelease(callback func(event *GdkEventButton)) { buttonReleaseFn = callback }

//export widgetMotionNotify
func widgetMotionNotify(self *C.GtkEventControllerMotion, x C.gdouble, y C.gdouble, data C.gpointer) {
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
		State: C.gdk_event_get_modifier_state(event),
	}
	motionNotifyFn(motionEvent)
}

// SignalMotionNotify connects a callback for "motion_notify_event" event
func SignalMotionNotify(callback func(event *GdkEventMotion)) { motionNotifyFn = callback }

//export widgetScroll
func widgetScroll(self *C.GtkEventControllerScroll, dx C.gdouble, dy C.gdouble, data C.gpointer) {
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
func SignalScroll(callback func(event *GdkEventScroll)) { scrollFn = callback }

func gSignalConnect(instance *C.GObject, signal string, handler C.GCallback, data C.gpointer) C.gulong {
	name := C.CString(signal)
	defer C.free(unsafe.Pointer(name))
	return C.GSignalConnect(instance, name, handler, data)
}

// AppSignalConnect connects application level signals
func (app *Application) AppSignalConnect() {
	gSignalConnect(app.GObject(), "activate", C.GCallback(C.app_activate), nil)
	gSignalConnect(app.GObject(), "shutdown", C.GCallback(C.app_shutdown), nil)
}

// AppSignalDisconnect disconnects application level signals
func (app *Application) AppSignalDisconnect() {
	C.GSignalHandlersDisconnectByFunc(app.GObject(), C.GCallback(C.app_activate), nil)
	C.GSignalHandlersDisconnectByFunc(app.GObject(), C.GCallback(C.app_shutdown), nil)
}

// TopSignalConnect connects application window level signals
func (top *TopWindow) TopSignalConnect() {
	gSignalConnect(top.widget.GObject(), "close-request", C.GCallback(C.close_request), nil)
}

// TopSignalDisconnect disconnects application window level signals
func (top *TopWindow) TopSignalDisconnect() {
	C.GSignalHandlersDisconnectByFunc(top.widget.GObject(), C.GCallback(C.close_request), nil)
}

// LayoutSignalConnect connects layout level signals
func (layout *Layout) LayoutSignalConnect() {
	sizeWidget = layout.scrolled
	gSignalConnect(layout.parent.Widget().GObject(), "notify::default-width", C.GCallback(C.size_notify), nil)
	gSignalConnect(layout.parent.Widget().GObject(), "notify::default-height", C.GCallback(C.size_notify), nil)
	gSignalConnect(C.adjustmentToGObject(layout.adjustment), "changed", C.GCallback(C.adjustment_notify), nil)
	gSignalConnect(C.adjustmentToGObject(layout.adjustment), "value-changed", C.GCallback(C.adjustment_notify), nil)
	C.size_notify_init()

	layout.keyEventController = C.gtk_event_controller_key_new()
	gSignalConnect(C.controllerToGObject(layout.keyEventController), "key-pressed", C.GCallback(C.key_pressed), nil)
	C.gtk_widget_add_controller(layout.scrolled, layout.keyEventController)

	layout.getstureConroller = C.gtk_gesture_click_new()
	C.gtk_gesture_single_set_button(C.gestureToGestureSingle(layout.getstureConroller), 0)
	C.gtk_gesture_single_set_touch_only(C.gestureToGestureSingle(layout.getstureConroller), 0) // false
	gSignalConnect(C.gestureToGObject(layout.getstureConroller), "pressed", C.GCallback(C.button_pressed), nil)
	gSignalConnect(C.gestureToGObject(layout.getstureConroller), "released", C.GCallback(C.button_released), nil)
	gSignalConnect(C.gestureToGObject(layout.getstureConroller), "unpaired-release", C.GCallback(C.button_released), nil)
	C.gtk_widget_add_controller(layout.widget.gtkWidget, C.gestureToEventController(layout.getstureConroller))

	layout.motionEventController = C.gtk_event_controller_motion_new()
	gSignalConnect(C.controllerToGObject(layout.motionEventController), "motion", C.GCallback(C.motion_notify), nil)
	C.gtk_widget_add_controller(layout.widget.gtkWidget, layout.motionEventController)

	layout.scrollEventController = C.gtk_event_controller_scroll_new(C.GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES)
	gSignalConnect(C.controllerToGObject(layout.scrollEventController), "scroll", C.GCallback(C.scroll_notify), nil)
	C.gtk_widget_add_controller(layout.widget.gtkWidget, layout.scrollEventController)
}

// LayoutSignalDisconnect disconnects layout level signals
func (layout *Layout) LayoutSignalDisconnect() {
	C.GSignalHandlersDisconnectByFunc(layout.parent.Widget().GObject(), C.GCallback(C.size_notify), nil)
	C.GSignalHandlersDisconnectByFunc(C.adjustmentToGObject(layout.adjustment), C.GCallback(C.adjustment_notify), nil)

	C.GSignalHandlersDisconnectByFunc(C.controllerToGObject(layout.keyEventController), C.GCallback(C.key_pressed), nil)
	C.gtk_widget_remove_controller(layout.scrolled, layout.keyEventController)

	C.GSignalHandlersDisconnectByFunc(C.gestureToGObject(layout.getstureConroller), C.GCallback(C.button_pressed), nil)
	C.GSignalHandlersDisconnectByFunc(C.gestureToGObject(layout.getstureConroller), C.GCallback(C.button_released), nil)
	C.gtk_widget_remove_controller(layout.widget.gtkWidget, C.gestureToEventController(layout.getstureConroller))

	C.GSignalHandlersDisconnectByFunc(C.controllerToGObject(layout.motionEventController), C.GCallback(C.motion_notify), nil)
	C.gtk_widget_remove_controller(layout.widget.gtkWidget, layout.motionEventController)

	C.GSignalHandlersDisconnectByFunc(C.controllerToGObject(layout.scrollEventController), C.GCallback(C.scroll_notify), nil)
	C.gtk_widget_remove_controller(layout.widget.gtkWidget, layout.scrollEventController)
}

// SignalDraw connects a callback parameters for "draw" event
func (drawing *Drawing) SignalDraw(h *cgo.Handle) {
	C.gtk_drawing_area_set_draw_func(C.widgetToGtkDrawingArea(drawing.Widget().GtkWidget()), C.GCallback(C.window_draw), C.gpointer(h), nil)
}

//export menuItemActivate
func menuItemActivate(action *C.GSimpleAction, parameter *C.GVariant, data C.gpointer) {
	menuActionFn((*cgo.Handle)(data))
}

// SignalMenuItemActivateCallback connects a callback for "activate" event
func SignalMenuItemActivateCallback(callback func(h *cgo.Handle)) { menuActionFn = callback }

// SignalMenuItemActivate connects a callback  parameters for "activate" event
func (action *MenuAction) SignalMenuItemActivate(h *cgo.Handle) {
	gSignalConnect(action.GObject(), "activate", C.GCallback(C.menu_item_activate), C.gpointer(h))
}

//export clipboardTextReceived
func clipboardTextReceived(source_object *C.GObject, res *C.GAsyncResult, data C.gpointer) {
	text := C.gdk_clipboard_read_text_finish(C.objectToGdkClipboard(source_object), res, nil)
	clipboardFn(C.GoString(text))
}

// RequestClipboardText connects a callback for clipboard receiving
func RequestClipboardText(callback func(text string)) {
	clipboardFn = callback
	clipboard := C.gdk_display_get_primary_clipboard(C.gdk_display_get_default())
	C.gdk_clipboard_read_text_async(clipboard, nil, C.GCallback(C.clipboard_text_received), nil)
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
