package hero

import (
	"fmt"

	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	s "github.com/WatchJani/Role-playing-game/model/skeleton"
)

type Hero struct {
	Killed int
	s.Skeleton
}

func EmptyHero() *Hero {
	return &Hero{}
}

func (h *Hero) UpdateHero(Height, Width, SpawnBorderHero float64) *Hero {
	return &Hero{
		Skeleton: s.NewSkeletonHero(
			Width,
			Height,
			h.Health,
			SpawnBorderHero,
			h.Weapon,
		),
	}
}

func TestHeroNew(x, y, health float64) Hero {
	return Hero{
		Skeleton: s.Skeleton{
			X:      x,
			Y:      y,
			Health: health,
		},
	}
}

func (h Hero) String() {
	fmt.Printf("Exp: %f | Health: %f | Weapon: %s | Killed: %d | X-coordination: %f | Y-coordination: %f\n", h.Exp, h.Health, h.Weapon.Name, h.Killed, h.X, h.Y)
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
