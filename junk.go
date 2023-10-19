package main

import (
	"fmt"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var t float32

func junk_animatedLines(dt float32) {
	seg := TWO_PI / float64(N)

	for i := 0; i < N; i++ {

		// Skip some lines
		// random := rand.Float32()
		// if random > 0.5 {
		// 	continue
		// }

		theta1 := seg * float64(i)
		y1 := float32(float64(r) * math.Sin(theta1))
		x1 := float32(float64(r) * math.Cos(theta1))

		theta2 := seg * (float64(int(t)%10+4)*float64(i) + (float64(t)*float64(int(t)%10+4)*0.01*TWO_PI)*float64(N))
		y2 := float32(float64(r) * math.Sin(theta2))
		x2 := float32(float64(r) * math.Cos(theta2))

		raylib.DrawLineV(centerVector(x1, y1), centerVector(x2, y2), raylib.Gray)
	}
}

func calculateLinesAlt() {
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
