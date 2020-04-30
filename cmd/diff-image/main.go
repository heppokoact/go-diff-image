package main

import "C"

import (
	"image"
	"image/png"
	"log"
	"os"

	diffimage "../.."
)

func mustOpen(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func mustLoadImage(filename string) image.Image {
	f := mustOpen(filename)
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func mustSaveImage(img image.Image, output string) {
	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(f, img)
}

func rate16(c uint32, r float64) uint16 {
	fc := float64(c)
	return uint16(65535 - (65535-fc)*r)
}

func main() {
}

//export diffImage
func diffImage(imgPath1 string, imgPath2 string, outputPath string) {
	img1 := mustLoadImage(imgPath1)
	img2 := mustLoadImage(imgPath2)

	dst := diffimage.DiffImage(img1, img2)

	mustSaveImage(dst, outputPath)
}
