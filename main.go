package main

import (
	"go-image/effects"
	"image"
	"image/png"
	"os"
)

func main() {
	f, _ := os.Open("gopher.png")
	defer f.Close()

	img, _, _ := image.Decode(f)

	dst := effects.Invert(img, 0, 200, 0, 200)

	nf, _ := os.Create("gopherPartialInvert.png")
	png.Encode(nf, dst)
}
