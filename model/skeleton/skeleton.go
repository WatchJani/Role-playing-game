package skeleton

import (
	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	w "github.com/WatchJani/Role-playing-game/model/weapon"
)

type Skeleton struct {
	X      float64
	Y      float64
	Exp    float64  `json:"Exp"`
	Health float64  `json:"Health"`
	Weapon w.Weapon `json:"Weapon"`
}

func (e Skeleton) GetX() float64 {
	return e.X
}

func (e Skeleton) GetY() float64 {
	return e.Y
}

func NewSkeletonHero(width, height, health, SpawnBorderHero float64, weapon w.Weapon) Skeleton {
	widthBorder, heightBorder := m.BorderPercentage(width, SpawnBorderHero), m.BorderPercentage(height, SpawnBorderHero)

	return Skeleton{
		X:      m.RandomNumber(widthBorder, width-widthBorder),
		Y:      m.RandomNumber(heightBorder, height-heightBorder),
		Health: health,
		Weapon: weapon,
	}
}
