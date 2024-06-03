package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"

type container interface {
	addChild(child *Widget)
	moveChild(child *Widget, x, y int)
	raiseChild(child *Widget)
	NewFixed() *FixedWidget
	NewDrawingArea() *DrawingWidget
}

// Widget is a GtkWidget wrapper
type Widget struct {
	w              *C.GtkWidget
	signalHandlers []C.gulong
}

func newWidget(w *C.GtkWidget) *Widget {
	return &Widget{
		w: w,
	}
}

// GtkWidget returns C.GtkWidget pointer
func (widget *Widget) GtkWidget() *C.GtkWidget { return widget.w }

// GPointer returns a C.gpointer address
func (widget *Widget) GPointer() C.gpointer { return C.widgetToGPointer(widget.w) }

// GObject returns C.GObject pointer from widget
func (widget *Widget) GObject() *C.GObject { return C.widgetToGObject(widget.w) }

// Destroy destroys widget
func (widget *Widget) Destroy() {
	for _, signalHandlerID := range widget.signalHandlers {
		C.g_signal_handler_disconnect(widget.GPointer(), signalHandlerID)
	}
	C.gtk_widget_destroy(widget.w)
}

// Show shows widget
func (widget *Widget) Show() { C.gtk_widget_show(widget.w) }

func (widget *Widget) size(width, height int) {
	C.gtk_widget_set_size_request(widget.w, C.int(width), C.int(height))
}

func (widget *Widget) addChild(child *Widget) {
	C.gtk_container_add(C.widgetToGtkContainer(widget.w), child.w)
}

func (widget *Widget) removeChild(child *Widget) {
	C.gtk_container_remove(C.widgetToGtkContainer(widget.w), child.w)
}

func (widget *Widget) raiseChild(child *Widget) {
	C.g_object_ref(child.GPointer())
	widget.removeChild(child)
	widget.addChild(child)
	C.g_object_unref(child.GPointer())
}

// type WindowWidget contains GtkWindow widget
type WindowWidget struct {
	*Widget
}

// NewWindow creates a WindowWidget
func (app *Application) NewWindow() *WindowWidget {
	return &WindowWidget{
		Widget: newWidget(C.gtk_application_window_new(app.GtkApplication())),
	}
}

// GtkWindow returns a C.GtkWindow pointer
func (widget *WindowWidget) GtkWindow() *C.GtkWindow { return C.widgetToGtkWindow(widget.w) }

// ShowAll shows a window with all child widgets
func (widget *WindowWidget) ShowAll() { C.gtk_widget_show_all(widget.w) }

// Size move and resize application window
func (widget *WindowWidget) Size(x, y, width, height int) {
	C.gtk_window_move(widget.GtkWindow(), C.int(x), C.int(y))
	C.gtk_window_resize(widget.GtkWindow(), C.int(width), C.int(height))
}

// LayoutWidget contains GtkLayout widget
type LayoutWidget struct {
	*Widget
}

// NewLayout creates a LayoutWidget
func (widget *WindowWidget) NewLayout() *LayoutWidget {
	layout := &LayoutWidget{
		Widget: newWidget(C.gtk_layout_new(nil, nil)),
	}
	widget.addChild(layout.Widget)
	return layout
}

// NewFixed creates a FixedWidget child
func (widget *LayoutWidget) NewFixed() *FixedWidget { return newFixedWidget(widget) }

// NewDrawingArea creates a DrawingWidget child
func (widget *LayoutWidget) NewDrawingArea() *DrawingWidget { return newDrawingArea(widget) }

// Size resizes widget size
func (widget *LayoutWidget) Size(width, height int) { widget.size(width, height) }

// Move is a empty method
func (widget *LayoutWidget) Move(x, y int) {}

// Raise is a empty method
func (widget *LayoutWidget) Raise() {}

func (widget *LayoutWidget) moveChild(child *Widget, x, y int) {
	C.gtk_layout_move(C.widgetToGtkLayout(widget.w), child.w, C.int(x), C.int(y))
}

// type FixedWidget contains GtkFixed widget
type FixedWidget struct {
	*Widget
	parent container
}

func newFixedWidget(parent container) *FixedWidget {
	fixed := &FixedWidget{
		Widget: newWidget(C.gtk_fixed_new()),
		parent: parent,
	}
	parent.addChild(fixed.Widget)
	return fixed
}

// NewFixed creates a FixedWidget child
func (widget *FixedWidget) NewFixed() *FixedWidget { return newFixedWidget(widget) }

// NewDrawingArea creates a DrawingWidget child
func (widget *FixedWidget) NewDrawingArea() *DrawingWidget { return newDrawingArea(widget) }

// Size resizes widget size
func (widget *FixedWidget) Size(width, height int) { widget.size(width, height) }

// Move shifts widget to a new coordinates
func (widget *FixedWidget) Move(x, y int) { widget.parent.moveChild(widget.Widget, x, y) }

// Raise raises widget
func (widget *FixedWidget) Raise() { widget.parent.raiseChild(widget.Widget) }

func (widget *FixedWidget) moveChild(child *Widget, x, y int) {
	C.gtk_fixed_move(C.widgetToGtkFixed(widget.w), child.w, C.int(x), C.int(y))
}

// DrawingWidget contains DrawingArea widget
type DrawingWidget struct {
	*Widget
	parent container
}

func newDrawingArea(parent container) *DrawingWidget {
	drawing := &DrawingWidget{
		Widget: newWidget(C.gtk_drawing_area_new()),
		parent: parent,
	}
	parent.addChild(drawing.Widget)
	return drawing
}

// Size resizes widget size
func (widget *DrawingWidget) Size(width, height int) { widget.size(width, height) }

// Move shifts widget to a new coordinates
func (widget *DrawingWidget) Move(x, y int) { widget.parent.moveChild(widget.Widget, x, y) }

// Raise raises widget
func (widget *DrawingWidget) Raise() { widget.parent.raiseChild(widget.Widget) }

// QueueDraw marks a widget to redrawing
func (widget *DrawingWidget) QueueDraw() {
	C.gtk_widget_queue_draw(widget.w)
}
