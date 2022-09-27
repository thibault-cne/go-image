package util

import (
	"image"
	"image/color"
)

func Rank(color color.RGBA) float64 {
	return float64(color.R)*0.3 + float64(color.G)*0.6 + float64(color.B)*0.1
}

func RGBAImageEqual(fImg, sImg *image.RGBA) bool {
	if !fImg.Rect.Eq(sImg.Rect) {
		return false
	}

	bounds := fImg.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			fPix := fImg.At(x, y).(color.RGBA)
			sPix := sImg.At(x, y).(color.RGBA)

			if fPix.R != sPix.R {
				return false
			}

			if fPix.G != sPix.G {
				return false
			}

			if fPix.B != sPix.B {
				return false
			}

			if fPix.A != sPix.A {
				return false
			}
		}

	}

	return true
}

func GrayImageEqual(fImg, sImg *image.Gray) bool {
	if !fImg.Rect.Eq(sImg.Rect) {
		return false
	}

	bounds := fImg.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			fPix := fImg.At(x, y).(color.Gray)
			sPix := sImg.At(x, y).(color.Gray)

			if fPix.Y != sPix.Y {
				return false
			}
		}

	}

	return true
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Luminance(c color.RGBA) float64 {
	r := 0.299
	g := 0.587
	b := 0.114

	return float64(r*float64(c.R) + g*float64(c.G) + b*float64(c.R))
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
