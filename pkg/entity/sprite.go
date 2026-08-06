package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	h "github.com/thuyencode/raylib-go-introduction/pkg/helper"
)

type Sprite struct {
	rl.Rectangle
	rl.Color
}

func NewSprite(pos rl.Vector2, size [2]float32) *Sprite {
	return &Sprite{
		Color:     rl.White,
		Rectangle: rl.NewRectangle(pos.X, pos.Y, size[0], size[1]),
	}
}

var colors = []rl.Color{rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Orange}

func (sprite *Sprite) RandomizeColor() {
	randColor := h.RandItem(colors)

	for randColor.A == sprite.A && randColor.B == sprite.B && randColor.G == sprite.G && randColor.R == sprite.R {
		randColor = h.RandItem(colors)
	}

	sprite.Color = randColor
}

func (sprite *Sprite) SetPos(pos rl.Vector2) {
	sprite.X = pos.X
	sprite.Y = pos.Y
}

func (sprite *Sprite) Draw() {
	rl.DrawRectangleRec(sprite.Rectangle, sprite.Color)
}
