package main

import (
	"go-image/effects"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	f, _ := os.Open("./assets/images/gopher_and_docker.jpeg")
	defer f.Close()

	img, _, _ := image.Decode(f)

	dst := effects.ColorFilter(img, color.RGBA{255, 0, 0, 255}, 6.15)

	nf, _ := os.Create("ColorFilter.jpeg")
	jpeg.Encode(nf, dst, &jpeg.Options{Quality: 95})
}
