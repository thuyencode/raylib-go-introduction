package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/thuyencode/raylib-go-introduction/pkg/configs"
)

type Sprite struct {
	Texture           rl.Texture2D
	Pos               rl.Vector2
	Speed             int
	Direction         rl.Vector2
	ShouldBeDiscarded bool
	CollisionRadius   float32
}

func NewSprite(texture rl.Texture2D, pos rl.Vector2, speed int, direction rl.Vector2) *Sprite {
	return &Sprite{
		Texture:           texture,
		Pos:               pos,
		Speed:             speed,
		Direction:         direction,
		ShouldBeDiscarded: false,
		CollisionRadius:   float32(texture.Height) / 2,
	}
}

func (sprite *Sprite) GetCenterPos() rl.Vector2 {
	return rl.NewVector2(
		sprite.Pos.X+float32(sprite.Texture.Width/2),
		sprite.Pos.Y+float32(sprite.Texture.Height/2),
	)
}

func (sprite *Sprite) MoveToCenterOfWindow() {
	sprite.Pos.X = float32(c.WINDOW_WIDTH/2 - sprite.Texture.Width/2)
	sprite.Pos.Y = float32(c.WINDOW_HEIGHT/2 - sprite.Texture.Height/2)
}

func (sprite *Sprite) CheckForDiscard() {
	if sprite.Pos.Y < -300 {
		sprite.ShouldBeDiscarded = true
	}
	if sprite.Pos.Y > c.WINDOW_HEIGHT+300 {
		sprite.ShouldBeDiscarded = true
	}
}

func (sprite *Sprite) Move(deltaTime float32) {
	sprite.Pos.X += sprite.Direction.X * float32(sprite.Speed) * deltaTime
	sprite.Pos.Y += sprite.Direction.Y * float32(sprite.Speed) * deltaTime
}

func (sprite *Sprite) Update(deltaTime float32) {
	sprite.Move(deltaTime)
	sprite.CheckForDiscard()
}

func (sprite *Sprite) Draw() {
	rl.DrawTextureV(sprite.Texture, sprite.Pos, rl.White)
}
