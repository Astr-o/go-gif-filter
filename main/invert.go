package main

import (
	"bufio"
	"errors"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func ReadBinaryFileToMemory(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}

func Decode24BitGif(filename string) (*gif.GIF, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	bufr := bufio.NewReader(file)

	img, decodeErr := gif.DecodeAll(bufr)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return img, nil

}

func Encode24BitGif(filename string, img *gif.GIF) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	bufw := bufio.NewWriter(file)

	encodeErr := gif.EncodeAll(bufw, img)

	if encodeErr != nil {
		return encodeErr
	}

	return nil

}

func Invert24BitGif(img *image.Paletted) error {

	bounds := img.Bounds()

	if bounds.Min == bounds.Max {
		return errors.New("invalid bounds")
	}

	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {

			c := img.At(x, y)
			newC := Invert24BitPixel(c)
			img.Set(x, y, newC)

		}
	}

	return nil

}

func Invert24BitPixel(c color.Color) color.Color {

	const COLOR_MAX_VAL uint8 = 255

	r, g, b, a := c.RGBA()

	newC := color.RGBA{
		R: COLOR_MAX_VAL - uint8(r),
		G: COLOR_MAX_VAL - uint8(g),
		B: COLOR_MAX_VAL - uint8(b),
		A: uint8(a),
	}

	return newC
}
