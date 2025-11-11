package gtk

// #cgo pkg-config: gtk+-3.0
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
	sizeAllocateFn  = func(allocation *GtkAllocation) {}
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
func widgetDelete(widget *C.GtkWidget, event *C.GdkEvent, data C.gpointer) { deleteFn() }

// SignalDelete connects a callback for "delete-event" event
func SignalDelete(callback func()) { deleteFn = callback }

//export widgetSizeAllocate
func widgetSizeAllocate(widget *C.GtkWidget, allocation *C.GtkAllocation, data C.gpointer) {
	sizeAllocateFn((*GtkAllocation)(allocation))
}

// SignalSizeAllocate connects a callback for "size-allocate" event
func SignalSizeAllocate(callback func(allocation *GtkAllocation)) { sizeAllocateFn = callback }

//export widgetKeyPress
func widgetKeyPress(widget *C.GtkWidget, event *C.GdkEventKey, data C.gpointer) {
	keyPressFn((*GdkEventKey)(event))
}

// SignalKeyPress connects a callback for "key_press_event" event
func SignalKeyPress(callback func(event *GdkEventKey)) { keyPressFn = callback }

var prevButtonTime C.guint32
var prevButtonType C.GdkEventType

//export widgetButtonPress
func widgetButtonPress(widget *C.GtkWidget, event *C.GdkEventButton, data C.gpointer) {
	if event.time == prevButtonTime && event._type == prevButtonType {
		return
	}
	prevButtonTime = event.time
	prevButtonType = event._type
	buttonPressFn((*GdkEventButton)(event))
}

// SignalButtonPress connects a callback for "button_press_event" event
func SignalButtonPress(callback func(event *GdkEventButton)) { buttonPressFn = callback }

//export widgetButtonRelease
func widgetButtonRelease(widget *C.GtkWidget, event *C.GdkEventButton, data C.gpointer) {
	if event.time == prevButtonTime && event._type == prevButtonType {
		return
	}
	prevButtonTime = event.time
	prevButtonType = event._type
	buttonReleaseFn((*GdkEventButton)(event))
}

// SignalButtonRelease connects a callback for "button_release_event" event
func SignalButtonRelease(callback func(event *GdkEventButton)) { buttonReleaseFn = callback }

//export widgetMotionNotify
func widgetMotionNotify(widget *C.GtkWidget, event *C.GdkEventMotion, data C.gpointer) {
	if event.time == prevButtonTime && event._type == prevButtonType {
		return
	}
	prevButtonTime = event.time
	prevButtonType = event._type
	motionNotifyFn((*GdkEventMotion)(event))
}

// SignalMotionNotify connects a callback for "motion_notify_event" event
func SignalMotionNotify(callback func(event *GdkEventMotion)) { motionNotifyFn = callback }

//export widgetScroll
func widgetScroll(widget *C.GtkWidget, event *C.GdkEventScroll, data C.gpointer) {
	scrollFn((*GdkEventScroll)(event))
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
	gSignalConnect(top.widget.GObject(), "delete-event", C.GCallback(C.widget_delete), nil)
	gSignalConnect(top.widget.GObject(), "key_press_event", C.GCallback(C.widget_key_press), nil)
	gSignalConnect(top.widget.GObject(), "button_press_event", C.GCallback(C.widget_button_press), nil)
	gSignalConnect(top.widget.GObject(), "button_release_event", C.GCallback(C.widget_button_release), nil)
	gSignalConnect(top.widget.GObject(), "motion_notify_event", C.GCallback(C.widget_motion_notify), nil)
	gSignalConnect(top.widget.GObject(), "scroll-event", C.GCallback(C.widget_scroll), nil)
}

// TopSignalDisconnect disconnects application window level signals
func (top *TopWindow) TopSignalDisconnect() {
	C.GSignalHandlersDisconnectByFunc(top.widget.GObject(), C.GCallback(C.widget_delete), nil)
	C.GSignalHandlersDisconnectByFunc(top.widget.GObject(), C.GCallback(C.widget_key_press), nil)
	C.GSignalHandlersDisconnectByFunc(top.widget.GObject(), C.GCallback(C.widget_button_press), nil)
	C.GSignalHandlersDisconnectByFunc(top.widget.GObject(), C.GCallback(C.widget_button_release), nil)
	C.GSignalHandlersDisconnectByFunc(top.widget.GObject(), C.GCallback(C.widget_motion_notify), nil)
	C.GSignalHandlersDisconnectByFunc(top.widget.GObject(), C.GCallback(C.widget_scroll), nil)
}

// LayoutSignalConnect connects layout level signals
func (layout *Layout) LayoutSignalConnect() {
	gSignalConnect(layout.widget.GObject(), "size-allocate", C.GCallback(C.widget_size_allocate), nil)
}

// LayoutSignalDisconnect disconnects layout level signals
func (layout *Layout) LayoutSignalDisconnect() {
	C.GSignalHandlersDisconnectByFunc(layout.widget.GObject(), C.GCallback(C.widget_size_allocate), nil)
}

// SignalDraw connects a callback parameters for "draw" event
func (drawing *Drawing) SignalDraw(h *cgo.Handle) {
	drawing.drawHandler = gSignalConnect(drawing.Widget().GObject(), "draw", C.GCallback(C.window_draw), C.gpointer(h))
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

//export clipboardTextReceived
func clipboardTextReceived(clipboard *C.GtkClipboard, text *C.gchar, data *C.gpointer) {
	clipboardFn(C.GoString(text))
}

// RequestClipboardText connects a callback for clipboard receiving
func RequestClipboardText(callback func(text string)) {
	clipboardFn = callback
	clipboard := C.gtk_clipboard_get(C.GDK_SELECTION_CLIPBOARD)
	C.gtk_clipboard_request_text(clipboard, C.GCallback(C.clipboard_text_received), nil)
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
