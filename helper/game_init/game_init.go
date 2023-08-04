package helper

import (
	g "github.com/WatchJani/Role-playing-game/model/game"
	file "github.com/WatchJani/Role-playing-game/model/reader"
)

func GameInit() *g.Game {
	gameSettings := g.NewSettings()
	file.NewRead("./config/game_settings.json", gameSettings)

	gameBlueprint := g.NewGameBlueprint()
	file.NewRead("./config/game_blueprint.json", gameBlueprint)

	return g.NewGame(*gameSettings, *gameBlueprint)
}
