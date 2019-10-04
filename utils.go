package colorutils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func hexToDecimal(hex string) (decimal uint8, err error) {
	// Ensure our string is long enough to avoid panics
	if len(hex) < 2 {
		err = fmt.Errorf("invalid hex string, expected a length of 2 and received a length of %d", len(hex))
		return
	}

	var u64 uint64
	// Parse a maximum of two characters from our inbound hex string to uint with a bitSize of 8 from base16
	if u64, err = strconv.ParseUint(hex[:2], 16, 8); err != nil {
		return
	}

	// Cast uint64 to uint8
	decimal = uint8(u64)
	return
}

func decimalToHex(decimal uint8) (hex string) {
	// Cast as uint64 so that we can pass the value to the formatting func
	u64 := uint64(decimal)
	// Format the decimal value (base10) as a hex value (base16) and return
	return strconv.FormatUint(u64, 16)
}

func sanitizeHex(hex string) (sanitized string, err error) {
	// Remove pound prefix (if exists)
	hex = strings.Replace(hex, "#", "", 1)

	// Switch on string length
	switch len(hex) {
	case 6:
		// We already have the proper length, no need to make any changes
		// Set the output as the input hex and return
		sanitized = hex
		return

	case 3:
		// A shorthand hex string has been provided, convert it to an expanded hex string
		sanitized = expandHex(hex)
		return

	default:
		err = fmt.Errorf("invalid hex, \"%s\" has a string length of %d - which is not a supported", hex, err)
		return
	}
}

func expandHex(hex string) (expanded string) {
	// Pre-allocate buffer to match the expanded string length
	buf := make([]rune, 0, 6)
	// Iterate through all the runes which make up the inbound hex string
	for _, r := range hex {
		// Append the rune to our buffer
		buf = append(buf, r, r)
	}

	// Convert buffer to string and return
	return string(buf)
}

func applyLuminance(value uint8, luminance float64) uint8 {
	// Convert value to float64
	f64 := float64(value)

	// Add the value multiplied by the luminance to our float value
	f64 += f64 * luminance

	// Round the float value
	f64 = math.Round(f64)

	// Ensure our value hasn't exceeded the range in either direction
	switch {
	case f64 < 0:
		// Float value is less than 0, set to zero
		f64 = 0
	case f64 > 255:
		// Float value is greater than 255, set to 255
		f64 = 255
	}

	// Convert float value to uint8 and return
	return uint8(f64)
}
