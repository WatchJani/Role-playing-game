package hero

import (
	"fmt"

	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	s "github.com/WatchJani/Role-playing-game/model/skeleton"
	w "github.com/WatchJani/Role-playing-game/model/weapon"
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
			X:      x,
			Y:      y,
			Health: health,
			// Weapon: make(map[string]Weapon),
		},
	}
}

func NewSkeletonEnemy(height, width, radius, exp, health float64, weapon w.Weapon, hero Hero) s.Skeleton {
	return s.Skeleton{
		X:      m.GeneratePosition(radius, width, hero.X),
		Y:      m.GeneratePosition(radius, height, hero.Y),
		Exp:    exp,
		Health: health,
		Weapon: weapon, //default weapon for enemy
	}
}

func (h Hero) String() {
	fmt.Printf("Exp: %f | Health: %f | Weapon: \n", h.Exp, h.Health)
}

func (h Hero) Damage() float64 {
	return m.RandomNumber(h.Weapon.MinDamageTake, h.Weapon.MaxDamageTake)
}

func (h *Hero) SetEXP(exp float64) {
	h.Exp += exp
}

func (h *Hero) TakeDamage(damage float64) {
	h.Health -= damage
}

func (h *Hero) IsAlive() bool {
	return h.Health >= 0
}
