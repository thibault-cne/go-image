package util

import (
	"errors"
	"image"
	"image/color"
	"math"
	"testing"
)

// Checks whether the relative error is below eps
func almosteq_eps(v1, v2, eps float64) bool {
	if math.Abs(v1) > delta {
		return math.Abs((v1-v2)/v1) < eps
	}
	return true
}

// Checks whether the relative error is below the 8bit RGB delta, which should be good enough.
const delta = 1.0 / 256.0

func almosteq(v1, v2 float64) bool {
	return almosteq_eps(v1, v2, delta)
}

func TestRank(t *testing.T) {
	cases := []struct {
		value    color.RGBA
		expected float64
	}{
		{
			value:    color.RGBA{0, 0, 0, 0},
			expected: 0,
		},
		{
			value:    color.RGBA{255, 255, 255, 255},
			expected: 255,
		},
		{
			value:    color.RGBA{130, 130, 130, 80},
			expected: 130,
		},
		{
			value:    color.RGBA{70, 230, 55, 255},
			expected: 164.5,
		},
	}

	for _, c := range cases {
		actual := Rank(c.value)

		if actual != c.expected {
			t.Errorf("%s : expected : %#v, actual : %#v", "Rank", c.expected, actual)
		}
	}
}

func TestRGBAImageEqual(t *testing.T) {
	cases := []struct {
		a        *image.RGBA
		b        *image.RGBA
		expected bool
	}{
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 3),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: false,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: false,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: true,
		},
		{
			a: &image.RGBA{},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: false,
		},
		{
			a:        &image.RGBA{},
			b:        &image.RGBA{},
			expected: true,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: true,
		},
	}

	for _, c := range cases {
		actual := RGBAImageEqual(c.a, c.b)

		if actual != c.expected {
			t.Errorf("%s : expected : %#v, actual : %#v. Images a : %#v b : %#v", "RGBAImageEqual", c.expected, actual, c.a, c.b)
		}
	}
}

func TestGrayImageEqual(t *testing.T) {
	cases := []struct {
		a        *image.Gray
		b        *image.Gray
		expected bool
	}{
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF,
				},
			},
			expected: false,
		},
		{
			a:        &image.Gray{},
			b:        &image.Gray{},
			expected: true,
		},
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			b:        &image.Gray{},
			expected: false,
		},
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			expected: true,
		},
	}

	for _, c := range cases {
		actual := GrayImageEqual(c.a, c.b)

		if actual != c.expected {
			t.Errorf("%s : expected : %#v, actual : %#v. Images a : %#v b : %#v", "GrayImageEqual", c.expected, actual, c.a, c.b)
		}
	}
}

func TestLuminance(t *testing.T) {
	cases := []struct {
		value    color.RGBA
		expected float64
	}{
		{
			value:    color.RGBA{0, 0, 0, 0},
			expected: 0,
		},
		{
			value:    color.RGBA{187, 187, 187, 0},
			expected: 187,
		},
	}

	for _, c := range cases {
		actual := Luminance(c.value)

		if actual != c.expected {
			t.Errorf("%s : expected : %#v, actual : %#v.", "Luminance", c.expected, actual)
		}
	}
}

func TestCheckError(t *testing.T) {
	cases := []struct {
		value error
		b     bool
	}{
		{
			value: nil,
			b:     false,
		},
		{
			value: errors.New("Test error"),
			b:     true,
		},
	}

	for _, c := range cases {
		func(b bool) {
			defer func() { recover() }()
			CheckError(c.value)
			if b {
				t.Errorf("should have panicked")
			}
		}(c.b)
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		x, y   int
		expect int
	}{
		{
			x:      3,
			y:      10,
			expect: 3,
		},
		{
			x:      3,
			y:      3,
			expect: 3,
		},
		{
			x:      -10,
			y:      10,
			expect: -10,
		},
	}

	for _, c := range cases {
		e := Min(c.x, c.y)

		if e != c.expect {
			t.Errorf("%s : expected : %#v, actual : %#v.", "Min", c.expect, e)
		}
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		x, y   int
		expect int
	}{
		{
			x:      3,
			y:      10,
			expect: 10,
		},
		{
			x:      3,
			y:      3,
			expect: 3,
		},
		{
			x:      -10,
			y:      10,
			expect: 10,
		},
	}

	for _, c := range cases {
		e := Max(c.x, c.y)

		if e != c.expect {
			t.Errorf("%s : expected : %#v, actual : %#v.", "Max", c.expect, e)
		}
	}
}

func TestGaussianDistribution(t *testing.T) {
	cases := []struct {
		x, y int
		s, e float64
	}{
		{
			x: 0,
			y: 0,
			s: 1.5,
			e: math.Exp(0.0) / (2 * math.Pi * 1.5 * 1.5),
		},
		{
			x: 1,
			y: -1,
			s: 1.5,
			e: math.Exp(-1.0/2.25) / (2 * math.Pi * 1.5 * 1.5),
		},
		{
			x: 2,
			y: -1,
			s: 1.5,
			e: math.Exp(-5.0/4.5) / (2 * math.Pi * 1.5 * 1.5),
		},
	}

	for _, c := range cases {
		e := GaussianDistribution(c.x, c.y, c.s)

		if !almosteq(c.e, e) {
			t.Errorf("%s : expected : %#v, actual : %#v.", "GaussianDistribution", c.e, e)
		}
	}
}

func TestTruncate(t *testing.T) {
	cases := []struct {
		x, e int
	}{
		{
			x: 200,
			e: 200,
		},
		{
			x: 0,
			e: 0,
		},
		{
			x: -1,
			e: 0,
		},
		{
			x: 260,
			e: 255,
		},
	}

	for _, c := range cases {
		e := Truncate(c.x)

		if e != c.e {
			t.Errorf("%s : expected : %#v, actual : %#v.", "Truncate", c.e, e)
		}
	}
}
