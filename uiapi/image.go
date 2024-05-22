package uiapi

import (
	"log"

	"github.com/codeation/itlog/gtk"
)

type image struct {
	bitmap *gtk.CairoBitmap
}

func (u *uiAPI) ImageNew(imageID int, width, height int, bitmap []byte) {
	u.images[imageID] = &image{
		bitmap: gtk.NewCairoBitmap(bitmap, width, height),
	}
}

func (u *uiAPI) ImageDrop(imageID int) {
	i, ok := u.images[imageID]
	if !ok {
		log.Printf("ImageDrop: image not found: %d", imageID)
		return
	}

	i.bitmap.Destroy()
	delete(u.images, imageID)
}
