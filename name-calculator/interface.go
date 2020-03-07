package namecalculator

import "image"

// Interface is the interface for determining new names for images
type Interface interface {
	Rename(image.Image) (string, error)
}
