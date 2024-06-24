package gtk4

// #cgo pkg-config: gtk4
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"
import (
	"unsafe"
)

// SetClipboardText sets clipboard content
func SetClipboardText(top *TopWindow, text string) {
	data := C.CString(text)
	defer C.free(unsafe.Pointer(data))
	var value C.GValue
	C.g_value_init(&value, C.G_TYPE_STRING)
	C.g_value_set_string(&value, data)
	clipboard := C.gtk_widget_get_clipboard(top.Widget().GtkWidget())
	C.gdk_clipboard_set_value(clipboard, &value)
	C.g_value_unset(&value)
}
