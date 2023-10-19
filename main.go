package main

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 1600
	HEIGHT = 800
	TWO_PI = math.Pi * 2

	PATH = "images/bird.png"
	N    = 8
)

type Line struct {
	startPos raylib.Vector2
	endPos   raylib.Vector2
	color    raylib.Color
}

type Pixel struct {
	pos   raylib.Vector2
	color raylib.Color
}

var r float32
var lines []Line

// Can do a line strip instead since we're doing the greedy algo
// DrawLineStrip(Vector2 *points, int pointCount, Color color);
// var lineSeq []raylib.Vector2

var pixels []Pixel

func processImage() {
	img := raylib.LoadImage(PATH)
	bounds := img.ToImage().Bounds()
	r = calculateRadius(bounds.Max.X, bounds.Max.Y)

	for i, c := range raylib.LoadImageColors(img) {
		x := (i % bounds.Max.X) - bounds.Max.X/2
		y := (i / bounds.Max.Y) - bounds.Max.Y/2

		pos := centerVector(float32(x), float32(y))
		toGray := toGrayScale(c)

		r, g, b, a := toGray.RGBA()
		whiteness := (r + g + b) / (255 + 255 + 255)

		if a > uint32(10) && whiteness < uint32(120) {
			pixels = append(pixels, Pixel{pos, toGrayScale(c)})
		}

	}
}

func addLine(startPos, endPos int) {
	seg := TWO_PI / float64(N)

	startTheta := seg * float64(startPos)
	startV := centerVector(float32(float64(r)*math.Cos(startTheta)), float32(float64(r)*math.Sin(startTheta)))

	endTheta := seg * float64(endPos)
	endV := centerVector(float32(float64(r)*math.Cos(endTheta)), float32(float64(r)*math.Sin(endTheta)))

	lines = append(lines, Line{startV, endV, raylib.DarkGray})
}

func calculateError() int {
	return 0
}

func findBestLine(startPos int, drawn map[int]bool) {

	// minError := calcCurrentError()
	// for all int m between [0,N] where startPoint != m & !drawn[nm]
	//		calc err with this line drawn
	//		if this line is better, store the endPos and minError
	//			minError = min(minError, err)
	//			endPos = m

	// if the minError was NOT improved, return

	// else:
	//	update inclusionScalar for this line to 1 (i.e. include this line)
	//	update drawn for this line
	//	look for new best line -> findBestLine(endPos, drawn)

	var endPos int
	found := false
	minError := calculateError()

	for p := 0; p < N; p++ {
		if startPos == p || drawn[toIntKey(startPos, p)] {
			continue
		}

		err := calculateError()
		if err < minError {
			minError = err
			endPos = p
			found = true
		}
	}

	if found {
		addLine(startPos, endPos)
		drawn[toIntKey(startPos, endPos)] = true
		findBestLine(endPos, drawn)
	}

}

func processLines() {
	findBestLine(0, make(map[int]bool))
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
		draw()
	}
}

func start() {
	// processImage()
	// processLines()
}

func draw() {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.RayWhite)

	// drawNodes()
	// drawLines()
	// drawImage()

}
