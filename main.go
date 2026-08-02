package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	et "github.com/thuyencode/raylib-go-introduction/pkg/entity"
)

const WINDOW_WIDTH = 960
const WINDOW_HEIGHT = 540
const SPEED = 300.0
const BLOCK_SIZE = 50

var LEVEL_MAP = []string{
	"1111111111111111111",
	"1010000000000000001",
	"1010000000001111111",
	"1000000000000000111",
	"1000000200000000011",
	"1000000000000100001",
	"1000000000000100001",
	"1001100000000100001",
	"1001100000000100001",
	"1001100000000100001",
	"1111111111111111111",
}

func main() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Collision in Go and Raylib")
	defer rl.CloseWindow()

	player := et.NewPlayer(rl.NewRectangle(0, 0, 30, 30), rl.Red, SPEED)
	player.SetPosToCenter(WINDOW_WIDTH, WINDOW_HEIGHT)

	level := et.NewLevel(LEVEL_MAP, rl.Gray, BLOCK_SIZE)

	for !rl.WindowShouldClose() {
		monitor := rl.GetCurrentMonitor()
		targetFPS := int32(rl.GetMonitorRefreshRate(monitor))
		rl.SetTargetFPS(targetFPS)

		player.Update()
		level.CheckCollision(et.XAxis, &player)
		level.CheckCollision(et.YAxis, &player)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		level.Draw()
		player.Draw()
		rl.DrawFPS(10, 10)
		player.DrawPlayerPos(10, 30, 20, rl.DarkBlue)

		rl.EndDrawing()
	}
}
