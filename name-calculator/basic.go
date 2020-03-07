package namecalculator

import (
	"fmt"
	"image"
)

// BasicNameCalculator is a struct that given an image will rename the image
type BasicNameCalculator struct{}

func (bnc *BasicNameCalculator) Rename(img image.Image) (string, error) {
	// jpeg (and possibly others) use different color models, convert them all to RGBA
	rgba := &image.RGBA{}
	pixelColor := rgba.ColorModel().Convert(img.At(0, 0))

	red, green, blue, _ := pixelColor.RGBA()
	// TODO: should actually add all the RGB in their own struct to see which color has the largest number maybe?
	redDominant := 0
	blueDominant := 0
	greenDominant := 0

	if red >= green && red >= blue {
		redDominant += 1
	} else if green >= red && green >= blue {
		greenDominant += 1
	} else {
		blueDominant += 1
	}

	// var pixelColor color.Color
	// if img.ColorModel() != color.RGBAModel {
	// 	rgba := &image.RGBA{}
	// 	pixelColor = rgba.ColorModel().Convert(img.At(0, 0))
	// } else {
	// 	pixelColor = img.At(0, 0)
	// }
	// fmt.Println(fmt.Sprintf("%+v", imag.ColorModel().Convert(img.ColorModel())))
	fmt.Println(fmt.Sprintf("%+v", pixelColor))
	return "", nil
}
