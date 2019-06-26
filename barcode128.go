package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func main() {
	// Create the barcode
	bcImg, _ := code128.Encode("12345")

	// Scale the barcode to 200x200 pixels
	bcImg, _ = barcode.Scale(bcImg, 200, 200)

	// create the output file
	file, _ := os.Create("128code.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, bcImg)
}
