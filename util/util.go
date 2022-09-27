package util

import (
	"errors"
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

func BlurMatrix(n int) [][]float64 {
	if n%2 == 0 {
		panic(errors.New("blur matrix constructor is not an odd number"))
	}

	var (
		b [][]float64
		t float64
		m int
	)

	m = (n / 2)
	t = 1.0 / float64(n*n-1)

	for i := 0; i < n; i++ {
		r := make([]float64, n)

		for j := 0; j < n; j++ {
			if i == m && j == m {
				r[j] = 0
				continue
			}
			r[j] = t
		}

		b = append(b, r)
	}

	return b
}
