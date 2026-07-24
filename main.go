package main

import (
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH = 700
const WINDOW_HEIGHT = 310

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Movement in Raylib")
	defer rl.CloseWindow()

	spaceshipPath := path.Join("assets", "spaceship.png")
	spaceshipTexture := rl.LoadTexture(spaceshipPath)
	spaceshipPosX := float32(WINDOW_WIDTH/2 - spaceshipTexture.Width/2)
	spaceshipPosY := float32(WINDOW_HEIGHT/2 - spaceshipTexture.Height/2)
	spaceshipPos := rl.NewVector2(spaceshipPosX, spaceshipPosY)
	spaceshipDirection := rl.NewVector2(1, 1)
	var spaceshipSpeed float32 = 50.0

	for !rl.WindowShouldClose() {
		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		deltaTime := rl.GetFrameTime()
		rl.SetTargetFPS(targetFPS)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(10, 10)
		rl.DrawTextureV(spaceshipTexture, spaceshipPos, rl.White)

		// Check for collision
		if WINDOW_WIDTH-spaceshipPos.X > float32(spaceshipTexture.Width) {
			spaceshipPos.X += deltaTime * spaceshipSpeed * spaceshipDirection.X
		}

		if WINDOW_HEIGHT-spaceshipPos.Y > float32(spaceshipTexture.Height) {
			spaceshipPos.Y += deltaTime * spaceshipSpeed * spaceshipDirection.Y
		}

		rl.EndDrawing()
	}
}
