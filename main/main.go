package main

import (
	"bufio"
	"fmt"
	"os"
)

const testFile = "./images/parrots.gif"
const testOutput = "./images/parrots-inverted.gif"

func main() {
	fmt.Println("Hello, world.")

	buff, readFileErr := ReadBinaryFileToMemory(testFile)

	if readFileErr != nil {
		panic(readFileErr)
	}

	fmt.Println("=== Buffer INT ===")
	fmt.Println(buff)

	fmt.Println("=== Buffer String ===")
	fmt.Println(string(buff))

	fmt.Printf("Decoding - %s \n", testFile)
	data, decodeErr := Decode24BitGif(testFile)

	if decodeErr != nil {
		panic(decodeErr)
	}

	img := (*data).Image[0]

	invertErr := Invert24BitGif(img)

	if invertErr != nil {
		panic(invertErr)
	}

	fmt.Printf("Encode - %s \n", testOutput)
	encodingErr := Encode24BitGif(testOutput, data)

	if encodingErr != nil {
		panic(invertErr)
	}

	fmt.Println("Press Enter to Exit")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

}
