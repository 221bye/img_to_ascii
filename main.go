package main

import (
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"os"
)

func main() {
	reader, err := os.Open("test_images/shrek.jpg")

	if err != nil {
		panic(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)

	if err != nil {
		panic(err)
	}

	grayImg := image.NewGray(img.Bounds())

	for y := 0; y < grayImg.Rect.Max.Y; y++ {
		for x := 0; x < grayImg.Rect.Max.X; x++ {
			c := img.At(x, y)
			grayImg.Set(x, y, c)
		}
	}

	outF, err := os.Create("test_images/output")
	defer outF.Close()

	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(outF, grayImg, &jpeg.Options{Quality: 90})

	if err != nil {
		panic(err)
	}
}
