package main

import (
	"go-image/effects"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	f, _ := os.Open("./assets/images/gopher_and_docker.jpeg")
	defer f.Close()

	img, _, _ := image.Decode(f)

	dst := effects.Threshold(img, 187)

	nf, _ := os.Create("halfThreshold.jpeg")
	jpeg.Encode(nf, dst, &jpeg.Options{Quality: 95})
}
