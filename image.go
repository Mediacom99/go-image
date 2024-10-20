package main

import (
	"os"
	"image"
	_ "image/jpeg" // we could initialize the decoder instead of importing this
	"log"
)

func main() {
	data, _ := os.Open("windypic.jpg")
	data_image, format_name, _ := image.Decode(data)
	log.Println("Format name:", format_name)
	log.Println("Image bounds:", data_image.Bounds().Max)
	log.Println(data_image.Bounds().At(0,0).RGBA)
}
