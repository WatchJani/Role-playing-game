package main

import (
	load "github.com/WatchJani/Role-playing-game/helper/game_init"
)

func main() {
	game, hero, enemies := load.GameInit()
	hero.String()

	game.String()

	for _, enemy := range *enemies {
		enemy.Spawn(*game)

		enemy.String()
	}
}
