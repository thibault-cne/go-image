package effects

import (
	"go-image/clone"
	"go-image/util"
	"image"
	"image/color"
	"math"
)

func Grayscale(img image.Image, config ...float64) *image.RGBA {
	// Set the weight for the grayscale luma method. By default it respects the ITU-R recommendations
	r := 0.2126
	g := 0.7152
	b := 0.0722

	// We verify that the sum of weight is equal to 1 else we apply the default values.
	if len(config) == 3 && float32(config[0])+float32(config[1])+float32(config[2]) == 1 {
		r = config[0]
		g = config[1]
		b = config[2]
	}

	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(bounds)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)
			gray := uint8(float64(pix.R)*r + float64(pix.G)*g + float64(pix.B)*b + 0.5)
			dst.Set(x, y, color.RGBA{gray, gray, gray, pix.A})
		}
	}

	return dst
}

func Threshold(img image.Image, level uint8) *image.Gray {
	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewGray(bounds)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)

			r := util.Rank(pix)

			if uint8(r) >= level || pix.R == 0 && pix.G == 0 && pix.B == 0 && pix.A == 0 {
				dst.Set(x, y, color.White)
			} else {
				dst.Set(x, y, color.Black)
			}
		}
	}

	return dst
}

// The x0, y0, x1, y1 values must be inside the picture
func Invert(img image.Image, x0, y0, x1, y1 int) *image.RGBA {
	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(bounds)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pix := src.At(x, y).(color.RGBA)

			var c color.RGBA

			if x >= x0 && x < x1 && y >= y0 && y < y1 {
				c = color.RGBA{uint8(255 - pix.R), uint8(255 - pix.G), uint8(255 - pix.B), pix.A}
			} else {
				c = color.RGBA{pix.R, pix.G, pix.B, pix.A}
			}

			dst.Set(x, y, c)
		}
	}

	return dst
}

func SobalEdge(img image.Image) *image.Gray {
	src := clone.CloneAsRGBA(img)
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	graySrc := Grayscale(src)

	dst := image.NewGray(bounds)

	hFilter := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	vFilter := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			gradient := [3][3]int{}

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					gradient[i][j] = int(util.Luminance(graySrc.RGBAAt(x-1+i, y-1+j)))
				}
			}

			gradientX := 0
			gradientY := 0

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					gradientX += gradient[i][j] * hFilter[i][j]
					gradientY += gradient[i][j] * vFilter[i][j]
				}
			}

			c := int(math.Sqrt(float64(gradientX*gradientX + gradientY*gradientY)))

			dst.SetGray(x, y, color.Gray{Y: uint8(c)})
		}
	}

	return dst
}
