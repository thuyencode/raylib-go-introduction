package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	p "github.com/thuyencode/raylib-go-introduction/pkg/player"
)

const WINDOW_WIDTH = 700
const WINDOW_HEIGHT = 700
const SPEED = 200.0

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Composition in Go and Raylib")
	defer rl.CloseWindow()

	player := p.NewPlayer(rl.NewVector2(0, 0), "spaceship.png", SPEED)
	player.SetPosToCenter(WINDOW_WIDTH, WINDOW_HEIGHT)

	for !rl.WindowShouldClose() {
		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		rl.SetTargetFPS(targetFPS)

		rl.BeginDrawing()
		rl.DrawFPS(10, 10)
		rl.ClearBackground(rl.RayWhite)

		player.Update()
		player.Draw()
		player.DrawPlayerPos(10, 30, 20, rl.Green)

		rl.EndDrawing()
	}
}
