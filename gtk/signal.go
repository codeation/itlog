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
	drawing.drawHandler = gSignalConnect(drawing.Widget().GObject(), "draw", C.GCallback(C.window_draw), C.gpointer(h))
}

//export widgetDelete
func widgetDelete(widget *C.GtkWidget, event *C.GdkEvent, data C.gpointer) { deleteFn() }

// SignalDelete connects a callback for "delete-event" event
func (widget *Widget) SignalDelete(callback func()) {
	deleteFn = callback
	gSignalConnect(widget.GObject(), "delete-event", C.GCallback(C.widget_delete), nil)
}

//export widgetSizeAllocate
func widgetSizeAllocate(widget *C.GtkWidget, allocation *C.GtkAllocation, data C.gpointer) {
	sizeAllocateFn((*GtkAllocation)(allocation))
}

// SignalSizeAllocate connects a callback for "size-allocate" event
func (layout *Layout) SignalSizeAllocate(callback func(allocation *GtkAllocation)) {
	sizeAllocateFn = callback
	gSignalConnect(layout.Widget().GObject(), "size-allocate", C.GCallback(C.widget_size_allocate), nil)
}

//export widgetKeyPress
func widgetKeyPress(widget *C.GtkWidget, event *C.GdkEventKey, data C.gpointer) {
	keyPressFn((*GdkEventKey)(event))
}

// SignalKeyPress connects a callback for "key_press_event" event
func (widget *Widget) SignalKeyPress(callback func(event *GdkEventKey)) {
	keyPressFn = callback
	gSignalConnect(widget.GObject(), "key_press_event", C.GCallback(C.widget_key_press), nil)
}

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
func (widget *Widget) SignalButtonPress(callback func(event *GdkEventButton)) {
	buttonPressFn = callback
	gSignalConnect(widget.GObject(), "button_press_event", C.GCallback(C.widget_button_press), nil)
}

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
func (widget *Widget) SignalButtonRelease(callback func(event *GdkEventButton)) {
	buttonReleaseFn = callback
	gSignalConnect(widget.GObject(), "button_release_event", C.GCallback(C.widget_button_release), nil)
}

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
func (widget *Widget) SignalMotionNotify(callback func(event *GdkEventMotion)) {
	motionNotifyFn = callback
	gSignalConnect(widget.GObject(), "motion_notify_event", C.GCallback(C.widget_motion_notify), nil)
}

//export widgetScroll
func widgetScroll(widget *C.GtkWidget, event *C.GdkEventScroll, data C.gpointer) {
	scrollFn((*GdkEventScroll)(event))
}

// SignalScroll connects a callback for "scroll-event" event
func (widget *Widget) SignalScroll(callback func(event *GdkEventScroll)) {
	scrollFn = callback
	gSignalConnect(widget.GObject(), "scroll-event", C.GCallback(C.widget_scroll), nil)
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
