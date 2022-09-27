package clone

import (
	"go-image/util"
	"image"
	"testing"
)

func TestCloneAsRGBA(t *testing.T) {
	cases := []struct {
		desc     string
		value    image.Image
		expected *image.RGBA
	}{
		{
			desc: "RGBA",
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "RGBA64",
			value: &image.RGBA64{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "NRGBA",
			value: &image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0x80,
					0xFF, 0xFF, 0xFF, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "NRGBA64",
			value: &image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0x80,
					0xFF, 0xFF, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
	}

	for _, c := range cases {
		actual := CloneAsRGBA(c.value)

		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s : expected : %#v, actual : %#v", "CloneAsRGBA from "+c.desc, c.expected, actual)
		}
	}
}
