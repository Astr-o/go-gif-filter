package invert

import (
	"bufio"
	"errors"
	"image"
	"image/color"
	"image/gif"
	"os"
)

// ReadBinaryFileToMemory  return raw byte array of file in memory
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

// Decode24BitGif return ptr to gif.GIF object decoded from file at filename
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

// Encode24BitGif encode and save content of img to file at filename
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

// Invert24BitGif apply a per pixel inverse filter
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

// Invert24BitPixel invert a single pixel using conversion to RGBA
func Invert24BitPixel(c color.Color) color.Color {

	const ColorMaxVal uint8 = 255

	r, g, b, a := c.RGBA()

	newC := color.RGBA{
		R: ColorMaxVal - uint8(r>>8),
		G: ColorMaxVal - uint8(g>>8),
		B: ColorMaxVal - uint8(b>>8),
		A: uint8(a >> 8),
	}

	return newC
}
