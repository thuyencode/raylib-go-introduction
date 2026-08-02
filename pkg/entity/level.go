package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Level struct {
	Blocks     []rl.Rectangle
	BlockColor rl.Color
	BlockSize  float32
}

type Axis int

const (
	XAxis Axis = iota
	YAxis
)

func NewLevel(levelMap []string, blockColor rl.Color, blockSize float32) Level {
	blocks := []rl.Rectangle{}

	for rowIndex, row := range levelMap {
		for colIndex, cell := range row {
			if cell == '1' {
				x := float32(colIndex) * blockSize
				y := float32(rowIndex) * blockSize
				block := rl.NewRectangle(x, y, blockSize, blockSize)
				blocks = append(blocks, block)
			}
		}
	}

	return Level{blocks, blockColor, blockSize}
}

func (level *Level) Draw() {
	for _, block := range level.Blocks {
		rl.DrawRectangleRec(block, level.BlockColor)
	}
}

func (level *Level) CheckCollision(axis Axis, player *Player) {
	for _, block := range level.Blocks {
		if rl.CheckCollisionRecs(block, player.Rectangle) {
			switch axis {
			case XAxis:
				// moving right
				if player.Direction.X > 0 {
					player.X = block.X - player.Width
				}

				// moving left
				if player.Direction.X < 0 {
					player.X = block.X + block.Width
				}

			case YAxis:
				// moving up
				if player.Direction.Y < 0 {
					player.Y = block.Y + block.Height
				}

				// moving down
				if player.Direction.Y > 0 {
					player.Y = block.Y - player.Height
				}
			}
		}
	}
}
