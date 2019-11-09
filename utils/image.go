package utils

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	_ "image/png"
	"io"
)

func Join(dist io.Writer, src []io.Reader, height int) error {
	sumWidth, imgs := 0, []image.Image{}
	for _, s := range src {
		img, err := imaging.Decode(s)
		if err != nil {
			return err
		}
		img = imaging.Resize(img, 0, height, imaging.Lanczos)
		imgs = append(imgs, img)
		sumWidth += img.Bounds().Size().X
	}
	base := imaging.New(sumWidth, height, color.Transparent)
	dw := 0
	for _, img := range imgs {
		base = imaging.Paste(base, img, image.Pt(dw, 0))
		dw += img.Bounds().Size().X
	}
	if err := imaging.Encode(dist, base, imaging.PNG); err != nil {
		return err
	}
	return nil
}