package kolor

import (
	"image/color"
	"math"
)

type Color struct {
	R, G, B float64
}

// Default alpha value
var alpha = 65535.0

// D65 default value
var D65 = [3]float64{95.0489, 100.0, 108.8840}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R*alpha + 0.5)
	g = uint32(c.G*alpha + 0.5)
	b = uint32(c.B*alpha + 0.5)
	a = uint32(alpha)
	return
}

// color.RGBA is alpha-premultiplied so we need to divide each color by alpha.
func MakeKolor(c color.Color) (Color, bool) {
	r, g, b, a := c.RGBA()

	if a == 0 {
		return Color{0, 0, 0}, false
	}

	r *= 0xffff
	r /= a
	g *= 0xffff
	g /= a
	b *= 0xffff
	b /= a

	return Color{float64(r) / alpha, float64(g) / alpha, float64(b) / alpha}, true
}

func sq(v float64) float64 {
	return v * v
}

// LinearRGB converts the color into the linear RGB space (see https://www.sjbrown.co.uk/posts/gamma-correct-rendering/)
func (c Color) LinearRGB() (r, g, b float64) {
	r = linearize(c.R)
	g = linearize(c.G)
	b = linearize(c.B)
	return
}

// sRGB to linear RGB
func linearize(v float64) float64 {
	if v <= 0.04045 {
		return v / 12.92
	}

	return math.Pow(((v + 0.055) / 1.055), 2.4)
}

// LinearRGB create a RGB color out of the given linear RGB color (see https://www.sjbrown.co.uk/posts/gamma-correct-rendering/)
func LinearRGB(r, g, b float64) Color {
	return Color{delinearize(r), delinearize(g), delinearize(b)}
}

// linear RGB to sRGB
func delinearize(v float64) float64 {
	if v <= 0.0031308 {
		return 12.92 * v
	}

	return 1.055*math.Pow(v, 1.0/2.4) - 0.055
}

// XYZ //

// Convert the color to the XYZ fomat
func (c Color) XYZ() (x, y, z float64) {
	return LinearRGBToXYZ(c.LinearRGB())
}

// Convert linear RGB to XYZ with D65 white point (see https://www.sjbrown.co.uk/posts/gamma-correct-rendering/)
func LinearRGBToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.4124*r + 0.3576*g + 0.1805*b
	y = 0.2126*r + 0.7152*g + 0.0722*b
	z = 0.0193*r + 0.1192*g + 0.9505*b
	return
}

// L*a*b* //

// L*a*b* func
func lab_f(t float64) float64 {
	var s = 6.0 / 29.0

	if t > math.Pow(s, 3) {
		return math.Pow(t, 1.0/3.0)
	}

	return t/(3*math.Pow(s, 2)) + (4.0 / 29.0)
}

// Convert XYZ to L*a*b* with D65 white ref
func XYZToLab(x, y, z float64) (l, a, b float64) {
	return XYZToLabWhiteRef(x, y, z, D65)
}

// Convert XYZ to L*a*b* with white ref (see https://en.wikipedia.org/wiki/CIELAB_color_space#From_CIEXYZ_to_CIELAB)
func XYZToLabWhiteRef(x, y, z float64, wref [3]float64) (l, a, b float64) {
	fy := lab_f(y / wref[1])
	l = 116*fy - 16
	a = 500 * (lab_f(x/wref[0]) - fy)
	b = 200 * (fy - lab_f(z/wref[2]))
	return
}

// Converts the given color to CIE L*a*b* space using D65 as reference white.
func (col Color) Lab() (l, a, b float64) {
	return XYZToLab(col.XYZ())
}

// L*a*b* distance is a goode mesurement
func (c1 Color) DistanceLAB(c2 Color) float64 {
	l1, a1, b1 := c1.Lab()
	l2, a2, b2 := c2.Lab()
	return math.Sqrt(sq(l1-l2) + sq(a1-a2) + sq(b1-b2))
}

// HSL //

// HSL convertion format
func (c Color) HSL() (h, s, l float64) {
	min := math.Min(math.Min(c.R, c.G), c.B)
	max := math.Max(math.Max(c.R, c.G), c.B)

	l = (max + min) / 2

	if min == max {
		s = 0
		h = 0
	} else {
		if l < 0.5 {
			s = (max - min) / (max + min)
		} else {
			s = (max - min) / (2.0 - max - min)
		}

		if max == c.R {
			h = (c.G - c.B) / (max - min)
		} else if max == c.G {
			h = 2.0 + (c.B-c.R)/(max-min)
		} else {
			h = 4.0 + (c.R-c.G)/(max-min)
		}

		h *= 60

		if h < 0 {
			h += 360
		}
	}

	return
}
