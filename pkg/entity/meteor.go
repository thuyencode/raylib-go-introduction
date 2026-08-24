package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/thuyencode/raylib-go-introduction/pkg/configs"
	u "github.com/thuyencode/raylib-go-introduction/pkg/util"
)

type Meteor struct {
	Sprite
	Rotation float32
	Rect     rl.Rectangle
}

func NewMeteor(texture rl.Texture2D) *Meteor {
	pos := rl.NewVector2(
		u.RandFloat32(c.WINDOW_WIDTH, 0),
		u.RandFloat32(-50, -150))
	speed := u.RandInt(
		c.METEOR_SPEED_RANGE[1],
		c.METEOR_SPEED_RANGE[0])
	direction := rl.NewVector2(u.RandFloat32(0.5, -0.5), 1)

	return &Meteor{
		Sprite:   *NewSprite(texture, pos, speed, direction),
		Rotation: 0,
		Rect:     rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
	}
}

func (meteor *Meteor) GetCenterPos() rl.Vector2 {
	return meteor.Pos
}

func (meteor *Meteor) Update(deltaTime float32) {
	meteor.Sprite.Update(deltaTime)
	meteor.Rotation += 50 * deltaTime
}

func (meteor *Meteor) Draw() {
	dest := rl.NewRectangle(
		meteor.Pos.X,
		meteor.Pos.Y,
		float32(meteor.Texture.Width),
		float32(meteor.Texture.Height))
	origin := rl.NewVector2(
		float32(meteor.Texture.Width)/2,
		float32(meteor.Texture.Height)/2,
	)

	rl.DrawTexturePro(meteor.Texture, meteor.Rect, dest, origin, meteor.Rotation, rl.White)
}
