package libgoimg

import "image"
import "image/color"


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

// Simple parser that converts string command to ImageMod token.
// The token should be used as input for libgi.ModEachPixel in order
func GetCommandToken(cmd string) ImageMod {
	switch cmd {
	case "greyscale":
		return Grayscale
	case "inverted":
		return Inverted
	default: //TODO handle this better
		return Grayscale
	}
}

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
