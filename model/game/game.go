package game

import (
	"fmt"
	"math"

	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	h "github.com/WatchJani/Role-playing-game/model/hero"
	i "github.com/WatchJani/Role-playing-game/model/item"
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
	Width      float64 `json:"Width"`
	Height     float64 `json:"Height"`
	Lvl        float64 `json:"Lvl"`
	NumEnemies int     `json:"EnemiesNumber"`
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
	fmt.Printf("Radius: %f | Spawn: %d | Width: %f | Height: %f | Lvl: %f | BorderHero: %f | DynamicRadius: %f | MaxRadius: %f | StartLvlPoint: %f | LvlBoost: %f | NumEnemies %d\n",
		g.Radius, g.Spawn, g.Width, g.Height, g.Lvl, g.SpawnBorderHero, g.DynamicRadius, g.MaxRadius, g.StartLvlPoint, g.LvlBoost, g.NumEnemies,
	)
}

func (g Game) Boosted(items *[]i.Item) []int {
	boostLen := len(*items)
	opt1, opt2, opt3 := m.GetNumber(boostLen), m.GetNumber(boostLen), m.GetNumber(boostLen)

	fmt.Printf("0)\n %s \n %s \n\n1)\n %s \n %s\n\n2)\n %s \n %s \n",
		(*items)[opt1].Name, (*items)[opt2].Description,
		(*items)[opt2].Name, (*items)[opt3].Description,
		(*items)[opt3].Name, (*items)[opt3].Description,
	)

	return []int{opt1, opt2, opt3}
}

func (g Game) InputChecker(answer int) bool {
	for i := 0; i < 3; i++ {
		if answer == i {
			return false
		}
	}

	return true
}
