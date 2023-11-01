package main

import (
	"image/color"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func debug_draw_circle() {
	c := centerVector(0, 0)
	raylib.DrawRing(c, r, r+1, 0, 360, 1, raylib.Red)
}

func debug_draw_pins() {
	for _, pin := range pins {
		raylib.DrawCircleV(pin, 1, raylib.Lime)
	}
}

func debug_draw_potential_lines() {
	for _, lineV := range lines {
		p1 := lineV[0]
		p2 := lineV[1]
		raylib.DrawLineV(p1, p2, raylib.LightGray)
	}
}

func debug_draw_pixelArr_image() {
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			i := y*bounds.Max.X + x

			// cx -> x position if center is (0,0)
			cx := float32(x - bounds.Max.X/2)

			// cy -> y position if center is (0,0)
			cy := float32(y - bounds.Max.Y/2)

			// If not in bounds of the circle
			isInBounds := math.Pow(float64(cx), 2)+math.Pow(float64(cy), 2) <= math.Pow(float64(r), 2)
			if !isInBounds {
				continue
			}

			v := uint8(grayscale[i])
			raylib.DrawPixelV(centerVector(cx, cy), color.RGBA{v, v, v, 255})
		}
	}
}
