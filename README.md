# Color utils
Color utils is a helper library to work with hex and rgb color values

## Usage

### NewRGB
```go
func ExampleNewRGB() {
	rgb := NewRGB(255, 100, 0)
	fmt.Printf("RGB value: %+v\n", rgb)
}
```

### NewRGBFromHex
```go
func ExampleNewRGBFromHex() {
	var (
		rgb RGB
		err error
	)

	if rgb, err = NewRGBFromHex("5c73da"); err != nil {
		log.Fatalf("unable to parse RGB value from Hex: %v", err)
	}

	fmt.Printf("RGB value: %+v\n", rgb)
}
```

### RGB.ApplyLuminance
```go
func ExampleRGB_ApplyLuminance() {
	rgb := NewRGB(92, 115, 218)
	rgb.ApplyLuminance(-0.5)
	fmt.Printf("New RGB value: %+v\n", rgb)
}
```

### RGB.Hex
```go
func ExampleRGB_Hex() {
	rgb := NewRGB(92, 115, 218)
	fmt.Printf("Hex value: %s\n", rgb.Hex())
}
```
