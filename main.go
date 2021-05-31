package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

//convert white to alpha channel in the given jpeg
func main() {

	//opening the image file
	images, _ := os.Open("img.jpeg")
	defer images.Close()

	//decoding the image to have all the pixel
	decodedImageJpeg, _ := jpeg.Decode(images)

	//rgba == 0,0,0,0
	img := image.NewRGBA(decodedImageJpeg.Bounds())

	//create a variable that cointains the size of the img
	size := img.Bounds().Size()

	//take every pixel and if is a pure white (65535, 65535, 65535) turn into alpha channel
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			r, g, b, a := decodedImageJpeg.At(x, y).RGBA()
			col := color.RGBA{}
			if r == 65535 && g == 65535 && b == 65535 {
				col = color.RGBA{0, 0, 0, 0}
			} else {
				//turn every r, g, b, a from 32bit into 8bit THIS STRUCT ACCEPT ONLY 8BIT INT
				col = color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			}
			//set the color converted in the image
			img.Set(x, y, col)
		}
	}

	//crete the image empty
	imgPng, _ := os.Create("test.png")
	defer imgPng.Close()

	//decoding the RGBA image into the empty one
	png.Encode(imgPng, img)
}
