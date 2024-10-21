package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/draw"
	"log"
	"os"
)

//FIXME
//there must be a better way
func logfat(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

// We deal with jpeg, make a function to turn png into jpeg
func convertJpegToPng() error {
	return nil
}

// Should check and initialize properly for jpeg and png
// considering for jpeg I have to create a new image of type RGBA
// and redraw the original that was an image.YCbCr
// For now I only handle jpeg images. I will add a function to convert
// a png to a jpeg image.

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	args := os.Args
	if(len(args) < 2) {
		fmt.Print("Not enough arguments\nUsage: ./image <image-name>.png/jpeg\n")
		os.Exit(1)
	}
	
	data, open_err := os.Open(args[1])
	logfat(open_err, "Could not open image:")
	defer data.Close()
	
	// JPEG decode
	data_image, decode_err := jpeg.Decode(data)
	logfat(decode_err, "Error decoding input image:")
	
	log.Println("Image bounds:", data_image.Bounds().Max)
	log.Printf("%T\n", data_image)

	new_image := RedrawImageIntoRgba(data_image)

	encode_err := EncodeRgbaToJpeg(&new_image)
	logfat(encode_err, "Error encoding image data into jpeg image:")
}

func EncodeRgbaToJpeg(image_data image.Image) error {	
	//Create new file to hold new image
	outfile, open_err := os.Create("NEWIMAGE.jpeg")
	if open_err != nil { return open_err }

	// JPEG encode
	encode_err := jpeg.Encode(outfile, image_data, nil)
	if encode_err != nil { return encode_err }
	
	defer outfile.Close()
	return nil
}

// Prints every pixel
func PrintPixels(image_data image.Image) {
	image_bounds := image_data.Bounds()
	for y := image_bounds.Min.Y; y < image_bounds.Max.Y; y++ {
		for x := image_bounds.Min.X; x < image_bounds.Max.X; x++ {
			r, g, b, a := image_data.At(x, y).RGBA()
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

//Redraws an image into a image.RGBA that is mutable
func RedrawImageIntoRgba(image_data image.Image) image.RGBA {
	new_image := image.NewRGBA(image_data.Bounds())
	draw.Draw(new_image, image_data.Bounds(), image_data, image.Point{}, draw.Over)
	
	//FIXME
	//make sure this is fine, should be because image.Image holds a pointer
	//to the actual data
	return *new_image;
}
