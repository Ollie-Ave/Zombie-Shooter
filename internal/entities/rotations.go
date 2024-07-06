package entities

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func calcuateAngleForDirectionVector(vector rl.Vector2) float64 {
	angleRad := math.Atan2(float64(vector.Y), float64(vector.X))

	angleDeg := angleRad * (180 / math.Pi)

	if angleDeg < 0 {
		angleDeg += 360
	}

	return angleDeg + 90
}
