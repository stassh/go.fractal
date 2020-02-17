package main

import (
	"flag"
	"fractal/cmdflags"
)

const Dst = "./out"

func main() {
	flag.BoolVar(&cmdflags.ShowBuildInfo, "version", false,
		"Shows build inforation and exists.")

	flag.IntVar(&cmdflags.HttpPort, "port", 8080,
		"Specify HTTP port.")

	flag.Parse()

	// generate some QR code look a like image
	imageMaxX := 800
	imageMaxY := 800

	b := Borders{X1: -2.5, X2: 2.5, Y1: -2.5, Y2: 2.5}

	img := GenerateFractalImage(imageMaxX, imageMaxY, b)

	// ok, write out the data into the new PNG file
	WriteToFile(img, "./output1.png")

	handleRequests()
}


