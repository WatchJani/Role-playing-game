package game_math

import (
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

func MinValue(value, min float64) float64 {
	if value < min {
		return min
	}
	return value
}

// func Distance(enemy, hero model.Position) float64 {
// 	return math.Sqrt(math.Pow(enemy.GetX()-hero.GetX(), 2) + math.Pow(enemy.GetY()-hero.GetY(), 2))
// }

// func NewVector(enemy, hero model.Position) (float64, float64) {
// 	return hero.GetX() - enemy.GetX(), hero.GetY() - enemy.GetY()
// }

func NormalizationVector(x, y, distance float64) (float64, float64) {
	return x / distance, y / distance
}
