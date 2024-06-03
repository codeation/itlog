package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"
import (
	"unsafe"
)

// SetClipboardText sets clipboard content
func SetClipboardText(text string) {
	data := C.CString(text)
	defer C.free(unsafe.Pointer(data))
	clipboard := C.gtk_clipboard_get(C.GDK_SELECTION_CLIPBOARD)
	C.gtk_clipboard_set_text(clipboard, data, -1) // TODO text only
}
