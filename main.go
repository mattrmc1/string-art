package main

import (
	"fmt"
	"image"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 1600
	HEIGHT = 800
	TWO_PI = math.Pi * 2

	PATH   = "images/mosdef.jpeg"
	N      = 8
	BUFFER = 1

	SEG   = TWO_PI / float64(N)
	MAX_L = (N * (N - 1)) / 2
)

type Line struct {
	start  raylib.Vector2
	end    raylib.Vector2
	pixels []raylib.Vector2
}

var r float32

var bounds image.Rectangle
var colors []raylib.Color
var grayscale []float64

var pins []raylib.Vector2
var lines = map[string]Line{}

var path []int

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
		for j := i + BUFFER; j < N; j++ {
			k := toStrKey(i, j)
			lines[k] = createLine(pins[i], pins[j])
		}
	}
}

func calculateCost(err []float64, l Line) float64 {
	// TODO
	return 1.0
}

func processLines() {

	var cost []float64 // white bounds - grayscale

	visited := map[string]bool{}
	startPin := 0

	for i := 0; i < MAX_L; i++ {
		endPin := -1
		maxCost := 0.0

		for n := BUFFER; n < N-BUFFER; n++ {
			p := (startPin + n) % N
			k := toStrKey(startPin, p)
			if visited[k] {
				continue
			}

			curCost := calculateCost(cost, lines[k])
			if curCost > maxCost {
				maxCost = curCost
				endPin = p
			}
		}

		if endPin == -1 {
			break
		}

		path = append(path, endPin)
		visited[toStrKey(startPin, endPin)] = true

		startPin = endPin
	}
}

func drawPath() {
	for i := range path {
		if i == 0 {
			continue
		}

		k := toStrKey(path[i], path[i-1])
		raylib.DrawLineV(lines[k].start, lines[k].end, raylib.LightGray)
	}
}

func printPath() {
	fmt.Println(path)
}

func main() {
	raylib.InitWindow(WIDTH, HEIGHT, "String Art")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	process()

	for !raylib.WindowShouldClose() {
		draw()
	}
}

func process() {
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

	debug_draw_image()
	debug_draw_circle()
	debug_draw_pins()
	debug_draw_potential_lines()

	// drawPath()
}
