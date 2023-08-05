package main

import (
	"fmt"

	load "github.com/WatchJani/Role-playing-game/helper/game_init"
)

func main() {
	game, enemies := load.GameInit()
	game.Hero.String()

	(*enemies)[0].Spawn(*game)

	fmt.Println((*enemies)[0])
}
