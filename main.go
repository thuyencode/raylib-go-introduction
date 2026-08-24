package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/thuyencode/raylib-go-introduction/pkg/configs"
	g "github.com/thuyencode/raylib-go-introduction/pkg/game"
)

func main() {
	game := g.NewGame(true, rl.LogError)
	game.SetPlayer(rl.NewVector2(0, 0), game.CreateLaser)
	game.SetTimer(c.METEOR_TIMER_DURATION, true, game.CreateMeteor)

	game.Run()
}
