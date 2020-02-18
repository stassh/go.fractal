package main

import (
	"encoding/base64"
	"fmt"
	"fractal/entity"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"net/url"

	proto "github.com/golang/protobuf/proto"
)

func GenerateFractalImage(imageMaxX int, imageMaxY int, b Borders) image.Image {
	imgRect := image.Rect(0, 0, imageMaxX, imageMaxY)
	img := image.NewRGBA(imgRect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for y := 0; y < imageMaxX; y += 1 {
		for x := 0; x < imageMaxY; x += 1 {
			variant := testPoint(b, x, y, imageMaxX, imageMaxY, 255)
			img.Set(x, y, &image.Uniform{rainbow(uint8(variant))})
		}
	}
	return img
}

func testPoint(b Borders, x int, y int, maxX int, maxY int, variants int) int {

	point := takeThePoint(b, x, y, maxX, maxY)
	variant := evaluate(point, variants)

	return variant
}

func takeThePoint(b Borders, x int, y int, maxX int, maxY int) point {

	xValue := b.X1 + (b.X2-b.X1)*float64(x)/float64(maxX)
	yValue := b.Y1 + (b.Y2-b.Y1)*float64(y)/float64(maxY)
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

func getFragmet(b Borders, i int, j int, maxI int, maxJ int) Borders {
	var borders Borders

	deltaX := (b.X2 - b.X1) / float64(maxI)
	deltaY := (b.Y2 - b.Y1) / float64(maxJ)

	borders.X1 = b.X1 + deltaX*float64(i)
	borders.X2 = borders.X1 + deltaX

	borders.Y1 = b.Y1 + deltaY*float64(j)
	borders.Y2 = borders.Y1 + deltaY

	return borders

}

func parseProto(value string) Borders {

	fractalInfo := entity.FractalInfo{}
	bytes, _ := base64.StdEncoding.DecodeString(value)

	if err := proto.Unmarshal(bytes, &fractalInfo); err != nil {
		log.Fatalln("Failed to parse fractal info:", err)
	}

	return Borders{X1: fractalInfo.Rectangle.Left, X2: fractalInfo.Rectangle.Right, Y1: fractalInfo.Rectangle.Bottom, Y2: fractalInfo.Rectangle.Top}

}

func getParams(b Borders) string {
	params := url.Values{}
	params.Add("x1", fmt.Sprintf("%f", b.X1))
	params.Add("x2", fmt.Sprintf("%f", b.X2))
	params.Add("y1", fmt.Sprintf("%f", b.Y1))
	params.Add("y2", fmt.Sprintf("%f", b.Y2))

	info := entity.FractalInfo{}
	info.Rectangle = &entity.Rectangle{Top: b.Y2, Bottom: b.Y1, Left: b.X1, Right: b.X2}

	log.Printf("Fractal info: %v", info)

	infoProto, err := proto.Marshal(&info)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	log.Printf("Fractal info: %v", infoProto)

	protoInfo := base64.StdEncoding.EncodeToString(infoProto)
	params.Add("p", protoInfo)

	return params.Encode()
}

type Borders struct {
	X1 float64
	X2 float64
	Y1 float64
	Y2 float64
}

type point struct {
	X float64
	Y float64
}
