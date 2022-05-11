package main

import (
	"fmt"

	"github.com/221bye/img_to_ascii/convert"
)

func main() {
	opts := convert.Options{
		SrHeight: 5,
		SrWidth:  4,
		Filepath: "test_images/shrek.jpg",
	}
	res := convert.ConvertImage(opts)
	fmt.Println(res)
}
