package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
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

// For now it only handles jpeg images.
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	args := os.Args
	if(len(args) < 2) {
		fmt.Print("Not enough arguments\nUsage: ./image <image-name>.jpeg\n")
		os.Exit(1)
	}

	data, open_err := os.Open(args[1])
	logfat(open_err, "Could not open image:")
	defer data.Close()

	// JPEG decode
	data_image, decode_err := jpeg.Decode(data)
	logfat(decode_err, "Error decoding input image:")

	log.Println("Image max bounds:", data_image.Bounds().Max)
	log.Println("Image min bounds:", data_image.Bounds().Min)

	// Convert image to image.RGBA for modifying
	new_image := RedrawImageIntoRgba(data_image)

	// Apply some modification to the image
	ModEachPixel(new_image, Inverted)
	// MakeSomeSquares(new_image, color.RGBA{0,0,0,0}, 17)

	// Encode the image into jpeg and save it to a file
	encode_err := EncodeImageToJpeg(&new_image)
	logfat(encode_err, "Error encoding image data into jpeg image:")
}

// Encode image with jpeg format and save it to file
func EncodeImageToJpeg(image_data image.Image) error {
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
//FIXME this does not work if the origin of the image is NOT (0,0)
func RedrawImageIntoRgba(image_data image.Image) image.RGBA {
	new_image := image.NewRGBA(image_data.Bounds())
	draw.Draw(new_image, image_data.Bounds(), image_data, image.Point{}, draw.Over)

	//FIXME
	//make sure this is fine, should be because image.Image holds a pointer
	//to the actual data
	return *new_image;
}

// Colors the pixels with `color` for which x%mod == y%mod == 0
func MakeGrid(image_data image.RGBA, color color.Color, mod int) {
	image_bounds := image_data.Bounds()
	for y := image_bounds.Min.Y; y < image_bounds.Max.Y; y++ {
		for x := image_bounds.Min.X; x < image_bounds.Max.X; x++ {
			if x%mod == 0 || y%mod == 0 {
				image_data.Set(x,y, color)
			}
		}
	}
}

type ImageMod int
const (
	Grayscale ImageMod = iota // 0
	Inverted                  // 1
)

// Modifies the image with a given mod chosen from the const vars defined above
// FIXME Should take the color model as input that I get from parsing the command line
// and move the switch logic in the parsing of the command line args
func ModEachPixel(image_data image.RGBA, mod ImageMod) {
	image_bounds := image_data.Bounds()
	for y := image_bounds.Min.Y; y < image_bounds.Max.Y; y++ {
		for x := image_bounds.Min.X; x < image_bounds.Max.X; x++ {
			old_color := image_data.At(x,y)
			var new_color color.Color

			switch mod {
			case Grayscale:
				new_color = color.GrayModel.Convert(old_color)
			case Inverted:
				new_color = InvertedModel.Convert(old_color)
			default:
				return;
			}

			image_data.Set(x,y,new_color)
		}
	}
}

var (
	InvertedModel color.Model = color.ModelFunc(invertedModel)
)

func invertedModel(old_color color.Color) color.Color {
	r,g,b,a := old_color.RGBA()
	return color.RGBA {
		R: uint8(255 - r >> 8),
		G: uint8(255 - g >> 8),
		B: uint8(255 - b >> 8),
		A: uint8(a>>8),
	}
}
