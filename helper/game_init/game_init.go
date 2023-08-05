package game_init

import (
	e "github.com/WatchJani/Role-playing-game/model/enemy"
	g "github.com/WatchJani/Role-playing-game/model/game"
	h "github.com/WatchJani/Role-playing-game/model/hero"
	file "github.com/WatchJani/Role-playing-game/model/reader"
)

func GameInit() (*g.Game, *[]e.Enemy) {
	gameSettings := g.NewSettings()
	file.NewRead("./config/game_settings.json", gameSettings)

	gameBlueprint := g.NewGameBlueprint()
	file.NewRead("./config/game_blueprint.json", gameBlueprint)

	gameHero := h.EmptyHero()
	file.NewRead("./config/hero.json", gameHero)

	enemy := e.EmptyEnemies()
	file.NewRead("./config/enemies.json", enemy)

	return g.NewGame(*gameSettings, *gameBlueprint, *gameHero), enemy
}
