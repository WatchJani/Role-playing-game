package game_math

import (
	"math"
	"math/rand"
)

func GeneratePosition(radius, max, coordination float64) float64 {
	for {
		position := RandomNumber(0, max)

		if position < coordination-radius || position > coordination+radius {
			return position
		}
	}
}

func RandomNumber(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GetNumber(max int) int {
	return rand.Intn(max)
}

func BorderPercentage(value float64, borderHero float64) float64 {
	return Percentage(value, borderHero)
}

func Percentage(value, percentage float64) float64 {
	return (value / 100) * percentage
}

func Distance(XA, YA, XB, YB float64) float64 {
	return math.Sqrt(math.Pow(XA-XB, 2) + math.Pow(YA-YB, 2))
}

func NewVector(XA, YA, XB, YB float64) (float64, float64) {
	return XB - XA, YB - YA
}

func NormalizationVector(x, y, distance float64) (float64, float64) {
	return x / distance, y / distance
}
