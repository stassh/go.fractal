package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func WriteToFile(img image.Image, fileName string) {
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Generated image to %v \n", fileName)
}
