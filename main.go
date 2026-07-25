package main

import (
	"fmt"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH = 700
const WINDOW_HEIGHT = 700

func main() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Movement in Raylib")
	defer rl.CloseWindow()

	spaceshipPath := path.Join("assets", "spaceship.png")
	spaceshipTexture := rl.LoadTexture(spaceshipPath)
	spaceshipHalfW := spaceshipTexture.Width / 2
	spaceshipHalfH := spaceshipTexture.Height / 2
	spaceshipPosX := float32(WINDOW_WIDTH/2 - spaceshipHalfW)
	spaceshipPosY := float32(WINDOW_HEIGHT/2 - spaceshipHalfH)
	spaceshipPos := rl.NewVector2(spaceshipPosX, spaceshipPosY)
	spaceshipDirection := rl.NewVector2(0, 0)
	const spaceshipSpeed float32 = 100.0

	for !rl.WindowShouldClose() {
		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		deltaTime := rl.GetFrameTime()
		rl.SetTargetFPS(targetFPS)

		spaceshipDirection.X = 0
		spaceshipDirection.Y = 0

		if rl.IsKeyDown(rl.KeyLeft) {
			spaceshipDirection.X = -1
		}
		if rl.IsKeyDown(rl.KeyRight) {
			spaceshipDirection.X = 1
		}
		if rl.IsKeyDown(rl.KeyUp) {
			spaceshipDirection.Y = -1
		}
		if rl.IsKeyDown(rl.KeyDown) {
			spaceshipDirection.Y = 1
		}

		spaceshipPos.X += spaceshipDirection.X * spaceshipSpeed * deltaTime
		spaceshipPos.Y += spaceshipDirection.Y * spaceshipSpeed * deltaTime
		spaceshipDirection = spaceshipDirection.Normalize()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(10, 10)
		rl.DrawText(fmt.Sprintf("Spaceship pos: %.0f %.0f", spaceshipPos.X, spaceshipPos.Y), 10, 30, 20, rl.DarkGreen)
		rl.DrawTextureV(spaceshipTexture, spaceshipPos, rl.White)

		rl.EndDrawing()
	}
}
