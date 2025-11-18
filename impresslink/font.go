package impresslink

import (
	"image"

	"github.com/codeation/impress/driver"
	"github.com/codeation/impress/joint/fontspec"
	gtk "github.com/codeation/itlog/gtk4"
)

type Font struct {
	a          *Application
	selection  *gtk.FontSelection
	height     int
	lineheight int
	baseline   int
	ascent     int
	descent    int
}

func (a *Application) NewFont(height int, attributes map[string]string) driver.Fonter {
	f := &Font{
		a:      a,
		height: height,
	}
	family, style, variant, weight, stretch := fontspec.Attributes(attributes)
	ready := make(chan struct{})
	f.a.commands <- func() {
		f.selection = gtk.NewFontSelection(height, family, style, variant, weight, stretch, f.a.top)
		f.lineheight, f.baseline, f.ascent, f.descent = f.selection.Metrics()
		ready <- struct{}{}
	}
	<-ready
	return f
}

func (f *Font) LineHeight() int {
	return f.lineheight
}

func (f *Font) Baseline() int {
	return f.baseline
}

func (f *Font) Ascent() int {
	return f.ascent
}

func (f *Font) Descent() int {
	return f.descent
}

func (f *Font) Close() {
	f.selection.Free()
}

func (f *Font) Split(text string, edge int, indent int) []string {
	var lengths []int
	ready := make(chan struct{})
	f.a.commands <- func() {
		lengths = f.selection.Split(text, edge, indent)
		ready <- struct{}{}
	}
	<-ready
	return fontspec.SplitByLengths(text, lengths)
}

func (f *Font) Size(text string) image.Point {
	var output image.Point
	ready := make(chan struct{})
	f.a.commands <- func() {
		output.X, output.Y = f.selection.Size(text)
		ready <- struct{}{}
	}
	<-ready
	return output
}

func getFontSelection(f driver.Fonter) *gtk.FontSelection {
	for {
		wrappedFont, ok := f.(interface{ Unwrap() driver.Fonter })
		if !ok {
			break
		}
		f = wrappedFont.Unwrap()
	}
	if localFont, ok := f.(*Font); ok {
		return localFont.selection
	}
	return nil
}
