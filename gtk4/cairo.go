package gtk4

// #cgo pkg-config: gtk4
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"
import (
	"unsafe"
)

// Cairo is a cairo_t wrapper
type Cairo C.cairo_t

// CairoFillPaint is a fill element parameters
type CairoFillPaint struct {
	x, y, width, height C.double
	r, g, b, a          C.double
}

// NewCairoFillPaint returns a fill element parameters
func NewCairoFillPaint(x, y, width, height int, r, g, b, a uint16) *CairoFillPaint {
	return &CairoFillPaint{
		x:      C.double(x),
		y:      C.double(y),
		width:  C.double(width),
		height: C.double(height),
		r:      C.double(r) / C.double(0xFFFF),
		g:      C.double(g) / C.double(0xFFFF),
		b:      C.double(b) / C.double(0xFFFF),
		a:      C.double(a) / C.double(0xFFFF),
	}
}

// Paint paints current element
func (e *CairoFillPaint) Paint(c *Cairo) {
	cr := (*C.cairo_t)(c)
	C.cairo_set_source_rgba(cr, e.r, e.g, e.b, e.a)
	C.cairo_rectangle(cr, e.x, e.y, e.width, e.height)
	C.cairo_fill(cr)
}

// CairoLinePaint is a line element parameters
type CairoLinePaint struct {
	x0, y0, x1, y1 C.double
	r, g, b, a     C.double
}

// NewCairoLinePaint returns a line element parameters
func NewCairoLinePaint(x0, y0, x1, y1 int, r, g, b, a uint16) *CairoLinePaint {
	var xOffset, yOffset C.double
	if x0 == x1 {
		xOffset = 0.5
	} else if y0 == y1 {
		yOffset = 0.5
	}
	return &CairoLinePaint{
		x0: C.double(x0) + xOffset,
		y0: C.double(y0) + yOffset,
		x1: C.double(x1) + xOffset,
		y1: C.double(y1) + yOffset,
		r:  C.double(r) / C.double(0xFFFF),
		g:  C.double(g) / C.double(0xFFFF),
		b:  C.double(b) / C.double(0xFFFF),
		a:  C.double(a) / C.double(0xFFFF),
	}
}

// Paint paints current element
func (e *CairoLinePaint) Paint(c *Cairo) {
	cr := (*C.cairo_t)(c)
	C.cairo_set_source_rgba(cr, e.r, e.g, e.b, e.a)
	C.cairo_set_line_width(cr, 1)
	C.cairo_move_to(cr, e.x0, e.y0)
	C.cairo_line_to(cr, e.x1, e.y1)
	C.cairo_stroke(cr)
}

// CairoTextPaint is a text element parameters
type CairoTextPaint struct {
	layout     *C.PangoLayout
	desc       *C.PangoFontDescription
	text       *C.char
	x, y       C.double
	r, g, b, a C.double
}

// NewCairoTextPaint returns a text element parameters
func NewCairoTextPaint(x, y int, r, g, b, a uint16, font *FontSelection, text string) *CairoTextPaint {
	return &CairoTextPaint{
		desc: font.desc,
		text: C.CString(text),
		x:    C.double(x),
		y:    C.double(y),
		r:    C.double(r) / C.double(0xFFFF),
		g:    C.double(g) / C.double(0xFFFF),
		b:    C.double(b) / C.double(0xFFFF),
		a:    C.double(a) / C.double(0xFFFF),
	}
}

// Paint paints current element
func (e *CairoTextPaint) Paint(c *Cairo) {
	cr := (*C.cairo_t)(c)
	if e.layout == nil {
		e.layout = C.pango_cairo_create_layout(cr)
		C.pango_layout_set_font_description(e.layout, e.desc)
		C.pango_layout_set_text(e.layout, e.text, -1)
	}
	C.cairo_set_source_rgba(cr, e.r, e.g, e.b, e.a)
	C.cairo_move_to(cr, e.x, e.y)
	C.pango_cairo_show_layout(cr, e.layout)
}

// Destroy destroys current element
func (e *CairoTextPaint) Destroy() {
	C.free(unsafe.Pointer(e.text))
	if e.layout != nil {
		C.g_object_unref(C.layoutToGPointer(e.layout))
	}
}

// CairoBitmap is a cairo_surface_t wrapper
type CairoBitmap struct {
	surface       *C.cairo_surface_t
	buffer        *C.uchar
	width, height int
}

// NewCairoBitmap returns a CairoBitmap from a image bitmap
func NewCairoBitmap(data []byte, width, height int) *CairoBitmap {
	const cFormat = C.CAIRO_FORMAT_ARGB32
	stride := int(C.cairo_format_stride_for_width(cFormat, C.int(width)))
	for i := 0; i < height; i++ {
		offset := i * stride
		for j := 0; j < width; j++ {
			data[offset+j*4+0], data[offset+j*4+2] = data[offset+j*4+2], data[offset+j*4+0]
		}
	}

	buffer := (*C.uchar)(C.CBytes(data))
	return &CairoBitmap{
		surface: C.cairo_image_surface_create_for_data(buffer, cFormat, C.int(width), C.int(height), C.int(stride)),
		buffer:  buffer,
		width:   width,
		height:  height,
	}
}

// Destroy destroys image bitmap buffers
func (e *CairoBitmap) Destroy() {
	C.cairo_surface_destroy(e.surface)
	C.free(unsafe.Pointer(e.buffer))
}

// CairoImagePaint is a image element parameters
type CairoImagePaint struct {
	surface        *C.cairo_surface_t
	scaleX, scaleY C.double
	x, y           C.double
}

// NewCairoImagePaint returns a image element parameters
func NewCairoImagePaint(x, y, width, height int, image *CairoBitmap) *CairoImagePaint {
	return &CairoImagePaint{
		surface: image.surface,
		scaleX:  C.double(image.width) / C.double(width),
		scaleY:  C.double(image.height) / C.double(height),
		x:       C.double(x),
		y:       C.double(y),
	}
}

// Paint paints current element
func (e *CairoImagePaint) Paint(c *Cairo) {
	cr := (*C.cairo_t)(c)
	C.cairo_surface_set_device_scale(e.surface, e.scaleX, e.scaleY)
	C.cairo_set_source_surface(cr, e.surface, e.x, e.y)
	C.cairo_paint(cr)
}
