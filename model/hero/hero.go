package hero

import (
	"fmt"

	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	s "github.com/WatchJani/Role-playing-game/model/skeleton"
)

type Hero struct {
	Killed    int
	MaxHealth float64
	s.Skeleton
}

func EmptyHero() *Hero {
	return &Hero{}
}

func (h *Hero) UpdateHero(Height, Width, SpawnBorderHero float64) *Hero {
	return &Hero{
		MaxHealth: h.Health,
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

func (h *Hero) String() {
	fmt.Printf("[HERO STATUS] Hero all status: Exp: %f | Health: %.2f | MaxHealth: %.2f | Weapon: %s | Weapon Damage: %f - %f | Killed: %d | X-coordination: %f | Y-coordination: %f\n",
		h.Exp, h.Health, h.MaxHealth, h.Weapon.Name, h.Weapon.MinDamageTake, h.Weapon.MaxDamageTake, h.Killed, h.X, h.Y,
	)
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

func (h Hero) IsAlive() bool {
	return h.Health >= 0
}

func (h *Hero) PlusHealth(value float64) {
	if h.MaxHealth <= h.Health+value {
		value = h.MaxHealth - h.Health
	}
	h.Health += value
	fmt.Println("[HEALTH] Boosted for: ", value)
}

func (h *Hero) PlusWeaponDamage(value float64) {
	h.Weapon.MinDamageTake += value
	h.Weapon.MaxDamageTake += value
}

func (h *Hero) Ability(name string, value float64) {
	posibleability := map[string]func(float64){
		"Health": h.PlusHealth,
		"Weapon": h.PlusWeaponDamage,
	}

	posibleability[name](value)
}
