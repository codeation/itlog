package gtk4

// #cgo pkg-config: gtk4
// #include <gtk/gtk.h>
// #include "macro.h"
// #include "watchio.h"
import "C"
import (
	"errors"
	"fmt"
	"os"
)

type fder interface {
	Fd() uintptr
}

var (
	streamReadFn  func()
	requestReadFn func()
)

//export streamDo
func streamDo() {
	streamReadFn()
}

//export requestDo
func requestDo() {
	requestReadFn()
}

//export chanErr
func chanErr() {
	fmt.Println("stream io error")
	os.Exit(1)
}

// WatchIO is a channel watch IDs
type WatchIO struct {
	inputID C.uint
	errorID C.uint
}

// NewStreamIO returns a watchers for async pipe
func NewStreamIO(streamFile fder, readFn func()) (*WatchIO, error) {
	streamReadFn = readFn

	gioChan := C.g_io_channel_unix_new(C.int(streamFile.Fd()))
	if gioChan == nil {
		return nil, errors.New("channel open error")
	}
	defer C.g_io_channel_unref(gioChan)

	if C.g_io_channel_set_encoding(gioChan, nil, nil) != C.G_IO_STATUS_NORMAL {
		return nil, errors.New("encoding status error")
	}
	C.g_io_channel_set_buffer_size(gioChan, 64*1024)

	return &WatchIO{
		inputID: C.g_io_add_watch(gioChan, C.G_IO_IN, C.GCallback(C.stream_read_chan), nil),
		errorID: C.g_io_add_watch(gioChan, C.G_IO_HUP, C.GCallback(C.chan_error_func), nil),
	}, nil
}

// NewRequestIO returns a watchers for full duplex pipe
func NewRequestIO(syncFile fder, readFn func()) (*WatchIO, error) {
	requestReadFn = readFn

	gioChan := C.g_io_channel_unix_new(C.int(syncFile.Fd()))
	if gioChan == nil {
		return nil, errors.New("channel open error")
	}
	defer C.g_io_channel_unref(gioChan)

	if C.g_io_channel_set_encoding(gioChan, nil, nil) != C.G_IO_STATUS_NORMAL {
		return nil, errors.New("encoding status error")
	}

	return &WatchIO{
		inputID: C.g_io_add_watch(gioChan, C.G_IO_IN, C.GCallback(C.request_read_chan), nil),
		errorID: C.g_io_add_watch(gioChan, C.G_IO_HUP, C.GCallback(C.chan_error_func), nil),
	}, nil
}

// Done removes channel watchers
func (w *WatchIO) Done() {
	C.g_source_remove(w.inputID)
	C.g_source_remove(w.errorID)
}
