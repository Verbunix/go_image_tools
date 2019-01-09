package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error. Missing required args.")
		fmt.Println(os.Args[0] + " [filename_image] [filename_new_blur]")
		return
	}

	filenameSource := os.Args[1]
	filenameNew := os.Args[2]

	img := gocv.IMRead(filenameSource, gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Error. Empty source image.")
		return
	}
	defer img.Close()

	imgWidth := img.Cols()
	imgHeight := img.Rows()
	pointX := float64(imgWidth) * 0.04
	pointY := float64(imgHeight) * 0.04
	gocv.Blur(img, &img, image.Point{X: int(pointX), Y: int(pointY)})
	gocv.IMWrite(filenameNew, img)
}
