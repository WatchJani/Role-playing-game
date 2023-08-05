package enemy

import (
	"fmt"

	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	g "github.com/WatchJani/Role-playing-game/model/game"
	h "github.com/WatchJani/Role-playing-game/model/hero"
	s "github.com/WatchJani/Role-playing-game/model/skeleton"
)

type Enemy struct {
	Class    string  `json:"Class"`
	Movement float64 `json:"Movement"`
	s.Skeleton
}

func EmptyEnemies() *[]Enemy {
	return &[]Enemy{}
}

func (e *Enemy) Spawn(game g.Game) {
	e.X = m.GeneratePosition(game.Radius, game.Width, game.X)
	e.Y = m.GeneratePosition(game.Radius, game.Height, game.Y)
}

func TestEnemyNew(x, y, movement, health float64) Enemy {
	return Enemy{
		Skeleton: s.Skeleton{
			X:      x,
			Y:      y,
			Health: health,
		},
		Movement: movement,
	}
}

func (e Enemy) TakeDame() float64 {
	return m.RandomNumber(e.Weapon.MinDamageTake, e.Weapon.MaxDamageTake)
}

func (e Enemy) GetEXP() float64 {
	return e.Exp
}

func (e *Enemy) GetDamage(damage float64) float64 {
	e.Health = -damage

	return e.Health
}

func (e Enemy) String() {
	fmt.Printf("Class Name: %s | Exp: %f | Health: %f | Movement: %f | Weapon: %s | X-coordination: %f | Y-coordination: %f\n", e.Class, e.Exp, e.Health, e.Movement, e.Weapon.Name, e.X, e.Y)
}

func (e *Enemy) Move(hero h.Hero) float64 {
	d := m.Distance(e, hero)

	if d <= e.Movement {
		return d
	}

	ABx, ABy := m.NewVector(e, hero)
	NewX, NewY := m.NormalizationVector(ABx, ABy, d)

	e.X, e.Y = e.GetX()+e.Movement*NewX, e.GetY()+e.Movement*NewY

	return m.Distance(e, hero)
}

func (e Enemy) IsAlive() bool {
	return e.Health >= 0
}
