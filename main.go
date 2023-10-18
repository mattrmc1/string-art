package main

import (
	"image/color"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 1200
	HEIGHT = 800
	TWO_PI = math.Pi * 2

	PATH = "images/test-image-2.jpeg"
	N    = 80
)

var radius float32

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
var testPixels []Pixel

func centerVector(x, y float32) raylib.Vector2 {
	return raylib.Vector2{X: WIDTH/2 + x, Y: HEIGHT/2 + y}
}

func calculateLinesTest() {
	seg := TWO_PI / float64(N)

	for i := 0; i < N; i++ {

		// Skip some lines
		// random := rand.Float32()
		// if random > 0.5 {
		// 	continue
		// }

		theta1 := seg * float64(i)
		y1 := float32(float64(radius) * math.Sin(theta1))
		x1 := float32(float64(radius) * math.Cos(theta1))

		theta2 := seg * (0.1*float64(i) + float64(N))
		y2 := float32(float64(radius) * math.Sin(theta2))
		x2 := float32(float64(radius) * math.Cos(theta2))

		lines = append(lines, Line{centerVector(x1, y1), centerVector(x2, y2), raylib.Black})
	}
}

func grayscale(c raylib.Color) raylib.Color {
	r, g, b, a := c.RGBA()
	y := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	r, g, b, _ = color.Gray{uint8(y / 256)}.RGBA()

	return raylib.NewColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

func calculateRadius(a, b int) {
	radius = float32(math.Sqrt((math.Pow(float64(a), 2) + math.Pow(float64(b), 2)) / 4))
}

func uploadImage() {
	img := raylib.LoadImage(PATH)
	if img == nil {
		return
	}

	bounds := img.ToImage().Bounds()
	calculateRadius(bounds.Max.X, bounds.Max.Y)
	colors := raylib.LoadImageColors(img)

	for i, c := range colors {
		x := (i % bounds.Max.X) - bounds.Max.X/2
		y := (i / bounds.Max.Y) - bounds.Max.Y/2

		pos := centerVector(float32(x), float32(y))
		testPixels = append(testPixels, Pixel{pos, grayscale(c)})
	}
}

func calculateLines() {
	// Real way to solve this:
	// 1. Upload image
	// 2. Quantify grayscale of pixels
	// 3. Need some sort of "cost function"
	// 4. Find the best line to reduce cost function
	// 5. From that line, find the next line that will best reduce cost func
	// 6. Repeat until the cost is below some threshold

	// Upload image
	uploadImage()
}

func drawNodes() {
	seg := TWO_PI / float64(N)

	for i := 0; i < N; i++ {
		theta := seg * float64(i)
		y := float32(float64(radius) * math.Sin(theta))
		x := float32(float64(radius) * math.Cos(theta))
		raylib.DrawCircleV(centerVector(x, y), 1, raylib.Black)
	}
}

func drawLines() {
	for _, l := range lines {
		raylib.DrawLineV(l.startPos, l.endPos, raylib.Black)
	}
}

func drawImage() {
	for _, p := range testPixels {
		raylib.DrawPixelV(p.pos, p.color)
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
	uploadImage()

	calculateLinesTest()
	// calculateLines()
}

func draw() {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.RayWhite)

	c := centerVector(0, 0)
	raylib.DrawCircleV(c, radius, raylib.LightGray)

	drawNodes()
	drawLines()

	drawImage()
}
