package animation

import rl "github.com/gen2brain/raylib-go/raylib"

type Size struct {
	width  int32
	height int32
}

type ExplosionAnimation struct {
	textures          []rl.Texture2D
	size              Size
	pos               rl.Vector2
	index             int
	ShouldBeDiscarded bool
}

func NewExplosionAnimation(textures []rl.Texture2D, pos rl.Vector2) *ExplosionAnimation {
	size := Size{textures[0].Width, textures[0].Height}
	centerPos := rl.NewVector2(pos.X-float32(size.width)/2, pos.Y-float32(size.height)/2)

	return &ExplosionAnimation{
		textures:          textures,
		size:              size,
		pos:               centerPos,
		index:             0,
		ShouldBeDiscarded: false,
	}
}

func (ani *ExplosionAnimation) Draw() {
	rl.DrawTextureV(ani.textures[ani.index], ani.pos, rl.White)
}

func (ani *ExplosionAnimation) Update(deltaTime float32) {
	if ani.index < len(ani.textures)-1 {
		ani.index += int(100 * deltaTime)
	} else {
		ani.ShouldBeDiscarded = true
	}
}
