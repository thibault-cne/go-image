package main

import (
	"go-image/effects"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	f, _ := os.Open("gopher_and_docker.jpeg")
	defer f.Close()

	img, _, _ := image.Decode(f)

	dst := effects.Invert(img, 0, 200, 0, 200)

	nf, _ := os.Create("partialInvert.jpeg")
	jpeg.Encode(nf, dst, &jpeg.Options{Quality: 95})
}
