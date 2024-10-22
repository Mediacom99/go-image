package libgoimg

import "log"
import "image"

//FIXME
// there must be a better way
func LogFat(err error, msg string) {
	if err != nil {
		log.Fatal(msg + ": ", err)
	}
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
