package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	// "image/color"
	libgi "goimg/libgoimg"
	"log"
	"os"
)

// For now it only handles jpeg images.
// TODO free main function from log setup, argument parsing and decoding from input
// so that eventually I can handle jpg, png and gif
func main() {
	
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	image_path, command := flagParser()

	data, open_err := os.Open(image_path)
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
	// TODO move all effects in the same big ApplyEffect function
	libgi.ModEachPixel(new_image, libgi.GetCommandToken(command))
	// libgi.MakeGrid(
	// 	new_image,
	// 	color.RGBA{123,32,189,45},
	// 	new_image.Bounds().Max.X/10,
	// )

	// Encode the image into jpeg and save it to a file
	encode_err := libgi.EncodeImageToJpeg(&new_image)
	libgi.LogFat(encode_err, "Error encoding image data into jpeg image") 
}

func flagParser() (string, string) {
	//TODO make sure that if there is not even one flag help is displayed
	// it's much more complicated than it needs to be
	if len(os.Args) < 2 {
		fmt.Print("No flags used. Please run `./goimg -h` to show available flags\n")
	}
	var (
		image_path string = "./images/windypic.jpg"
		command string = "greyscale"
	)
	cmd_ptr := flag.String("c", "greyscale", `Main command to run on the input image, available commands:
 greyscale
 inverted`)
	img_path_ptr := flag.String("if", "./images/windypic.jpg", "path for the input image")
	flag.Parse()
	image_path = *img_path_ptr
	command = *cmd_ptr
	log.Println("Image path flag:", image_path)
	log.Println("Command flag:", command)
	return image_path, command
}
