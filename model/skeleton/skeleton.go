package model

import m "github.com/WatchJani/Role-playing-game/helper/game_math"

type Skeleton struct {
	X      float64
	Y      float64
	Exp    float64
	Health float64
	// Weapon Weapon
}

func (e Skeleton) GetX() float64 {
	return e.X
}

func (e Skeleton) GetY() float64 {
	return e.Y
}

// func NewSkeletonEnemy(height, width, radius, exp, health float64, weapon Weapon, hero Hero) Skeleton {
// 	return Skeleton{
// 		X:      m.GeneratePosition(radius, width, hero.X),
// 		Y:      m.GeneratePosition(radius, height, hero.Y),
// 		Exp:    exp,
// 		Health: health,
// Weapon: weapon, //default weapon for enemy
// 	}
// }

func NewSkeletonHero(width, height, health, SpawnBorderHero float64) Skeleton {
	widthBorder, heightBorder := m.BorderPercentage(width, SpawnBorderHero), m.BorderPercentage(height, SpawnBorderHero)

	return Skeleton{
		X:      m.RandomNumber(widthBorder, width-widthBorder),
		Y:      m.RandomNumber(heightBorder, height-heightBorder),
		Health: health,
		// Weapon: NewWeapon("Scattergun", 500, 150, 300), //default weapon for hero
	}
}
