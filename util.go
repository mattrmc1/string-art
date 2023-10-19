package main

import (
	"image/color"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func centerVector(x, y float32) raylib.Vector2 {
	return raylib.Vector2{X: WIDTH/2 + x, Y: HEIGHT/2 + y}
}

func centerVectorV(v raylib.Vector2) raylib.Vector2 {
	return raylib.Vector2{X: WIDTH/2 + v.X, Y: HEIGHT/2 + v.Y}
}

func grayscale(c raylib.Color) raylib.Color {
	r, g, b, a := c.RGBA()
	y := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	r, g, b, _ = color.Gray{uint8(y / 256)}.RGBA()

	return raylib.NewColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

// TODO
func pixelImportance(c raylib.Color) float64 {
	return float64(0)
}
