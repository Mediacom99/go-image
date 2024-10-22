package main

import (
	"fmt"
	"image"
	"image/jpeg"
	// "image/color"
	"log"
	"os"
	libgi "goimg/libgoimg"
)

// For now it only handles jpeg images.
// TODO free main function from log setup, argument parsing and decoding from input
// so that eventually I can handle jpg, png and gif
func main() {
	
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	args := os.Args
	if(len(args) < 2) {
		fmt.Print("Not enough arguments\nUsage: ./image <image-name>.jpeg\n")
		os.Exit(1)
	}

	data, open_err := os.Open(args[1])
	libgi.LogFat(open_err, "Could not open image")
	defer data.Close()

	// JPEG decode
	data_image, decode_err := jpeg.Decode(data)
	libgi.LogFat(decode_err, "Error decoding input image")

	log.Println("Image max bounds:", data_image.Bounds().Max)
	log.Println("Image min bounds:", data_image.Bounds().Min)

	// Convert image to image.RGBA for modifying
	var new_image image.RGBA = libgi.RedrawImageIntoRgba(data_image)

	// Apply some modification to the image
	libgi.ModEachPixel(new_image, libgi.Inverted)
	// libgi.MakeGrid(
	// 	new_image,
	// 	color.RGBA{123,32,189,45},
	// 	new_image.Bounds().Max.X/10,
	// )

	// Encode the image into jpeg and save it to a file
	encode_err := libgi.EncodeImageToJpeg(&new_image)
	libgi.LogFat(encode_err, "Error encoding image data into jpeg image") 
}
