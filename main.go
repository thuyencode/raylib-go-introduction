package main

import (
	"fmt"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/thuyencode/raylib-go-introduction/pkg/helper"
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

	const spaceshipSpeed float32 = 200.0

	for !rl.WindowShouldClose() {
		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		deltaTime := rl.GetFrameTime()
		rl.SetTargetFPS(targetFPS)

		spaceshipDirection.X = float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyRight)) - helper.BoolToInt(rl.IsKeyDown(rl.KeyLeft)))
		spaceshipDirection.Y = float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyDown)) - helper.BoolToInt(rl.IsKeyDown(rl.KeyUp)))
		spaceshipDirection = spaceshipDirection.Normalize()

		spaceshipPos.X += spaceshipDirection.X * spaceshipSpeed * deltaTime
		spaceshipPos.Y += spaceshipDirection.Y * spaceshipSpeed * deltaTime

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(10, 10)
		rl.DrawText(fmt.Sprintf("Spaceship pos: %.0f %.0f", spaceshipPos.X, spaceshipPos.Y), 10, 30, 20, rl.DarkGreen)
		rl.DrawTextureV(spaceshipTexture, spaceshipPos, rl.White)

		rl.EndDrawing()
	}
}
