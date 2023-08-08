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

func RandomNoRepeatable(object int) (int, int, int) {
	var count int = 3

	if object < count {
		panic("Game configuration is not valid\n Check all items\n Min number of items i 3")
	}

	uniqueNumbers := make([]int, 0, count)
	numberSet := make(map[int]bool)

	for len(uniqueNumbers) < count {
		num := GetNumber(3)
		if !numberSet[num] {
			numberSet[num] = true
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}

	return uniqueNumbers[0], uniqueNumbers[1], uniqueNumbers[2]
}
