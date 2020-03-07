package image_calculation

type ImageCalculator interface {
	CalculateDominantColour() (string, error)
}
