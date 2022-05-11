package main

import (
	"fmt"
	"image"
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
	//ascii := "@#8&o:*. "
	ascii := " .*:o&8#@"

	avgIdx, avgIdy := 0, 0
	averages := make([][]int, grayImg.Rect.Max.Y/5)
	for i := 0; i < len(averages); i++ {
		averages[i] = make([]int, grayImg.Rect.Max.X/4)
	}

	//initialize slice for gray values of pixels
	grayValues := make([][]int, 5)
	for i := 0; i < 5; i++ {
		grayValues[i] = make([]int, 4)
	}

	for y := 0; y < grayImg.Rect.Max.Y/5; y++ {
		for x := 0; x < grayImg.Rect.Max.X/4; x++ {
			for ySub := 0; ySub < 5; ySub++ {
				for xSub := 0; xSub < 4; xSub++ {
					dx := xSub + (4 * x)
					dy := ySub + (5 * y)
					c := img.At(dx, dy)
					grayImg.Set(dx, dy, c)
					tmp := int(grayImg.GrayAt(dx, dy).Y)
					grayValues[ySub][xSub] = tmp
				}
			}
			//calculate average of gray values in slice
			sum := 0
			for i := 0; i < 5; i++ {
				for j := 0; j < 4; j++ {
					sum += grayValues[i][j]
				}
			}
			average := sum / (5 * 4)

			//append calculated average to slice
			averages[avgIdy][avgIdx] = average
			fmt.Println(avgIdx)
			avgIdx++
		}
		avgIdx = 0
		avgIdy++
	}

	for i := 0; i < len(averages); i++ {
		for j := 0; j < len(averages); j++ {
			fmt.Printf("%d ", averages[i][j])
		}
		fmt.Println()
	}

	result := ""
	for i := 0; i < len(averages); i++ {
		for j := 0; j < len(averages[i]); j++ {
			idx := averages[i][j] / 29
			result += string(ascii[idx])
		}
		result += "\n"
	}
	fmt.Println(result)
}
