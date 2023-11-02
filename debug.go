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

func debug_draw_potential_lines_old() {
	for _, lineV := range lines_old {
		p1 := lineV[0]
		p2 := lineV[1]
		raylib.DrawLineV(p1, p2, raylib.LightGray)
	}
}

func debug_draw_potential_lines_seperate() {
	for k, xs := range lines_x {
		points := make([]raylib.Vector2, len(xs))
		ys := lines_y[k]
		for i, x := range xs {
			y := ys[i]
			points[i] = raylib.Vector2{float32(x), float32(y)}
		}
		raylib.DrawLineStrip(points, int32(len(xs)), raylib.Green)
	}
}

func debug_draw_potential_lines_alt() {
	for _, l := range linesAlt {
		points := make([]raylib.Vector2, len(l))
		xs := l[0]
		ys := l[1]
		for i, x := range xs {
			y := ys[i]
			points[i] = raylib.Vector2{float32(x), float32(y)}
		}
		raylib.DrawLineStrip(points, int32(len(xs)), raylib.Blue)
	}
}

func debug_draw_potential_lines() {
	for _, l := range lines {
		raylib.DrawLineStrip(l.pixels, int32(len(l.pixels)), raylib.DarkGreen)
	}
}

func debug_draw_image() {
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
