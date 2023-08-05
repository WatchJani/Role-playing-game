package game

import (
	"fmt"
	"math"

	h "github.com/WatchJani/Role-playing-game/model/hero"
)

type Game struct {
	Radius float64
	Spawn  int
	*h.Hero
	Settings
	GameBlueprint
}

type Settings struct {
	SpawnBorderHero float64 `json:"SpawnBorderHero"`
	DynamicRadius   float64 `json:"DynamicRadius"`
	MaxRadius       float64 `json:"MaxRadius"`
	StartLvlPoint   float64 `json:"StartLvlPoint"`
	LvlBoost        float64 `json:"LvlBoost"`
}

type GameBlueprint struct {
	Width  float64 `json:"Width"`
	Height float64 `json:"Height"`
	Lvl    float64 `json:"Lvl"`
}

func NewGame() *Game {
	return &Game{}
}

func (game *Game) Update(hero h.Hero) {
	var (
		height = game.Height
		width  = game.Width
	)

	radius := math.Min(height, width) / game.DynamicRadius

	if radius > game.MaxRadius {
		radius = game.MaxRadius
	}

	height = math.Min(height, 100)
	width = math.Min(game.Width, 100)

	game.Radius = radius
	game.Hero = hero.UpdateHero(height, width, game.SpawnBorderHero)
}

func (g *Game) BoostLvl() {
	g.Lvl *= g.LvlBoost
}

func (g Game) String() {
	fmt.Printf("Radius: %f | Spawn: %d | Width: %f | Height: %f | Lvl: %f | BorderHero: %f | DynamicRadius: %f | MaxRadius: %f | StartLvlPoint: %f | LvlBoost: %f\n",
		g.Radius, g.Spawn, g.Width, g.Height, g.Lvl, g.SpawnBorderHero, g.DynamicRadius, g.MaxRadius, g.StartLvlPoint, g.LvlBoost,
	)
}
