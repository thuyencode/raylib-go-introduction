package main

import (
	"fmt"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	h "github.com/thuyencode/raylib-go-introduction/pkg/helper"
)

const WINDOW_WIDTH = 960
const WINDOW_HEIGHT = 540
const ANIMATION_FRAME_NUM = 8
const SPEED = 5

func main() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Animation in Go and Raylib")
	defer rl.CloseWindow()

	animationFrames := make([]rl.Texture2D, ANIMATION_FRAME_NUM)

	for i := range ANIMATION_FRAME_NUM {
		animationFrames[i] = rl.LoadTexture(path.Join("assets", "animation", fmt.Sprintf("%d.png", i)))
	}

	var animationFrameIndex float32 = 0

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		animationFrameIndex += SPEED * dt

		monitor := rl.GetCurrentMonitor()
		targetFPS := int32(rl.GetMonitorRefreshRate(monitor))
		rl.SetTargetFPS(targetFPS)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(10, 10)

		frame := animationFrames[int(animationFrameIndex)%len(animationFrames)]
		posX, posY := h.GetCenterPos(WINDOW_WIDTH, WINDOW_HEIGHT, frame.Width, frame.Height)
		rl.DrawTexture(frame, posX, posY, rl.White)

		rl.EndDrawing()
	}
}
