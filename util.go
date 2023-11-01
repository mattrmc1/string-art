package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type CircleRender int

const (
	CIRCUMSCRIBE CircleRender = iota // Circumscribe around rect
	INSCRIBE_MIN                     // Inscribe with diameter set to smallest side
	INSCRIBE_MAX                     // Inscribe with diameter set to largest side
)

func centerVector(x, y float32) raylib.Vector2 {
	return raylib.Vector2{X: WIDTH/2 + x, Y: HEIGHT/2 + y}
}

func calculateRadius(a, b int, render CircleRender) float32 {
	switch render {
	case INSCRIBE_MIN:
		return float32(math.Min(float64(a), float64(b)) / 2)
	case INSCRIBE_MAX:
		return float32(math.Max(float64(a), float64(b)) / 2)
	case CIRCUMSCRIBE:
		return float32(math.Sqrt((math.Pow(float64(a), 2) + math.Pow(float64(b), 2)) / 4))
	default:
		return float32(math.Sqrt((math.Pow(float64(a), 2) + math.Pow(float64(b), 2)) / 4))
	}
}

func toStrKey(n, m int) string {
	arr := []int{n, m}
	sort.Ints(arr)
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", "", -1), "[]")
}
