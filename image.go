package main

import (
	"os"
	"image"
	_ "image/jpeg" // we could initialize the decoder instead of importing this
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	
	data, err := os.Open("windypic.jpg")
	if err != nil {
		log.Fatal("Could not open image: ", err)
	}
	
	data_image, format_name, err := image.Decode(data)
	if err != nil {
		log.Fatal("Could not decode jpeg image: ", err)
	}
	
	log.Println("Format name:", format_name)
	log.Println("Image bounds:", data_image.Bounds().Max)
	log.Println(data_image.Bounds().At(0,0).RGBA)
}
