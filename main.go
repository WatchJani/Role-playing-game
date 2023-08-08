package main

import (
	"fmt"
	"time"

	load "github.com/WatchJani/Role-playing-game/helper/game"
)

func main() {
	//game load
	game, hero, Enemies, items := load.GameInit()
	//hero status
	hero.String()

	//all enemies on battlefield, which have been created and which will be created
	enemies := load.NewBattlefield(game)

	//game loop
	for load.InGame(game) {
		//spawn new enemy
		load.EnemySpawn(enemies, Enemies, game)
		// enemy.ThisEnemy() //just for print

		//make enemy move
		if load.EnemyMove(hero, game, enemies) {
			//game status end
			fmt.Println("End Game")
			return
		}

		//hero make damage
		load.HeroMakeDamage(enemies, hero)

		//hero kill someone
		load.HeroKilled(enemies, hero)

		//hero lvl up
		load.LvlUp(hero, game, items)

		//game increase spawn
		game.NewSpawn()

		//make simulation time for move
		time.Sleep(500 * time.Millisecond)
	}

	//game status end
	fmt.Println("You are Winner!")
}
