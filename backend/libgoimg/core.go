package libgoimg

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

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

