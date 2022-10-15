package kolor

import (
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

func TestMakeKolor(t *testing.T) {
	cases := []struct {
		c color.Color
		e Color
		s bool
	}{
		{
			c: color.RGBA{R: 255, G: 255, B: 255, A: 0},
			e: Color{0, 0, 0},
			s: false,
		},
		{
			c: color.RGBA{R: 255, G: 255, B: 255, A: 1},
			e: Color{255, 255, 255},
			s: true,
		},
		{
			c: color.RGBA{R: 255, G: 255, B: 255, A: 2},
			e: Color{255.0 / 2.0, 255.0 / 2.0, 255.0 / 2.0},
			s: true,
		},
		{
			c: color.RGBA{R: 127, G: 0, B: 255, A: 2},
			e: Color{127.0 / 2.0, 0.0 / 2.0, 255.0 / 2.0},
			s: true,
		},
	}

	for _, c := range cases {
		e, s := MakeKolor(c.c)

		if !almosteq(e.R, c.e.R) || !almosteq(e.G, c.e.G) || !almosteq(e.B, c.e.B) || s != c.s {
			t.Errorf("%s : expected : %#v, actual : %#v.", "MakeKolor", c.e, e)
		}
	}
}

func TestSq(t *testing.T) {
	cases := []struct {
		v, e float64
	}{
		{
			v: 1.0,
			e: 1.0,
		},
		{
			v: -1.0,
			e: 1.0,
		},
		{
			v: 10.0,
			e: 100.0,
		},
	}

	for _, c := range cases {
		e := sq(c.v)

		if e != c.e {
			t.Errorf("%s : expected : %#v, actual : %#v.", "MakeKolor", c.e, e)
		}
	}
}
