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
	for _, l := range lines {
		raylib.DrawLineStrip(l.pixels, int32(len(l.pixels)), raylib.DarkGreen)
	}
}

func debug_draw_potential_line_px() {
	for _, l := range lines {
		for _, px := range l.pixels {
			raylib.DrawPixelV(px, raylib.Brown)
		}

	}
}

func debug_draw_potential_lines_img() {
	for _, l := range lines {
		for _, px := range l.pixels {
			x := int(px.X) - WIDTH/2 + bounds.Max.X/2
			y := int(px.Y) - HEIGHT/2 + bounds.Max.Y/2
			i := y*bounds.Max.X + x

			var c color.RGBA
			if i < 0 || i > len(grayscale)-1 {
				c = raylib.Red
			} else {
				g := uint8(grayscale[i])
				c = color.RGBA{g, g, g, 255}
			}

			raylib.DrawPixelV(px, c)
		}
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
