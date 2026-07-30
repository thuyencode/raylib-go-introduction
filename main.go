package main

import (
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH = 700
const WINDOW_HEIGHT = 700

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Audio in Raylib")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	laserSound := rl.LoadSound(path.Join("assets", "laser.wav"))
	musicStream := rl.LoadMusicStream(path.Join("assets", "music.wav"))

	rl.PlayMusicStream(musicStream)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(musicStream)

		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		rl.SetTargetFPS(targetFPS)

		rl.BeginDrawing()
		rl.DrawFPS(10, 10)
		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeyP) {
			rl.PlaySound(laserSound)
		}

		rl.EndDrawing()
	}
}
