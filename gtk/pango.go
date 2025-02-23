package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "macro.h"
import "C"
import (
	"unsafe"
)

// FontSelection is a PangoFontDescription wrapper
type FontSelection struct {
	desc        *C.PangoFontDescription
	family      *C.char
	layout      *C.PangoLayout
	splitLayout *C.PangoLayout
}

func fontStyle(style int) C.PangoStyle {
	switch style {
	case 0:
		return C.PANGO_STYLE_NORMAL
	case 1:
		return C.PANGO_STYLE_OBLIQUE
	case 2:
		return C.PANGO_STYLE_ITALIC
	default:
		return C.PANGO_STYLE_NORMAL
	}
}

func fontVariant(variant int) C.PangoVariant {
	switch variant {
	case 0:
		return C.PANGO_VARIANT_NORMAL
	case 1:
		return C.PANGO_VARIANT_SMALL_CAPS
	case 2:
		return C.PANGO_VARIANT_ALL_SMALL_CAPS
	case 3:
		return C.PANGO_VARIANT_PETITE_CAPS
	case 4:
		return C.PANGO_VARIANT_ALL_PETITE_CAPS
	case 5:
		return C.PANGO_VARIANT_UNICASE
	case 6:
		return C.PANGO_VARIANT_TITLE_CAPS
	default:
		return C.PANGO_VARIANT_NORMAL
	}
}

func fontWeight(weigth int) C.PangoWeight {
	switch weigth {
	case 100:
		return C.PANGO_WEIGHT_THIN
	case 200:
		return C.PANGO_WEIGHT_ULTRALIGHT
	case 300:
		return C.PANGO_WEIGHT_LIGHT
	case 350:
		return C.PANGO_WEIGHT_SEMILIGHT
	case 380:
		return C.PANGO_WEIGHT_BOOK
	case 400:
		return C.PANGO_WEIGHT_NORMAL
	case 500:
		return C.PANGO_WEIGHT_MEDIUM
	case 600:
		return C.PANGO_WEIGHT_SEMIBOLD
	case 700:
		return C.PANGO_WEIGHT_BOLD
	case 800:
		return C.PANGO_WEIGHT_ULTRABOLD
	case 900:
		return C.PANGO_WEIGHT_HEAVY
	case 1000:
		return C.PANGO_WEIGHT_ULTRAHEAVY
	default:
		return C.PANGO_WEIGHT_NORMAL
	}
}

func fontStretch(stretch int) C.PangoStretch {
	switch stretch {
	case 0:
		return C.PANGO_STRETCH_ULTRA_CONDENSED
	case 1:
		return C.PANGO_STRETCH_EXTRA_CONDENSED
	case 2:
		return C.PANGO_STRETCH_CONDENSED
	case 3:
		return C.PANGO_STRETCH_SEMI_CONDENSED
	case 4:
		return C.PANGO_STRETCH_NORMAL
	case 5:
		return C.PANGO_STRETCH_SEMI_EXPANDED
	case 6:
		return C.PANGO_STRETCH_EXPANDED
	case 7:
		return C.PANGO_STRETCH_EXTRA_EXPANDED
	case 8:
		return C.PANGO_STRETCH_ULTRA_EXPANDED
	default:
		return C.PANGO_STRETCH_NORMAL
	}
}

// NewFontSelection creates FontSelection
func NewFontSelection(height int, family string, style int, variant int, weight int, stretch int,
	top *TopWindow,
) *FontSelection {
	pangoContext := C.gtk_widget_get_pango_context(top.Widget().GtkWidget())
	f := &FontSelection{
		desc:        C.pango_font_description_new(),
		family:      C.CString(family),
		layout:      C.pango_layout_new(pangoContext),
		splitLayout: C.pango_layout_new(pangoContext),
	}
	C.pango_font_description_set_family(f.desc, f.family)
	C.pango_font_description_set_absolute_size(f.desc, C.double(C.PANGO_SCALE*height))
	C.pango_font_description_set_style(f.desc, fontStyle(style))
	C.pango_font_description_set_variant(f.desc, fontVariant(variant))
	C.pango_font_description_set_weight(f.desc, fontWeight(weight))
	C.pango_font_description_set_stretch(f.desc, fontStretch(stretch))
	C.pango_layout_set_font_description(f.layout, f.desc)
	C.pango_layout_set_font_description(f.splitLayout, f.desc)
	C.pango_layout_set_wrap(f.splitLayout, C.PANGO_WRAP_WORD_CHAR)
	return f
}

// Free destroys font selection
func (f *FontSelection) Free() {
	C.g_object_unref(C.layoutToGPointer(f.layout))
	C.g_object_unref(C.layoutToGPointer(f.splitLayout))
	C.pango_font_description_free(f.desc)
	C.free(unsafe.Pointer(f.family))
}

// Metrics returns font selection lineheight, baseline, ascent, descent
func (f *FontSelection) Metrics() (int, int, int, int) {
	baseline := int(float64(C.pango_layout_get_baseline(f.layout)) / float64(C.PANGO_SCALE))
	metrics := C.pango_context_get_metrics(C.pango_layout_get_context(f.layout), f.desc, nil)
	lineheight := int(float64(C.pango_font_metrics_get_height(metrics)) / float64(C.PANGO_SCALE))
	ascent := int(float64(C.pango_font_metrics_get_ascent(metrics)) / float64(C.PANGO_SCALE))
	descent := int(float64(C.pango_font_metrics_get_descent(metrics)) / float64(C.PANGO_SCALE))
	C.pango_font_metrics_unref(metrics)
	return lineheight, baseline, ascent, descent
}

// Split splits text string to substrings
func (f *FontSelection) Split(text string, edge, indent int) []int {
	C.pango_layout_set_width(f.splitLayout, C.int(C.PANGO_SCALE*edge))
	C.pango_layout_set_indent(f.splitLayout, C.int(C.PANGO_SCALE*indent))
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.pango_layout_set_text(f.splitLayout, cText, -1)
	output := make([]int, 0, 8)
	for e := C.pango_layout_get_lines_readonly(f.splitLayout); e != nil; e = e.next {
		output = append(output, int((*C.PangoLayoutLine)(e.data).length))
	}
	return output
}

// Size returns text string width and height in pixels
func (f *FontSelection) Size(text string) (int, int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.pango_layout_set_text(f.layout, cText, -1)
	var w, h C.int
	C.pango_layout_get_pixel_size(f.layout, &w, &h)
	return int(w), int(h)
}
