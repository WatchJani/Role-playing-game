package game_init

import (
	e "github.com/WatchJani/Role-playing-game/model/enemy"
	g "github.com/WatchJani/Role-playing-game/model/game"
	h "github.com/WatchJani/Role-playing-game/model/hero"
	i "github.com/WatchJani/Role-playing-game/model/item"
	file "github.com/WatchJani/Role-playing-game/model/reader"
)

func GameInit() (*g.Game, *h.Hero, *[]e.Enemy, *[]i.Item) {
	game := g.NewGame()
	file.NewRead("./config/game.json", game)

	gameHero := h.EmptyHero()
	file.NewRead("./config/hero.json", gameHero)

	game.Update(*gameHero)

	enemy := e.EmptyEnemies()
	file.NewRead("./config/enemies.json", enemy)

	items := i.EmptyItems()
	file.NewRead("./config/items.json", items)

	return game, game.Hero, enemy, items
}
