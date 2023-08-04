package skeleton

import (
	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	w "github.com/WatchJani/Role-playing-game/model/weapon"
)

type Skeleton struct {
	X      float64
	Y      float64
	Exp    float64
	Health float64
	Weapon w.Weapon
}

func (e Skeleton) GetX() float64 {
	return e.X
}

func (e Skeleton) GetY() float64 {
	return e.Y
}

func NewSkeletonHero(width, height, health, SpawnBorderHero float64) Skeleton {
	widthBorder, heightBorder := m.BorderPercentage(width, SpawnBorderHero), m.BorderPercentage(height, SpawnBorderHero)

	return Skeleton{
		X:      m.RandomNumber(widthBorder, width-widthBorder),
		Y:      m.RandomNumber(heightBorder, height-heightBorder),
		Health: health,
		// Weapon: NewWeapon("Scattergun", 500, 150, 300), //default weapon for hero
	}
}
