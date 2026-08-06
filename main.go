package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
	e "github.com/thuyencode/raylib-go-introduction/pkg/entity"
)

const WINDOW_WIDTH = 960
const WINDOW_HEIGHT = 540

func main() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Collision in Go and Raylib")
	defer rl.CloseWindow()

	sprite := e.NewSprite(rl.NewVector2(100, 100), [2]float32{30, 30})

	timer1 := e.NewTimer(1.5, true, true, sprite.RandomizeColor)

	timer2 := e.NewTimer(4, true, true, func() {
		sprite.SetPos(rl.NewVector2(float32(rand.IntN(WINDOW_WIDTH)), float32(rand.IntN(WINDOW_HEIGHT))))
	})

	for !rl.WindowShouldClose() {
		timer1.Update()
		timer2.Update()

		monitor := rl.GetCurrentMonitor()
		targetFPS := int32(rl.GetMonitorRefreshRate(monitor))
		rl.SetTargetFPS(targetFPS)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(10, 10)

		sprite.Draw()

		rl.EndDrawing()
	}
}
