package game

import (
	"math"

	h "github.com/WatchJani/Role-playing-game/model/hero"
)

type Settings struct {
	SpawnBorderHero float64 `json:"SpawnBorderHero"`
	DynamicRadius   float64 `json:"DynamicRadius"`
	MaxRadius       float64 `json:"MaxRadius"`
	StartLvlPoint   float64 `json:"StartLvlPoint"`
	LvlBoost        float64 `json:"LvlBoost"`
}

func NewSettings() *Settings {
	return &Settings{}
}

type GameBlueprint struct {
	Width  float64 `json:"Width"`
	Height float64 `json:"Height"`
	Lvl    float64 `json:"Lvl"`
}

type Game struct {
	Radius float64
	Spawn  int
	h.Hero
	Settings
	GameBlueprint
}

func NewGameBlueprint() *GameBlueprint {
	return &GameBlueprint{}
}

func NewGame(settings Settings, gameBlueprint GameBlueprint, hero h.Hero) *Game {
	var (
		height = gameBlueprint.Height
		width  = gameBlueprint.Width
	)

	radius := math.Min(height, width) / settings.DynamicRadius

	if radius > settings.MaxRadius {
		radius = settings.MaxRadius
	}

	height = math.Min(height, 100)
	width = math.Min(gameBlueprint.Width, 100)

	return &Game{
		Radius:        radius,
		Settings:      settings,
		GameBlueprint: gameBlueprint,
		Hero:          hero.UpdateHero(height, width, settings.SpawnBorderHero),
	}
}

// func (g Game) String() {
// 	fmt.Printf("Hero Health: %f \n", g.HeroHealth, g.Exp, g.X, g.Y, g.StartLvlPoint, g.Lvl, g.Radius, g.)
// }
