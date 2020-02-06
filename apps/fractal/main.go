package main

import (
	"flag"
	"fmt"
	"fractal/cmdflags"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

const Dst = "./out"

type borders struct {
	X1 float64
	X2 float64
	Y1 float64
	Y2 float64
}

type point struct {
	X float64
	Y float64
}

func main() {
	flag.BoolVar(&cmdflags.ShowBuildInfo, "version", false,
		"Shows build inforation and exists.")

	flag.IntVar(&cmdflags.HttpPort, "port", 8080,
		"Specify HTTP port.")

	flag.Parse()

	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// generate some QR code look a like image

	imageMaxX := 800
	imageMaxY := 800

	b := borders{X1: -2.5, X2: 2.5, Y1: -2.5, Y2: 2.5}

	imgRect := image.Rect(0, 0, imageMaxX, imageMaxY)
	img := image.NewRGBA(imgRect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for y := 0; y < imageMaxX; y += 1 {
		for x := 0; x < imageMaxY; x += 1 {
			variant := testPoint(b, x, y, imageMaxX, imageMaxY, 255)
			img.Set(x, y, &image.Uniform{rainbow(uint8(variant))})
		}
	}

	// ok, write out the data into the new PNG file

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generated image to output.png \n")
}

func testPoint(b borders, x int, y int, maxX int, maxY int, variants int) int {

	point := takeThePoint(b, x, y, maxX, maxY)
	variant := evaluate(point, variants)

	return variant
}

func takeThePoint(b borders, x int, y int, maxX int, maxY int) point {

	xValue := b.X1 + (b.X2-b.X1)*float64(x)/float64(maxX)
	yValue := b.Y1 + (b.Y2-b.Y1)*float64(y)/float64(maxY)
	// log.Printf("x, y %v:%v", xValue, yValue)
	return point{X: xValue, Y: yValue}
}

func evaluate(p point, variants int) int {

	var a, b, t float64
	var loops int
	for loops = 0; loops < variants; loops++ {
		t = a*a - b*b
		b = 2 * a * b
		a = t
		a = a + p.X
		b = b + p.Y

		var m = a*a + b*b
		if m > 10 {
			return loops
		}
	}

	return variants
}

func rainbow(n uint8) color.RGBA {
	return color.RGBA{wave(n), wave(n + 85), wave(n + 170), 255}
}

func wave(n uint8) uint8 {
	return uint8(125*math.Sin(float64(n)/15)) + 126
}
