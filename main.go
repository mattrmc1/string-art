package main

import (
	"image"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 1600
	HEIGHT = 800
	TWO_PI = math.Pi * 2

	PATH = "images/mosdef.jpeg"
	N    = 8

	SEG = TWO_PI / float64(N)
)

var r float32

var bounds image.Rectangle
var colors []raylib.Color
var grayscale []float64

var pins []raylib.Vector2
var lines [][2]raylib.Vector2

func calculateLinePointPosition(n int) raylib.Vector2 {
	theta := SEG * float64(n)
	return centerVector(float32(float64(r)*math.Cos(theta)), float32(float64(r)*math.Sin(theta)))
}

func processImage() {
	imgRaw := raylib.LoadImage(PATH)
	raylib.ImageColorGrayscale(imgRaw)

	bounds = imgRaw.ToImage().Bounds()
	r = calculateRadius(bounds.Max.X, bounds.Max.Y, INSCRIBE_MIN)

	colors = raylib.LoadImageColors(imgRaw)
	grayscale = make([]float64, len(colors))
	for i, c := range colors {

		if c.A > 50 {
			grayscale[i] = float64(c.R)
		} else {
			grayscale[i] = 255
		}
	}

	raylib.UnloadImage(imgRaw)
}

func processPins() {
	for i := 0; i < N; i++ {
		pos := calculateLinePointPosition(i)
		pins = append(pins, pos)
	}
}

func processAllPotentialLines() {
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			lines = append(lines, [2]raylib.Vector2{pins[i], pins[j]})
		}
	}
}

func processLines() {
	// 	Greedy approach:
	//		-> start at n=0
	//		-> for all lines connected to pins[n]
	//			-> if already visited, continue
	//			-> if too close to pin, continue
	//			-> find line of best fit
	//		-> set best line as visited
	//		-> set pos = second point of best line
	//		-> loop
	//		-
	//		-> break if no line improves the cost? (definition of "cost"?)
	//		-> break if max iter (i.e. max lines) is reached?
	//		-> break if all connected lines have already been visited?

	// * find line of best fit?
	//		-> The high grayscale sum along the line?
	//		-> buffer around pixels surrounding line?
}

func drawLines() {
	// Option 1: Render all calculated lines
	// Option 2: Animate a line for each dt
}

func printStringPath() {
	// Print array of n positions in order for string art
	// For example:
	//	given path is [0, 4, 3]
	//	we should start at 0 and connect a string from 0 -> 4 -> 3
}

func main() {
	raylib.InitWindow(WIDTH, HEIGHT, "String Art")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	start()

	for !raylib.WindowShouldClose() {
		draw()
	}
}

func start() {
	// Upload image and store in grayscale pixel array
	processImage()

	// Calculate the pins vector2 array
	processPins()

	// Calculate and store pixel array of all potential lines
	processAllPotentialLines()

	// Greedy algo to find line of best fit until threshold
	processLines()
}

func draw() {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.RayWhite)

	// debug_draw_pixelArr_image()
	// debug_draw_circle()
	debug_draw_pins()
	debug_draw_potential_lines()
}
