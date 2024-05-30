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

//export appActivate
func appActivate(application *C.GApplication, data C.gpointer) { activateFn() }

func (app *Application) SignalActivate(callback func()) {
	activateFn = callback
	name := C.CString("activate")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(app.GObject(), name, C.GCallback(C.app_activate), nil)
	app.signalHandlers = append(app.signalHandlers, signalHandlerID)
}

//export appShutdown
func appShutdown(application *C.GApplication, data C.gpointer) { shutdownFn() }

func (app *Application) SignalShutdown(callback func()) {
	shutdownFn = callback
	name := C.CString("shutdown")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(app.GObject(), name, C.GCallback(C.app_shutdown), nil)
	app.signalHandlers = append(app.signalHandlers, signalHandlerID)
}

//export windowDraw
func windowDraw(widget *C.GtkWidget, cr *C.cairo_t, data C.gpointer) {
	drawFn((*Cairo)(cr), (*cgo.Handle)(data))
}

func SignalDrawCallback(callback func(cr *Cairo, h *cgo.Handle)) {
	drawFn = callback
}

func (widget *Widget) SignalDraw(h *cgo.Handle) {
	name := C.CString("draw")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.window_draw), C.gpointer(h))
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
}

//export widgetDelete
func widgetDelete(widget *C.GtkWidget, data C.gpointer) { deleteFn() }

func (widget *Widget) SignalDelete(callback func()) {
	deleteFn = callback
	name := C.CString("delete-event")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.widget_delete), nil)
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
}

//export widgetSizeAllocate
func widgetSizeAllocate(widget *C.GtkWidget, allocation *C.GtkAllocation, data C.gpointer) {
	sizeAllocateFn((*GtkAllocation)(allocation))
}

func (widget *Widget) SignalSizeAllocate(callback func(allocation *GtkAllocation)) {
	sizeAllocateFn = callback
	name := C.CString("size-allocate")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.widget_size_allocate), nil)
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
}

//export widgetKeyPress
func widgetKeyPress(widget *C.GtkWidget, event *C.GdkEventKey, data C.gpointer) {
	keyPressFn((*GdkEventKey)(event))
}

func (widget *Widget) SignalKeyPress(callback func(event *GdkEventKey)) {
	keyPressFn = callback
	name := C.CString("key_press_event")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.widget_key_press), nil)
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
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

func (widget *Widget) SignalButtonPress(callback func(event *GdkEventButton)) {
	buttonPressFn = callback
	name := C.CString("button_press_event")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.widget_button_press), nil)
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
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

func (widget *Widget) SignalButtonRelease(callback func(event *GdkEventButton)) {
	buttonReleaseFn = callback
	name := C.CString("button_release_event")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.widget_button_release), nil)
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
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

func (widget *Widget) SignalMotionNotify(callback func(event *GdkEventMotion)) {
	motionNotifyFn = callback
	name := C.CString("motion_notify_event")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.widget_motion_notify), nil)
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
}

//export widgetScroll
func widgetScroll(widget *C.GtkWidget, event *C.GdkEventScroll, data C.gpointer) {
	scrollFn((*GdkEventScroll)(event))
}

func (widget *Widget) SignalScroll(callback func(event *GdkEventScroll)) {
	scrollFn = callback
	name := C.CString("scroll-event")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(widget.GObject(), name, C.GCallback(C.widget_scroll), nil)
	widget.signalHandlers = append(widget.signalHandlers, signalHandlerID)
}

//export menuItemActivate
func menuItemActivate(action *C.GSimpleAction, parameter *C.GVariant, data C.gpointer) {
	menuActionFn((*cgo.Handle)(data))
}

func SignalMenuItemActivateCallback(callback func(h *cgo.Handle)) {
	menuActionFn = callback
}

func (action *MenuAction) SignalMenuItemActivate(h *cgo.Handle) {
	name := C.CString("activate")
	defer C.free(unsafe.Pointer(name))
	signalHandlerID := C.GSignalConnect(action.GObject(), name, C.GCallback(C.menu_item_activate), C.gpointer(h))
	action.signalHandlers = append(action.signalHandlers, signalHandlerID)
}

//export clipboardTextReceived
func clipboardTextReceived(clipboard *C.GtkClipboard, text *C.gchar, data *C.gpointer) {
	clipboardFn(C.GoString(text))
}

func RequestClipboardText(callback func(text string)) {
	clipboardFn = callback
	clipboard := C.gtk_clipboard_get(C.GDK_SELECTION_CLIPBOARD)
	C.gtk_clipboard_request_text(clipboard, C.GCallback(C.clipboard_text_received), nil)
}

//export objectWeakRef
func objectWeakRef(data C.gpointer, where_the_object_was *C.GObject) {
	weakRefFn(uintptr(data))
}

func NotifyWeakRef(object *C.GObject, callback func(data uintptr), data uintptr) {
	weakRefFn = callback
	C.g_object_weak_ref(object, C.GCallback(C.object_weak_notify), C.gpointer(data))
}
