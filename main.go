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

	dst := effects.SobalEdge(img)

	nf, _ := os.Create("SobalEdge.jpeg")
	jpeg.Encode(nf, dst, &jpeg.Options{Quality: 95})
}
