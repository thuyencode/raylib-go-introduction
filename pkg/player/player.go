package player

import (
	"fmt"
	"image/color"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	h "github.com/thuyencode/raylib-go-introduction/pkg/helper"
)

type Player struct {
	Pos       rl.Vector2
	Direction rl.Vector2
	Texture   rl.Texture2D
	Speed     float32
}

func NewPlayer(pos rl.Vector2, assetFilePath string, speed float32) Player {
	texture := rl.LoadTexture(path.Join("assets", assetFilePath))
	direction := rl.Vector2{X: 0, Y: 0}
	return Player{pos, direction, texture, speed}
}

func (player *Player) SetPosToCenter(windowWidth int32, windowHeight int32) {
	player.Pos.X = float32(windowWidth/2 - player.Texture.Width/2)
	player.Pos.Y = float32(windowHeight/2 - player.Texture.Height/2)

}

func (player *Player) Draw() {
	rl.DrawTextureV(player.Texture, player.Pos, rl.White)
}

func (player *Player) Update() {
	player.Direction.X = float32(h.BoolToInt(rl.IsKeyDown(rl.KeyRight)) - h.BoolToInt(rl.IsKeyDown(rl.KeyLeft)))
	player.Direction.Y = float32(h.BoolToInt(rl.IsKeyDown(rl.KeyDown)) - h.BoolToInt(rl.IsKeyDown(rl.KeyUp)))
	player.Direction = player.Direction.Normalize()

	dt := rl.GetFrameTime()

	player.Pos.X += player.Direction.X * dt * player.Speed
	player.Pos.Y += player.Direction.Y * dt * player.Speed
}

func (player *Player) DrawPlayerPos(posX int32, posY int32, fontSize int32, col color.RGBA) {
	rl.DrawText(fmt.Sprintf("Player POS: %.0f %.0f", player.Pos.X, player.Pos.Y), posX, posY, fontSize, col)
}
