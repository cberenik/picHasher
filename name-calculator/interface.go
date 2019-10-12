package namecalculator

import "image"

type Interface interface {
	Rename(image.Image) (string, error)
}
