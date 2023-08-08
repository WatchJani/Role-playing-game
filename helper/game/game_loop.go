package game

import (
	"fmt"

	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	"github.com/WatchJani/Role-playing-game/model/enemy"
	"github.com/WatchJani/Role-playing-game/model/game"
	"github.com/WatchJani/Role-playing-game/model/hero"
	"github.com/WatchJani/Role-playing-game/model/item"
)

// add new enemy on the battlefield
func EnemySpawn(allEnemy, kindOfEnemy *[]enemy.Enemy, game *game.Game) *enemy.Enemy {
	//add new enemy on battlefield
	(*allEnemy)[game.Spawn] = (*kindOfEnemy)[m.GetNumber(9)]
	//set random coordination of spawned enemies
	(*allEnemy)[game.Spawn].Spawn(*game)
	return &(*allEnemy)[game.Spawn]
}

// is the end of lvl
func InGame(game *game.Game) bool {
	return game.Spawn < game.NumEnemies
}

func HeroTakeDamage(distance float64, enemy *enemy.Enemy, hero *hero.Hero) {
	if distance <= (*enemy).Weapon.PowerRange {
		enemyDamage := (*enemy).TakeDame()
		hero.TakeDamage(enemyDamage)
		fmt.Printf("[HERO HEALTH]: %f [DAMAGE TAKE] %f\n", hero.Health, enemyDamage)
	}
}

func HeroMakeDamage(enemies *[]enemy.Enemy, hero *hero.Hero) {
	(*enemies)[hero.Killed].GetDamage(hero.Damage())
}

func EnemyMove(hero *hero.Hero, game *game.Game, enemies *[]enemy.Enemy) bool {

	for kill, spawn := hero.Killed, game.Spawn; kill <= spawn; kill++ {
		//if hero dead, game is over
		if !hero.IsAlive() {
			return true
		}

		//make move and return new distance
		distance := (*enemies)[spawn].Move(*hero)

		//make damage on hero if range of enemy weapon in radius
		HeroTakeDamage(distance, &(*enemies)[spawn], hero)
	}
	return false
}

func HeroKilled(enemies *[]enemy.Enemy, hero *hero.Hero) {
	if !(*enemies)[hero.Killed].IsAlive() {
		hero.SetEXP((*enemies)[hero.Killed].GetEXP())
		fmt.Println("[GAME EXP] EXP: ", hero.Exp)
		hero.Killed++
	}
}

func LvlUp(hero *hero.Hero, game *game.Game, items *[]item.Item) {
	if hero.Exp > game.Lvl {
		for {
			var answer int
			fmt.Println()
			opt := game.Boosted(items)
			fmt.Println()
			fmt.Printf("Your answer: ")
			fmt.Scanln(&answer)

			if game.InputChecker(answer) {
				fmt.Printf("===============Input is not valid!=================\n\n\n")
				continue
			}

			hero.Ability((*items)[opt[answer]].Name, (*items)[opt[answer]].Value)
			hero.String()
			game.BoostLvl()

			break
		}
	}
}

func NewBattlefield(game *game.Game) *[]enemy.Enemy {
	var enemies []enemy.Enemy = make([]enemy.Enemy, game.NumEnemies)

	return &enemies
}
