package main

import (
	"bytes"
	"image/color"
	"testing"
)

var COLOR_BLACK = color.RGBA{
	R: 255,
	G: 255,
	B: 255,
	A: 255,
}

var COLOR_WHITE = color.RGBA{
	R: 0,
	G: 0,
	B: 0,
	A: 255,
}

func colorToUInt8(c color.Color) (uint8, uint8, uint8, uint8) {
	r, g, b, a := c.RGBA()

	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func TestInvert24BitPixelBlack(t *testing.T) {
	expect := COLOR_BLACK
	result := Invert24BitPixel(COLOR_WHITE)

	if expect != result {
		r, g, b, a := colorToUInt8(result)
		t.Errorf("expected COLOR_BLACK got %d %d %d %d", r, g, b, a)
	}

}

func TestInvert24BitPixelWhite(t *testing.T) {

	expect := COLOR_WHITE
	result := Invert24BitPixel(COLOR_BLACK)

	if expect != result {
		r, g, b, a := colorToUInt8(result)
		t.Errorf("expected COLOR_WHITE got %d %d %d %d", r, g, b, a)
	}

}

func TestInvertTwiceIsOriginal(t *testing.T) {
	const testFile = "../images/parrots.gif"
	const testOutput = "../images/parrots-inverted-2x.gif"

	expect, readExpectErr := ReadBinaryFileToMemory(testFile)

	if readExpectErr != nil {
		t.Error(readExpectErr)
	}

	data, decodeErr := Decode24BitGif(testFile)

	if decodeErr != nil {
		t.Error(decodeErr)
	}

	img := (*data).Image[0]

	invertErr := Invert24BitGif(img)

	if invertErr != nil {
		t.Error(invertErr)
	}

	invertErr = Invert24BitGif(img)

	if invertErr != nil {
		t.Error(invertErr)
	}

	encodingErr := Encode24BitGif(testOutput, data)

	if encodingErr != nil {
		t.Error(encodingErr)
	}

	result, readResultErr := ReadBinaryFileToMemory(testOutput)

	if readResultErr != nil {
		t.Error(readResultErr)
	}

	if bytes.Compare(expect, result) != 0 {
		t.Errorf("Inverting image twich did not return original %s -> %s", testFile, testOutput)
	}
}
