package main

import (
	"fmt"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 1600
	HEIGHT = 800
	TWO_PI = math.Pi * 2

	PATH = "images/bird.png"
	N    = 200
)

var r float32 = float32(380)

type Line struct {
	startPos raylib.Vector2
	endPos   raylib.Vector2
	color    raylib.Color
}

type Pixel struct {
	pos   raylib.Vector2
	color raylib.Color
}

var lines []Line
var pixels []Pixel

func calculateRadius(a, b int) {
	r = float32(math.Sqrt((math.Pow(float64(a), 2) + math.Pow(float64(b), 2)) / 4))
}

func uploadImage() {
	img := raylib.LoadImage(PATH)
	bounds := img.ToImage().Bounds()
	calculateRadius(bounds.Max.X, bounds.Max.Y)

	for i, c := range raylib.LoadImageColors(img) {
		x := (i % bounds.Max.X) - bounds.Max.X/2
		y := (i / bounds.Max.Y) - bounds.Max.Y/2

		pos := centerVector(float32(x), float32(y))
		toGray := grayscale(c)

		r, g, b, a := toGray.RGBA()
		whiteness := (r + g + b) / (255 + 255 + 255)

		if a > uint32(10) && whiteness < uint32(120) {
			pixels = append(pixels, Pixel{pos, grayscale(c)})
		}

	}
}

func calculateLines() {
	// Real way to solve this:
	// - greedy algo:
	// - start at arbitrary node
	// - find line that will minimize error the most
	// - repeat process starting at prev line endPos
	// - escape when no line will minimize the error

	fmt.Printf("\n\n START \n\n")
	fmt.Printf("pixels %v", len(pixels))

	seg := TWO_PI / float64(N)
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			theta_i := seg * float64(i)
			y_i := float32(float64(r) * math.Sin(theta_i))
			x_i := float32(float64(r) * math.Cos(theta_i))

			theta_j := seg * float64(j)
			y_j := float32(float64(r) * math.Sin(theta_j))
			x_j := float32(float64(r) * math.Cos(theta_j))

			lineVector_i := centerVector(x_i, y_i)
			lineVector_j := centerVector(x_j, y_j)

			isIntersecting := false
			for _, p := range pixels {
				if raylib.CheckCollisionPointLine(p.pos, lineVector_i, lineVector_j, 1) {
					isIntersecting = true
					break
				}
			}

			if !isIntersecting {
				lines = append(lines, Line{lineVector_i, lineVector_j, raylib.DarkGray})
			}
		}
	}

	fmt.Printf("\n\n END \n\n")
	fmt.Printf("lines %v", len(lines))

}

func drawNodes() {
	seg := TWO_PI / float64(N)
	for i := 0; i < N; i++ {
		theta := seg * float64(i)
		y := float32(float64(r) * math.Sin(theta))
		x := float32(float64(r) * math.Cos(theta))
		raylib.DrawCircleV(centerVector(x, y), 1, raylib.Black)
	}
}

func drawLines() {
	for _, l := range lines {
		raylib.DrawLineV(l.startPos, l.endPos, l.color)
		// raylib.DrawLineEx(l.startPos, l.endPos, 5, l.color)
	}
}

func drawImage() {
	for _, p := range pixels {
		raylib.DrawPixelV(p.pos, raylib.ColorAlpha(p.color, 1))
	}
}

func main() {
	raylib.InitWindow(WIDTH, HEIGHT, "String Art")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	start()

	for !raylib.WindowShouldClose() {
		dt := raylib.GetFrameTime()
		t += dt
		draw(dt)
	}
}

func start() {
	uploadImage()
	calculateLinesAlt()
}

func draw(dt float32) {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.RayWhite)

	drawLines()
	drawImage()
}
