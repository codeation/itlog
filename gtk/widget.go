package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"

type Widget struct {
	w              *C.GtkWidget
	signalHandlers []C.gulong
}

func (widget *Widget) Widget() *C.GtkWidget { return widget.w }
func (widget *Widget) GPointer() C.gpointer { return C.widgetToGPointer(widget.w) }
func (widget *Widget) GObject() *C.GObject  { return C.widgetToGObject(widget.w) }

func (widget *Widget) Destroy() {
	for _, signalHandlerID := range widget.signalHandlers {
		C.g_signal_handler_disconnect(widget.GPointer(), signalHandlerID)
	}
	C.gtk_widget_destroy(widget.w)
}

func (widget *Widget) Show()    { C.gtk_widget_show(widget.w) }
func (widget *Widget) ShowAll() { C.gtk_widget_show_all(widget.w) }

// type WindowWidget contain GtkApplicationWindow widget
type WindowWidget struct {
	*Widget
}

func (app *Application) NewWindow() *WindowWidget {
	return &WindowWidget{
		Widget: &Widget{
			w: C.gtk_application_window_new(app.GtkApplication()),
		},
	}
}

func (widget *WindowWidget) GtkWindow() *C.GtkWindow { return C.widgetToGtkWindow(widget.w) }

func (widget *WindowWidget) Size(x, y, width, height int) {
	C.gtk_window_move(widget.GtkWindow(), C.int(x), C.int(y))
	C.gtk_window_resize(widget.GtkWindow(), C.int(width), C.int(height))
}

// type FrameWidget contain GtkLayout, GtkFixed or DrawingArea widget
type FrameWidget struct {
	*Widget
	parent *Widget
}

func (parent *Widget) NewLayout() *FrameWidget {
	layout := &FrameWidget{
		Widget: &Widget{
			w: C.gtk_layout_new(nil, nil),
		},
		parent: parent,
	}
	C.gtk_container_add(C.widgetToGtkContainer(parent.w), layout.w)
	return layout
}

func (parent *Widget) NewFixed() *FrameWidget {
	fixed := &FrameWidget{
		Widget: &Widget{
			w: C.gtk_fixed_new(),
		},
		parent: parent,
	}
	C.gtk_container_add(C.widgetToGtkContainer(parent.w), fixed.w)
	return fixed
}

func (parent *FrameWidget) NewDrawingArea() *FrameWidget {
	drawing := &FrameWidget{
		Widget: &Widget{
			w: C.gtk_drawing_area_new(),
		},
		parent: parent.Widget,
	}
	C.gtk_container_add(C.widgetToGtkContainer(parent.w), drawing.w)
	return drawing
}

func (frame *FrameWidget) Size(width, height int) {
	C.gtk_widget_set_size_request(frame.w, C.int(width), C.int(height))
}

func (frame *FrameWidget) Move(x, y int) {
	if C.widgetIsLayout(frame.parent.w) != 0 {
		C.gtk_layout_move(C.widgetToGtkLayout(frame.parent.w), frame.w, C.int(x), C.int(y))
	} else {
		C.gtk_fixed_move(C.widgetToGtkFixed(frame.parent.w), frame.w, C.int(x), C.int(y))
	}
}

func (frame *FrameWidget) Raise() {
	C.g_object_ref(frame.GPointer())
	C.gtk_container_remove(C.widgetToGtkContainer(frame.parent.w), frame.w)
	C.gtk_container_add(C.widgetToGtkContainer(frame.parent.w), frame.w)
	C.g_object_unref(frame.GPointer())
}

func (frame *FrameWidget) QueueDraw() {
	C.gtk_widget_queue_draw(frame.w)
}
