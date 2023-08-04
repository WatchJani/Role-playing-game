package main

import (
	load "github.com/WatchJani/Role-playing-game/helper/game_init"
)

func main() {
	game := load.GameInit()

	game.Hero.String()
}
