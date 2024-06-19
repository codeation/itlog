package gtk4

// #cgo pkg-config: gtk4
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"

// Widget is a GtkWidget wrapper
type Widget struct {
	gtkWidget *C.GtkWidget
}

func newWidget(widget *C.GtkWidget) *Widget {
	return &Widget{
		gtkWidget: widget,
	}
}

// GtkWidget returns C.GtkWidget pointer
func (widget *Widget) GtkWidget() *C.GtkWidget { return widget.gtkWidget }

// GObject returns C.GObject pointer from widget
func (widget *Widget) GObject() *C.GObject { return C.widgetToGObject(widget.gtkWidget) }

// type TopWindow contains GtkWindow widget
type TopWindow struct {
	widget *Widget
}

// NewTopWindow creates a WindowWidget
func (app *Application) NewTopWindow() *TopWindow {
	return &TopWindow{
		widget: newWidget(C.gtk_application_window_new(app.GtkApplication())),
	}
}

// Widget returns a Widget pointer
func (window *TopWindow) Widget() *Widget { return window.widget }

// Destroy destroys widget
func (window *TopWindow) Destroy() {
	C.gtk_window_destroy(C.widgetToGtkWindow(window.widget.gtkWidget))
}

// ShowAll shows a window with all child widgets
func (window *TopWindow) ShowAll() {
	C.gtk_application_window_set_show_menubar(C.widgetToGtkApplicationWindow(window.widget.gtkWidget), 1)
	C.gtk_window_present(C.widgetToGtkWindow(window.widget.gtkWidget))
}

// Size move and resize application window
func (window *TopWindow) Size(x, y, width, height int) {
	C.gtk_window_set_default_size(C.widgetToGtkWindow(window.widget.gtkWidget), C.int(width), C.int(height))
}

type container interface {
	addChild(child *Widget)
	removeChild(child *Widget)
	moveChild(child *Widget, x, y int)
}

func raiseChild(parent container, child *Widget) {
	C.g_object_ref(C.widgetToGPointer(child.gtkWidget))
	parent.removeChild(child)
	parent.addChild(child)
	C.g_object_unref(C.widgetToGPointer(child.gtkWidget))
}

// Layout contains GtkLayout widget
type Layout struct {
	widget     *Widget
	parent     *TopWindow
	scrolled   *C.GtkWidget
	adjustment *C.GtkAdjustment
}

// NewLayout creates a LayoutWidget
func (window *TopWindow) NewLayout() *Layout {
	scrolled := C.gtk_scrolled_window_new()
	C.gtk_window_set_child(C.widgetToGtkWindow(window.widget.gtkWidget), scrolled)
	C.gtk_scrolled_window_set_policy(C.widgetToGtkScrolledWindow(scrolled), C.GTK_POLICY_EXTERNAL, C.GTK_POLICY_EXTERNAL)

	adjustment := C.gtk_adjustment_new(0, 0, 1, 0, 0, 1)
	C.gtk_scrolled_window_set_hadjustment(C.widgetToGtkScrolledWindow(scrolled), adjustment)
	C.gtk_scrolled_window_set_vadjustment(C.widgetToGtkScrolledWindow(scrolled), adjustment)

	fixed := C.gtk_fixed_new()
	C.gtk_scrolled_window_set_child(C.widgetToGtkScrolledWindow(scrolled), fixed)
	C.gtk_widget_set_overflow(fixed, C.GTK_OVERFLOW_VISIBLE)

	return &Layout{
		widget:     newWidget(fixed),
		parent:     window,
		scrolled:   scrolled,
		adjustment: adjustment,
	}
}

// Widget returns a Widget pointer
func (layout *Layout) Widget() *Widget { return layout.widget }

// Destroy destroys widget
func (layout *Layout) Destroy() {}

// NewFixed creates a FixedWidget child
func (layout *Layout) NewFixed() *Fixed { return newFixed(layout) }

// NewDrawingArea creates a DrawingWidget child
func (layout *Layout) NewDrawingArea() *Drawing { return newDrawing(layout) }

// Show shows widget
func (layout *Layout) Show() {
	C.gtk_widget_show(layout.scrolled)
	C.gtk_widget_show(layout.widget.gtkWidget)
}

// Size resizes widget size
func (layout *Layout) Size(width, height int) {}

// Move is a empty method
func (layout *Layout) Move(x, y int) {}

// Raise is a empty method
func (layout *Layout) Raise() {}

func (layout *Layout) addChild(child *Widget) {
	C.gtk_widget_set_parent(child.gtkWidget, layout.widget.gtkWidget)
}

func (layout *Layout) removeChild(child *Widget) {
	C.gtk_widget_unparent(child.gtkWidget)
}

func (layout *Layout) moveChild(child *Widget, x, y int) {
	C.gtk_fixed_move(C.widgetToGtkFixed(layout.widget.gtkWidget), child.gtkWidget, C.double(x), C.double(y))
}

// type Fixed contains GtkFixed widget
type Fixed struct {
	widget *Widget
	parent container
}

func newFixed(parent container) *Fixed {
	fixed := &Fixed{
		widget: newWidget(C.gtk_fixed_new()),
		parent: parent,
	}
	C.gtk_widget_set_overflow(fixed.widget.gtkWidget, C.GTK_OVERFLOW_HIDDEN)
	parent.addChild(fixed.widget)
	return fixed
}

// Widget returns a Widget pointer
func (fixed *Fixed) Widget() *Widget { return fixed.widget }

// Destroy destroys widget
func (fixed *Fixed) Destroy() { C.gtk_widget_unparent(fixed.widget.gtkWidget) }

// NewFixed creates a FixedWidget child
func (fixed *Fixed) NewFixed() *Fixed { return newFixed(fixed) }

// NewDrawingArea creates a DrawingWidget child
func (fixed *Fixed) NewDrawingArea() *Drawing { return newDrawing(fixed) }

// Show shows widget
func (fixed *Fixed) Show() { C.gtk_widget_show(fixed.widget.gtkWidget) }

// Size resizes widget size
func (fixed *Fixed) Size(width, height int) {
	C.gtk_widget_set_size_request(fixed.widget.gtkWidget, C.int(width), C.int(height))
}

// Move shifts widget to a new coordinates
func (fixed *Fixed) Move(x, y int) { fixed.parent.moveChild(fixed.widget, x, y) }

// Raise raises widget
func (fixed *Fixed) Raise() { raiseChild(fixed.parent, fixed.widget) }

func (fixed *Fixed) addChild(child *Widget) {
	C.gtk_widget_set_parent(child.gtkWidget, fixed.widget.gtkWidget)
}

func (fixed *Fixed) removeChild(child *Widget) {
	C.gtk_widget_unparent(child.gtkWidget)
}

func (fixed *Fixed) moveChild(child *Widget, x, y int) {
	C.gtk_fixed_move(C.widgetToGtkFixed(fixed.widget.gtkWidget), child.gtkWidget, C.double(x), C.double(y))
}

// Drawing contains DrawingArea widget
type Drawing struct {
	widget *Widget
	parent container
}

func newDrawing(parent container) *Drawing {
	drawing := &Drawing{
		widget: newWidget(C.gtk_drawing_area_new()),
		parent: parent,
	}
	parent.addChild(drawing.widget)
	return drawing
}

// Widget returns a Widget pointer
func (drawing *Drawing) Widget() *Widget { return drawing.widget }

// Destroy destroys widget
func (drawing *Drawing) Destroy() { C.gtk_widget_unparent(drawing.widget.gtkWidget) }

// Show shows widget
func (drawing *Drawing) Show() { C.gtk_widget_show(drawing.widget.gtkWidget) }

// Size resizes widget size
func (drawing *Drawing) Size(width, height int) {
	C.gtk_drawing_area_set_content_width(C.widgetToGtkDrawingArea(drawing.widget.gtkWidget), C.int(width))
	C.gtk_drawing_area_set_content_height(C.widgetToGtkDrawingArea(drawing.widget.gtkWidget), C.int(height))
}

// Move shifts widget to a new coordinates
func (drawing *Drawing) Move(x, y int) { drawing.parent.moveChild(drawing.widget, x, y) }

// Raise raises widget
func (drawing *Drawing) Raise() { raiseChild(drawing.parent, drawing.widget) }

// QueueDraw marks a widget to redrawing
func (drawing *Drawing) QueueDraw() { C.gtk_widget_queue_draw(drawing.widget.gtkWidget) }
