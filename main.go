package main

import (
	"fmt"
	"time"

	load "github.com/WatchJani/Role-playing-game/helper/game_init"
	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	"github.com/WatchJani/Role-playing-game/model/enemy"
)

func main() {
	game, hero, Enemies := load.GameInit()

	var enemies []enemy.Enemy = make([]enemy.Enemy, game.NumEnemies)

GameLoop:
	for game.Spawn < game.NumEnemies {
		enemies[game.Spawn] = (*Enemies)[m.GetNumber(9)]
		enemies[game.Spawn].String()

		for kill, spawn := hero.Killed, game.Spawn; kill <= spawn; kill++ {
			d := enemies[spawn].Move(*hero)

			fmt.Printf("Spawn Point: %f | Radius: %f | Between %f\n", d, game.Radius, d-game.Radius)

			fmt.Printf("Enemy position %f | My position %f | isSame: %t \n",
				d, m.Distance(enemies[spawn].GetX(), enemies[spawn].GetY(), hero.GetX(), hero.Y), m.Distance(enemies[spawn].GetX(), enemies[spawn].GetY(), hero.GetX(), hero.GetY()) == d,
			)

			if !hero.IsAlive() {
				fmt.Println("End Game")
				break GameLoop
			}

			if d < enemies[spawn].Weapon.PowerRange {
				hero.TakeDamage(enemies[spawn].TakeDame())
				fmt.Println(hero.Health)
			}
		}

		enemies[hero.Killed].GetDamage(hero.Damage())

		if !enemies[hero.Killed].IsAlive() {
			hero.SetEXP(enemies[hero.Killed].GetEXP())
			fmt.Println("Game EXP", hero.Exp)
			hero.Killed++
		}

		// if game.Hero.Exp > game.Lvl {
		// 	game.Boosted()
		// 	fmt.Scanln(&answer)
		// 	game.BoostLvl()
		// }

		game.Spawn++
		fmt.Println("spawned:", game.Spawn)

		time.Sleep(500 * time.Millisecond)
	}
}
