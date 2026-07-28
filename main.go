package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/thuyencode/raylib-go-introduction/pkg/helper"
)

const WINDOW_WIDTH = 700
const WINDOW_HEIGHT = 700
const SPEED float32 = 200.0
const NUM_OF_CIRCLES = 100
const NUM_OF_COLORS = 5

type CircleItem struct {
	pos    rl.Vector2
	radius float32
	color  rl.Color
}

func main() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Movement in Raylib")
	defer rl.CloseWindow()

	var colors = [NUM_OF_COLORS]rl.Color{rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Orange}
	var circles [NUM_OF_CIRCLES]CircleItem

	for i := range NUM_OF_CIRCLES {
		circles[i].pos = rl.NewVector2(helper.RandFloat32(2000, -2000), helper.RandFloat32(1000, -1000))
		circles[i].radius = helper.RandFloat32(200, 50)
		circles[i].color = colors[rand.IntN(NUM_OF_COLORS)]
	}

	direction := rl.NewVector2(0, 0)
	playerPos := rl.NewVector2(WINDOW_WIDTH/2, WINDOW_HEIGHT/2)
	// The offset value should be the center coord of the window,
	// which is the current `playerPos`. That's why
	camera := rl.NewCamera2D(playerPos, playerPos, 0, 1)

	for !rl.WindowShouldClose() {
		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		deltaTime := rl.GetFrameTime()
		rl.SetTargetFPS(targetFPS)

		direction.X = float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyRight))) - float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyLeft)))
		direction.Y = float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyDown))) - float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyUp)))
		direction = direction.Normalize()

		playerPos.X += direction.X * deltaTime * SPEED
		playerPos.Y += direction.Y * deltaTime * SPEED

		rotateDirection := float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyS)) - helper.BoolToInt(rl.IsKeyDown(rl.KeyA)))
		zoomDirection := float32(helper.BoolToInt(rl.IsKeyDown(rl.KeyW)) - helper.BoolToInt(rl.IsKeyDown(rl.KeyQ)))
		camera.Target = playerPos
		camera.Rotation += rotateDirection * deltaTime * (SPEED / 2)
		camera.Zoom += zoomDirection * deltaTime * (SPEED / 100)
		camera.Zoom = max(0.1, min(5, camera.Zoom))

		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)

		for _, circle := range circles {
			rl.DrawCircleV(circle.pos, circle.radius, circle.color)
		}

		rl.DrawCircleV(playerPos, 10, rl.Black)

		rl.EndMode2D()
		rl.EndDrawing()
	}
}
