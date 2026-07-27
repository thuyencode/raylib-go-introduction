package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH = 700
const WINDOW_HEIGHT = 700
const SPEED float32 = 200.0

func main() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Movement in Raylib")
	defer rl.CloseWindow()

	player := rl.NewRectangle(0, 0, 100, 100)
	rectangle := rl.NewRectangle(WINDOW_WIDTH-100, 0, 200, 200)

	for !rl.WindowShouldClose() {
		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		rl.SetTargetFPS(targetFPS)

		player.X = rl.GetMousePosition().X
		player.Y = rl.GetMousePosition().Y
		overlapRec := rl.GetCollisionRec(player, rectangle)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(10, 10)

		rl.DrawRectangleRec(player, rl.Blue)
		rl.DrawRectangleRec(rectangle, rl.Yellow)
		rl.DrawRectangleRec(overlapRec, rl.Red)

		rl.EndDrawing()
	}
}
