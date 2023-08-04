package model

import (
	"fmt"

	s "github.com/WatchJani/Role-playing-game/model/skeleton"
)

type Hero struct {
	s.Skeleton
}

func NewHero(Height, Width, HeroHealth, SpawnBorderHero float64) Hero {
	return Hero{
		s.NewSkeletonHero(
			Width,
			Height,
			HeroHealth,
			SpawnBorderHero,
		),
	}
}

func TestHeroNew(x, y, health float64) Hero {
	return Hero{
		Skeleton: s.Skeleton{
			X: x,
			Y: y,
			// Health: health,
			// Weapon: make(map[string]Weapon),
		},
	}
}

func (h Hero) String() {
	fmt.Printf("Exp: %f | Health: %f | Weapon: \n", h.Exp, h.Health)
}

// func (h Hero) Damage() float64 {
// 	return m.RandomNumber(h.Weapon.MinDamageTake, h.Weapon.MaxDamageTake)
// }

func (h *Hero) SetEXP(exp float64) {
	h.Exp += exp
}

func (h *Hero) TakeDamage(damage float64) {
	h.Health -= damage
}

func (h *Hero) IsAlive() bool {
	return h.Health >= 0
}
