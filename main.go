package main

import (
	"fmt"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 400, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	spaceshipPath := path.Join("assets", "spaceship.png")
	spaceshipImage := rl.LoadImage(spaceshipPath)
	spaceshipPosX := 400 - (spaceshipImage.Width / 2)
	spaceshipPosY := 200 - (spaceshipImage.Height / 2)
	spaceshipTexture := rl.LoadTextureFromImage(spaceshipImage)

	cowboyPath := path.Join("assets", "animation", "0.png")
	cowboyImage := rl.LoadImage(cowboyPath)
	cowboyPosX := 800 - cowboyImage.Width
	cowboyPosY := 400 - cowboyImage.Height
	rl.ImageColorInvert(cowboyImage)
	cowboyTexture := rl.LoadTextureFromImage(cowboyImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 10, 10, 24, rl.Green)

		rl.DrawTexture(spaceshipTexture, spaceshipPosX, spaceshipPosY, rl.White)
		rl.DrawTexture(cowboyTexture, cowboyPosX, cowboyPosY, rl.White)

		rl.DrawLineEx(rl.NewVector2(100, 100), rl.NewVector2(300, 100), 10, rl.Blue)

		rl.EndDrawing()
	}
}
