package game_init

import (
	e "github.com/WatchJani/Role-playing-game/model/enemy"
	g "github.com/WatchJani/Role-playing-game/model/game"
	h "github.com/WatchJani/Role-playing-game/model/hero"
	file "github.com/WatchJani/Role-playing-game/model/reader"
)

func GameInit() (*g.Game, *h.Hero, *[]e.Enemy) {
	game := g.NewGame()
	file.NewRead("./config/game.json", game)

	gameHero := h.EmptyHero()
	file.NewRead("./config/hero.json", gameHero)

	game.Update(*gameHero)

	enemy := e.EmptyEnemies()
	file.NewRead("./config/enemies.json", enemy)

	return game, game.Hero, enemy
}
