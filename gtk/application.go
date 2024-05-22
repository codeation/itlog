package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"
import (
	"os"
	"unsafe"
)

type Application struct {
	a              *C.GtkApplication
	signalHandlers []C.gulong
}

func NewApplication() *Application {
	return &Application{
		a: C.gtk_application_new(nil, C.G_APPLICATION_FLAGS_NONE),
	}
}

func (app *Application) GtkApplication() *C.GtkApplication { return app.a }
func (app *Application) GObject() *C.GObject               { return C.appToGObject(app.a) }
func (app *Application) GApplication() *C.GApplication     { return C.appToGApplication(app.a) }
func (app *Application) GActionMap() *C.GActionMap         { return C.appToGActionMap(app.a) }
func (app *Application) GPointer() C.gpointer              { return C.appToGPointer(app.a) }

func (app *Application) Quit() {
	for _, signalHandlerID := range app.signalHandlers {
		C.g_signal_handler_disconnect(app.GPointer(), signalHandlerID)
	}
	C.g_application_quit(app.GApplication())
}

func (app *Application) Run() {
	procName := C.CString(os.Args[0])
	defer C.free(unsafe.Pointer(procName))
	C.g_application_run(app.GApplication(), 1, &procName)
	C.g_object_unref(app.GPointer())
}

func (app *Application) SetName(name string) {
	appName := C.CString(name)
	defer C.free(unsafe.Pointer(appName))
	C.g_set_application_name(appName)
}
