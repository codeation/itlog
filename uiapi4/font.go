package uiapi4

import (
	"log"

	"github.com/codeation/itlog/gtk4"
)

type font struct {
	id        int
	height    int
	selection *gtk4.FontSelection
}

func (u *uiAPI) FontNew(fontID int, height int, style, variant, weight, stretch int, family string) (int, int, int, int) {
	f := &font{
		id:        fontID,
		height:    height,
		selection: gtk4.NewFontSelection(height, family, style, variant, weight, stretch),
	}
	u.fonts[fontID] = f

	return f.selection.Metrics(u.top)
}

func (u *uiAPI) FontDrop(fontID int) {
	f, ok := u.fonts[fontID]
	if !ok {
		log.Printf("FontDrop: font not found: %d", fontID)
		return
	}

	f.selection.Free()
	delete(u.fonts, fontID)
}

func (u *uiAPI) FontSplit(fontID int, text string, edge, indent int) []int {
	f, ok := u.fonts[fontID]
	if !ok {
		log.Printf("FontSplit: font not found: %d", fontID)
		return nil
	}

	return f.selection.Split(u.top, text, edge, indent)
}

func (u *uiAPI) FontSize(fontID int, text string) (int, int) {
	f, ok := u.fonts[fontID]
	if !ok {
		log.Printf("FontSize: font not found: %d", fontID)
		return 0, 0
	}

	return f.selection.Size(u.top, text)
}
