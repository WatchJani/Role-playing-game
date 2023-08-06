package main

import (
	"testing"

	m "github.com/WatchJani/Role-playing-game/helper/game_math"
	e "github.com/WatchJani/Role-playing-game/model/enemy"
	h "github.com/WatchJani/Role-playing-game/model/hero"
)

func TestNormalizationVector(t *testing.T) {
	data := []struct {
		x          float64
		y          float64
		distance   float64
		expected_X float64
		expected_Y float64
	}{
		{10, 0, 10, 1, 0},
	}

	for _, test := range data {
		if x, y := m.NormalizationVector(test.x, test.y, test.distance); x != test.expected_X || y != test.expected_Y {
			t.Errorf("Test Failed: (%f, %f) inputted, (%f,%f) expected, received: (%f, %f)", test.x, test.y, test.expected_X, test.expected_Y, x, y)
		}
	}
}

func TestDistance(t *testing.T) {
	enemy := e.TestEnemyNew(0, 2, 10, 10)
	hero := h.TestHeroNew(10, 2, 100)

	if d := m.Distance(enemy.GetX(), enemy.GetY(), hero.GetX(), hero.GetY()); d != 10 {
		t.Errorf("Test Failed: %f, expected: %f", d, 10.0)
	}
}

func TestNewVector(t *testing.T) {
	enemy := e.TestEnemyNew(0, 2, 10, 10)
	hero := h.TestHeroNew(10, 2, 100)

	if x, y := m.NewVector(enemy.GetX(), enemy.GetY(), hero.GetX(), hero.GetY()); x != 10 || y != 0 {
		t.Errorf("Test Failed: (%f, %f), expected: (%f, %f)", x, y, 10.0, 0.0)
	}

}

func TestMove(t *testing.T) {
	enemy := e.TestEnemyNew(0, 2, 2, 10)
	hero := h.TestHeroNew(10, 2, 100)

	enemy.Move(hero)

	if x, y := enemy.X, enemy.Y; x != 2 || y != 2 {
		t.Errorf("Test Failed: (%f, %f), expected: (%f, %f)", x, y, 2.0, 2.0)
	}
}
