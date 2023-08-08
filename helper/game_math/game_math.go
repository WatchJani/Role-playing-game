package game_math

import (
	"math"
	"math/rand"
)

// random coordination in
func GeneratePosition(radius, max, coordination float64) float64 {
	for {
		position := RandomNumber(0, max)

		if position < coordination-radius || position > coordination+radius {
			return position
		}
	}
}

// random number between two number
func RandomNumber(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// random number in range of number
func GetNumber(max int) int {
	return rand.Intn(max)
}



// Percentage of value (20% of 150)
func Percentage(value, percentage float64) float64 {
	return (value / 100) * percentage
}

// distance between two points
func Distance(XA, YA, XB, YB float64) float64 {
	return math.Sqrt(math.Pow(XA-XB, 2) + math.Pow(YA-YB, 2))
}

// create new vector
func NewVector(XA, YA, XB, YB float64) (float64, float64) {
	return XB - XA, YB - YA
}

// vector normalization
func NormalizationVector(x, y, distance float64) (float64, float64) {
	return x / distance, y / distance
}
