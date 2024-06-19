package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"
import (
	"os"
	"unsafe"
)

// Application is a GtkApplication wrapper
type Application struct {
	a *C.GtkApplication
}

// NewApplication creates Application
func NewApplication() *Application {
	return &Application{
		a: C.gtk_application_new(nil, C.G_APPLICATION_FLAGS_NONE),
	}
}

// GtkApplication returns C.GtkApplication pointer
func (app *Application) GtkApplication() *C.GtkApplication { return app.a }

// GObject returns C.GObject pointer
func (app *Application) GObject() *C.GObject { return C.appToGObject(app.a) }

// GApplication returns C.GApplication type pointer
func (app *Application) GApplication() *C.GApplication { return C.appToGApplication(app.a) }

// GActionMap returns C.GActionMap type pointer
func (app *Application) GActionMap() *C.GActionMap { return C.appToGActionMap(app.a) }

// GPointer returns C.gpointer address
func (app *Application) GPointer() C.gpointer { return C.appToGPointer(app.a) }

// Quit quits application
func (app *Application) Quit() {
	C.g_application_quit(app.GApplication())
}

// Run runs application
func (app *Application) Run() {
	procName := C.CString(os.Args[0])
	defer C.free(unsafe.Pointer(procName))
	C.g_application_run(app.GApplication(), 1, &procName)
	C.g_object_unref(app.GPointer())
}

// SetName sets application name
func (app *Application) SetName(name string) {
	appName := C.CString(name)
	defer C.free(unsafe.Pointer(appName))
	C.g_set_application_name(appName)
}
