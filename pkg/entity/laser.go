package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/thuyencode/raylib-go-introduction/pkg/configs"
)

type Laser struct {
	Sprite
}

func NewLaser(texture rl.Texture2D, pos rl.Vector2) *Laser {
	return &Laser{
		Sprite: *NewSprite(texture, pos, c.LASER_SPEED, rl.NewVector2(0, -1)),
	}
}

func (laser *Laser) GetRect() rl.Rectangle {
	return rl.NewRectangle(
		laser.Pos.X,
		laser.Pos.Y,
		float32(laser.Texture.Width),
		float32(laser.Texture.Height))
}
