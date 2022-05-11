package convert

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

//srHeight - sub rect Height
//srWidth  - sub rect width
type Options struct {
	SrHeight int
	SrWidth  int
	Filepath string
}

func handleErr(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func ConvertImage(opt Options) string {
	reader, err := os.Open(opt.Filepath)
	handleErr(err, "Couldn't open file")
	defer reader.Close()

	img, _, err := image.Decode(reader)
	handleErr(err, "Couldn't decode image")

	averages := calculateAvgGreyValues(img, opt.SrHeight, opt.SrWidth)

	result := ""
	ascii := " .*:o&8#@"
	for i := 0; i < len(averages); i++ {
		for j := 0; j < len(averages[i]); j++ {
			idx := averages[i][j] / (255/(len(ascii)) + 1)
			result += string(ascii[idx])
		}
		result += "\n"
	}

	return result
}

// split image into rectangles with given width (w) and height(h)
// then calculate average grey value for every rectangle
func calculateAvgGreyValues(img image.Image, h, w int) [][]int {
	grayImg := image.NewGray(img.Bounds())
	maxX := grayImg.Rect.Max.X
	rectsX := maxX / w
	maxY := grayImg.Rect.Max.Y
	rectsY := maxY / h

	averages := make([][]int, rectsY)
	for i := 0; i < len(averages); i++ {
		averages[i] = make([]int, rectsX)
	}

	for y := 0; y < rectsY; y++ {
		for x := 0; x < rectsX; x++ {
			graySum := 0
			for srY := 0; srY < h; srY++ {
				for srX := 0; srX < w; srX++ {
					imgX := srX + (w * x)
					imgY := srY + (h * y)

					clr := img.At(imgX, imgY)
					grayImg.Set(imgX, imgY, clr)

					graySum += int(grayImg.GrayAt(imgX, imgY).Y)
				}
			}
			averageGray := graySum / (h * w)
			averages[y][x] = averageGray
		}
	}

	return averages
}
