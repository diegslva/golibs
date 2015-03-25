// +build local

package main

import (
	"github.com/simonwaldherr/golibs/graphics"
	"github.com/simonwaldherr/golibs/ansi"
	"image/jpeg"
	"log"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./flowers.jpg")
	defer file.Close()

	if err != nil {
		log.Printf("EachPixel Test failed: %v", err)
	}
	img := graphics.EachPixel(file, func(r, g, b, a uint8) (uint8, uint8, uint8, uint8) {
		fmt.Printf("r: %v\tg: %v\tb: %v\n", ansi.Color(r, ansi.Red), ansi.Color(g, ansi.Green), ansi.Color(b, ansi.Blue))
		return uint8(255-r), uint8(255-g), uint8(255-b), a
	})
	fd, err := os.Create("./inv.jpg")
	if err != nil {
		log.Printf("EachPixel Test failed: %v", err)
	}

	err = jpeg.Encode(fd, img, nil)
	if err != nil {
		log.Printf("EachPixel Test failed: %v", err)
	}

	err = fd.Close()
	if err != nil {
		log.Printf("EachPixel Test failed: %v", err)
	}
}