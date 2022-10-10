package filter

import (
	"image"
	"image/color"

	"go-image/clone"
	"go-image/kolor"
)

type Filter struct {
	Color *color.RGBA
}

// We only take colors greater than the filter color part.
func (f *Filter) Apply(img image.Image) *image.RGBA {
	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(bounds)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pix := src.At(x, y).(color.RGBA)
			newP := color.RGBA{pix.R, pix.G, pix.B, pix.A}

			if f.Color.R > pix.R {
				newP.R = 0
			}
			if f.Color.G > pix.G {
				newP.G = 0
			}
			if f.Color.B > pix.B {
				newP.B = 0
			}

			dst.Set(x, y, newP)
		}
	}

	return dst
}

// We keep only pix close to a certain color of a distance in the L*a*b* color representation.
// Other pixels are greyscaled
//
//	d : the maximal distance to the color you want to keep
func (f Filter) ColorFilter(img image.Image, d float64) *image.RGBA {
	// Set the weight for the grayscale luma method. By default it respects the ITU-R recommendations
	r := 0.2126
	g := 0.7152
	b := 0.0722

	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(bounds)

	k, _ := kolor.MakeKolor(f.Color)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pix := src.At(x, y).(color.RGBA)
			p, _ := kolor.MakeKolor(pix)

			if k.DistanceLAB(p) < d {
				dst.Set(x, y, color.RGBA{R: pix.R, G: pix.G, B: pix.B, A: pix.A})
			} else {
				gray := uint8(float64(pix.R)*r + float64(pix.G)*g + float64(pix.B)*b + 0.5)
				dst.Set(x, y, color.RGBA{gray, gray, gray, pix.A})
			}
		}
	}

	return dst
}
