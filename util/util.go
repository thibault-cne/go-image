package util

import (
	"image"
	"image/color"
	"math"
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

func GaussianBlurMatrix(n int) [][]float64 {
	var (
		b [][]float64
		s float64
	)

	for i := -n; i < n+1; i++ {
		var r []float64

		for j := -n; j < n+1; j++ {
			t := GaussianDistribution(i, j, 1.5)
			r = append(r, t)
			s += t
		}

		b = append(b, r)
	}

	for i := 0; i < n*2+1; i++ {
		for j := 0; j < n*2+1; j++ {
			b[i][j] /= s
		}
	}

	return b
}

func GaussianDistribution(x, y int, s float64) float64 {
	var (
		e  float64
		f  float64
		s2 float64
	)

	s2 = s * s
	f = 1 / (2 * math.Pi * s2)
	e = float64(x*x+y*y) / (2 * s2)

	return f * math.Exp(-e)
}
