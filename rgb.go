package colorutils

// NewRGB will return a new RGB value
func NewRGB(red, green, blue uint8) (r RGB) {
	r.Red = red
	r.Green = green
	r.Blue = blue
	return
}

// NewRGBFromHex will return a new RGB value from a HEX string
func NewRGBFromHex(hex string) (r RGB, err error) {
	if hex, err = sanitizeHex(hex); err != nil {
		return
	}

	if r.Red, err = hexToDecimal(hex[:2]); err != nil {
		return
	}

	if r.Green, err = hexToDecimal(hex[2:4]); err != nil {
		return
	}

	if r.Blue, err = hexToDecimal(hex[4:]); err != nil {
		return
	}

	return
}

// RGB represents an RGB defined color
type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

// ApplyLuminance will change the luminance of the selected RGB value
// Luminance value should range between -1 (black) and 1 (white)
func (r *RGB) ApplyLuminance(luminance float64) {
	r.Red = applyLuminance(r.Red, luminance)
	r.Green = applyLuminance(r.Green, luminance)
	r.Blue = applyLuminance(r.Blue, luminance)
}

// Hex will return a Hex value from an RGB value
func (r *RGB) Hex() (hex string) {
	red := decimalToHex(r.Red)
	green := decimalToHex(r.Green)
	blue := decimalToHex(r.Blue)
	return red + green + blue
}
