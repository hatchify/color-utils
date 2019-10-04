package colorutils

import (
	"fmt"
	"log"
	"testing"
)

func TestNewRGB(t *testing.T) {
	var rgb RGB
	rgb = NewRGB(255, 100, 0)

	if rgb.Red != 255 {
		t.Fatalf("invalid value, expected \"%+v\" but received \"%+v\"", 255, rgb.Red)
	}

	if rgb.Green != 100 {
		t.Fatalf("invalid value, expected \"%+v\" but received \"%+v\"", 100, rgb.Green)
	}

	if rgb.Blue != 0 {
		t.Fatalf("invalid value, expected \"%+v\" but received \"%+v\"", 0, rgb.Blue)
	}
}

func TestNewRGBFromHex(t *testing.T) {
	var (
		rgb RGB
		err error
	)

	if rgb, err = NewRGBFromHex("5c73da"); err != nil {
		t.Fatal(err)
	}

	if rgb.Red != 92 {
		t.Fatalf("invalid value, expected \"%+v\" but received \"%+v\"", 92, rgb.Red)
	}

	if rgb.Green != 115 {
		t.Fatalf("invalid value, expected \"%+v\" but received \"%+v\"", 115, rgb.Green)
	}

	if rgb.Blue != 218 {
		t.Fatalf("invalid value, expected \"%+v\" but received \"%+v\"", 218, rgb.Blue)
	}
}

func TestRGB_ApplyLuminance(t *testing.T) {
	var (
		rgb RGB
		hex string
	)

	rgb = NewRGB(92, 115, 218)
	rgb.ApplyLuminance(-0.5)

	hex = rgb.Hex()
	if hex != "2e3a6d" {
		t.Fatalf("invalid value, expected \"5c73da\" but received \"%+v\"", hex)
	}
}

func TestRGB_Hex(t *testing.T) {
	var (
		rgb RGB
		hex string
	)

	rgb = NewRGB(92, 115, 218)

	hex = rgb.Hex()
	if hex != "5c73da" {
		t.Fatalf("invalid value, expected \"5c73da\" but received \"%+v\"", hex)
	}
}

func ExampleNewRGB() {
	rgb := NewRGB(255, 100, 0)
	fmt.Printf("RGB value: %+v\n", rgb)
}

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

func ExampleRGB_ApplyLuminance() {
	rgb := NewRGB(92, 115, 218)
	rgb.ApplyLuminance(-0.5)
	fmt.Printf("New RGB value: %+v\n", rgb)
}

func ExampleRGB_Hex() {
	rgb := NewRGB(92, 115, 218)
	fmt.Printf("Hex value: %s\n", rgb.Hex())
}
