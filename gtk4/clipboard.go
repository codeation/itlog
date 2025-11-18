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
	clipboard := C.gtk_widget_get_clipboard(top.Widget().GtkWidget())
	C.gdk_clipboard_set_text(clipboard, data)
}
