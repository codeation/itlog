package impresslink

import (
	"image"
	"image/draw"

	"github.com/codeation/impress/driver"
	gtk "github.com/codeation/itlog/gtk4"
)

type Image struct {
	a      *Application
	bitmap *gtk.CairoBitmap
	size   image.Point
}

func (a *Application) NewImage(img image.Image) driver.Imager {
	pix, ok := img.(*image.NRGBA)
	if !ok {
		pix = image.NewNRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
		draw.Draw(pix, pix.Bounds(), img, image.Pt(0, 0), draw.Src)
	}
	i := &Image{
		a:    a,
		size: img.Bounds().Size(),
	}
	ready := make(chan struct{})
	i.a.commands <- func() {
		i.bitmap = gtk.NewCairoBitmap(pix.Pix, i.size.X, i.size.Y)
		ready <- struct{}{}
	}
	<-ready
	return i
}

func (i *Image) Size() image.Point { return i.size }

func (i *Image) Close() {
	i.a.commands <- func() {
		i.bitmap.Destroy()
	}
}

func getImageBitmap(i driver.Imager) *gtk.CairoBitmap {
	for {
		wrappedImage, ok := i.(interface{ Unwrap() driver.Imager })
		if !ok {
			break
		}
		i = wrappedImage.Unwrap()
	}
	if localImage, ok := i.(*Image); ok {
		return localImage.bitmap
	}
	return nil
}
