package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Star struct {
	Pos   rl.Vector2
	Scale float32
}

func NewStar(pos rl.Vector2, scale float32) *Star {
	return &Star{pos, scale}
}
