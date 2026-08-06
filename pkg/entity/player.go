package entity

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	h "github.com/thuyencode/raylib-go-introduction/pkg/helper"
)

type Player struct {
	rl.Rectangle
	Color     rl.Color
	Direction rl.Vector2
	Speed     float32
}

func NewPlayer(rectangle rl.Rectangle, color rl.Color, speed float32) *Player {
	direction := rl.Vector2{X: 0, Y: 0}
	return &Player{rectangle, color, direction, speed}
}

func (player *Player) SetPosToCenter(windowWidth int32, windowHeight int32) {
	player.X = float32(windowWidth/2) - player.Width/2
	player.Y = float32(windowHeight/2) - player.Height/2
}

func (player *Player) Draw() {
	rl.DrawRectangleRec(player.Rectangle, player.Color)
}

func (player *Player) Update() {
	player.Direction.X = float32(h.BoolToInt(rl.IsKeyDown(rl.KeyRight)) - h.BoolToInt(rl.IsKeyDown(rl.KeyLeft)))
	player.Direction.Y = float32(h.BoolToInt(rl.IsKeyDown(rl.KeyDown)) - h.BoolToInt(rl.IsKeyDown(rl.KeyUp)))
	player.Direction = player.Direction.Normalize()

	dt := rl.GetFrameTime()

	player.X += player.Direction.X * dt * player.Speed
	player.Y += player.Direction.Y * dt * player.Speed
}

func (player *Player) DrawPlayerPos(posX int32, posY int32, fontSize int32, col color.RGBA) {
	rl.DrawText(fmt.Sprintf("Player POS: %.0f %.0f", player.X, player.Y), posX, posY, fontSize, col)
}
