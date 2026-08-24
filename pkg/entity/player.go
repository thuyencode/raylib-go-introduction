package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/thuyencode/raylib-go-introduction/pkg/configs"
	h "github.com/thuyencode/raylib-go-introduction/pkg/util"
)

type Player struct {
	Sprite
	SetLaserPos func(pos rl.Vector2)
}

func NewPlayer(texture rl.Texture2D, pos rl.Vector2, laserCallback func(pos rl.Vector2)) *Player {
	return &Player{
		Sprite:      *NewSprite(texture, pos, c.PLAYER_SPEED, rl.NewVector2(0, 0)),
		SetLaserPos: laserCallback,
	}
}

func (player *Player) Input() {
	player.Direction.X = float32(
		h.BoolToInt(rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)) -
			h.BoolToInt(rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)))
	player.Direction.Y = float32(
		h.BoolToInt(rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS)) -
			h.BoolToInt(rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW)))
	player.Direction = player.Direction.Normalize()

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) {
		player.SetLaserPos(
			rl.NewVector2(
				player.Pos.X+float32(player.Texture.Width/2),
				player.Pos.Y-60,
			),
		)
	}
}

func (player *Player) Constraint() {
	player.Pos.X = rl.Clamp(player.Pos.X, 0, float32(c.WINDOW_WIDTH-player.Texture.Width))
	player.Pos.Y = rl.Clamp(player.Pos.Y, 0, float32(c.WINDOW_HEIGHT-player.Texture.Height))
}

func (player *Player) Update(deltaTime float32) {
	player.Input()
	player.Move(deltaTime)
	player.Constraint()
}
