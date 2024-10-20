package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // we could initialize the decoder instead of importing this
	_ "image/png"
	"log"
	"os"
)

func printPixels(image image.Image) {
	image_bounds := image.Bounds()
	for y := image_bounds.Min.Y; y < image_bounds.Max.Y; y++ {
		for x := image_bounds.Min.X; x < image_bounds.Max.X; x++ {
			r, g, b, a := image.At(x, y).RGBA()
			log.Printf("Pixel at (%d, %d): R: %d, G: %d, B: %d, A: %d\n",
				x,
				y,
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
				uint8(a >> 8))
 		}
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	args := os.Args
	if(len(args) < 2) {
		fmt.Print("Not enough arguments\nUsage: ./image <image-name>.png/jpeg\n")
		os.Exit(1)
	}
	//TODO should check correct name and format like : <not too long name>.jpeg/png
	// even if os.Open checks for existence of file with given path
	
	data, err := os.Open(args[1])
	if err != nil {
		log.Fatal("Could not open image: ", err)
	}
	defer data.Close()
	
	data_image, format_name, err := image.Decode(data)
	if err != nil {
		log.Fatal("Could not decode jpeg image: ", err)
	}
	
	log.Println("Format name:", format_name)
	log.Println("Image bounds:", data_image.Bounds().Max)

	printPixels(data_image)
	// r, g, b, a := data_image.Bounds().At(639, 800).RGBA()
	// log.Println(r,g,b,a)
	
}
