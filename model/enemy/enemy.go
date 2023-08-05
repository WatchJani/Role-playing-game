package enemy

import (
	s "github.com/WatchJani/Role-playing-game/model/skeleton"
)

type Enemy struct {
	s.Skeleton
	Movement float64
	Class    string
}

// func (g *Game) NewEnemy(movement, exp, health float64, class string, weapon Weapon) Enemy {
// 	return Enemy{
// 		Skeleton: NewSkeletonEnemy(g.Height, g.Width, g.Radius, exp, health, weapon, g.Hero),
// 		Movement: movement,
// 		Class:    class,
// 	}
// }

// func TestEnemyNew(x, y, movement, health float64) Enemy {
// 	return Enemy{
// 		Skeleton: Skeleton{
// 			X:      x,
// 			Y:      y,
// 			Health: health,
// 			// Weapon: make(map[string]Weapon),
// 		},
// 		Movement: movement,
// 	}
// }

// func (e Enemy) TakeDame() float64 {
// 	return helper.RandomNumber(e.Weapon.MinDamageTake, e.Weapon.MaxDamageTake)
// }

// func (e Enemy) GetEXP() float64 {
// 	return e.Exp
// }

// func (e *Enemy) GetDamage(damage float64) float64 {
// 	e.Health = -damage

// 	return e.Health
// }

// func (e Enemy) String() {
// 	fmt.Printf("Class Name: %s | Exp: %f | Health: %f | Movement: %f | Weapon: %s\n", e.Class, e.Exp, e.Health, e.Movement, e.Weapon.Name)
// }

// func (e *Enemy) Move(hero Hero) float64 {
// 	d := helper.Distance(e, hero)

// 	if d <= e.Movement {
// 		return d
// 	}

// 	ABx, ABy := helper.NewVector(e, hero)
// 	NewX, NewY := helper.NormalizationVector(ABx, ABy, d)

// 	e.X, e.Y = e.GetX()+e.Movement*NewX, e.GetY()+e.Movement*NewY

// 	return helper.Distance(e, hero)
// }

// func (e Enemy) IsAlive() bool {
// 	return e.Health >= 0
// }
