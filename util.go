package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func centerVector(x, y float32) raylib.Vector2 {
	return raylib.Vector2{X: WIDTH/2 + x, Y: HEIGHT/2 + y}
}

func toGrayScale(c raylib.Color) raylib.Color {
	r, g, b, a := c.RGBA()
	y := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	r, g, b, _ = color.Gray{uint8(y / 256)}.RGBA()

	return raylib.NewColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

func calculateRadius(a, b int) float32 {
	return float32(math.Sqrt((math.Pow(float64(a), 2) + math.Pow(float64(b), 2)) / 4))
}

func toStrKey(n, m int) string {
	return strings.Trim(strings.Replace(fmt.Sprint([]int{n, m}), " ", "", -1), "[]")
}

func toIntKey(n, m int) int {
	return n*N + m
}
