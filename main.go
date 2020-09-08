package main

import (
	"go-k8s-chaos/characters"
	"go-k8s-chaos/level"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()

	level := level.Level{
		BaseLevel: level.Init(),
	}

	level.Construct()

	player := characters.Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		Level:  level.BaseLevel,
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)

	game.Screen().SetLevel(level)
	game.Start()
}
