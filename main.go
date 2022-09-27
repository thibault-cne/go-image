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

	dst := effects.SobalEdge(img)

	nf, _ := os.Create("gopherEdge.png")
	png.Encode(nf, dst)
}
