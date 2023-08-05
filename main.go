package main

import (
	load "github.com/WatchJani/Role-playing-game/helper/game_init"
)

func main() {
	game, enemies := load.GameInit()
	game.Hero.String()

	for _, enemy := range *enemies {
		enemy.String()
	}

}
