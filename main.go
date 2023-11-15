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

	PATH        = "images/mosdef.jpeg"
	N           = 288
	BUFFER      = 10
	LINE_WEIGHT = 20

	SEG = TWO_PI / float64(N)
	// MAX_L = (N * (N - 1)) / 2
	MAX_L = 3000
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
		grayscale[i] = 255 - float64(c.R)
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

func calculateCost(grayscaleCopy []float64, l Line) float64 {
	sum := 0.0
	for _, px := range l.pixels {
		x := int(px.X) - WIDTH/2 + bounds.Max.X/2
		y := int(px.Y) - HEIGHT/2 + bounds.Max.Y/2
		i := y*bounds.Max.X + x

		if i < 0 || i > len(grayscaleCopy)-1 {
			continue
		}
		sum += math.Max(grayscaleCopy[i], 0)
	}
	return sum
}

func processLines() {

	var grayscaleCopy = make([]float64, len(grayscale))
	copy(grayscaleCopy, grayscale)

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

			curCost := calculateCost(grayscaleCopy, lines[k])
			if curCost > maxCost {
				maxCost = curCost
				endPin = p
			}
		}

		// if endPin == -1 {
		// 	fmt.Printf("\n\n\nBROKE LOOP at %v, cost %v\n\n\n", i, maxCost)
		// 	break
		// }

		path = append(path, endPin)
		k := toStrKey(startPin, endPin)
		visited[k] = true

		for _, px := range lines[k].pixels {
			x := int(px.X) - WIDTH/2 + bounds.Max.X/2
			y := int(px.Y) - HEIGHT/2 + bounds.Max.Y/2
			i := y*bounds.Max.X + x
			if i < 0 || i > len(grayscaleCopy)-1 {
				continue
			}
			grayscaleCopy[i] = math.Max(grayscaleCopy[i]-LINE_WEIGHT, 0)
		}

		startPin = endPin
		// fmt.Printf("Processing... (%v/%v)\n", len(path), MAX_L)
	}
}

func drawPath() {
	for i := range path {
		if i == 0 {
			continue
		}

		k := toStrKey(path[i], path[i-1])
		raylib.DrawLineV(lines[k].start, lines[k].end, raylib.Black)
	}
}

var pIdx = 1

func animatePath(dt float32, repeat bool) {

	if repeat && pIdx > len(path)-1 {
		pIdx = 1
	}

	idx := int(math.Min(float64(pIdx), float64(len(path)-1)))
	for i := 1; i <= idx; i++ {
		k := toStrKey(path[i], path[i-1])
		raylib.DrawLineV(lines[k].start, lines[k].end, raylib.DarkGray)
	}

	raylib.DrawText(fmt.Sprintf("Line count: %v", pIdx), 12.0, 12.0, 32, raylib.Black)
	raylib.DrawText(fmt.Sprintf("dt: %v", dt), 12.0, HEIGHT-12-32-12-32, 32, raylib.Black)
	raylib.DrawText(fmt.Sprintf("FPS: %v", raylib.GetFPS()), 12.0, HEIGHT-12-32, 32, raylib.Black)

	if pIdx < len(path) || repeat {
		pIdx += 1
	}
}

func printPath() {
	fmt.Println()
	fmt.Println(path)
	fmt.Println()
}

func main() {
	raylib.InitWindow(WIDTH, HEIGHT, "String Art")
	defer raylib.CloseWindow()

	// raylib.SetTargetFPS(60)

	process()

	for !raylib.WindowShouldClose() {
		dt := raylib.GetFrameTime()
		draw(dt)
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

	// printPath()
}

func draw(dt float32) {
	raylib.BeginDrawing()
	defer raylib.EndDrawing()

	raylib.ClearBackground(raylib.RayWhite)

	// debug_draw_image()
	// debug_draw_circle()
	// debug_draw_pins()
	// debug_draw_potential_lines()
	// debug_draw_potential_line_px()
	// debug_draw_potential_lines_img()

	// drawPath()
	animatePath(dt, true)
}
