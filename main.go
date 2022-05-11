package main

import (
	"flag"
	"fmt"
	"github.com/221bye/img_to_ascii/convert"
)

func main() {
	var fFlag = flag.String("f", "test_images/shrek.jpg", "path to image you want to convert")
	var srWFlag = flag.Int("srW", 4, "width of rectangle that one symbol represents")
	var srHFlag = flag.Int("srH", 5, "height of rectangle that one symbol represents")

	flag.Parse()

	opts := convert.Options{
		SrHeight: *srHFlag,
		SrWidth:  *srWFlag,
		Filepath: *fFlag,
	}
	res := convert.ConvertImage(opts)
	fmt.Println(res)
}
