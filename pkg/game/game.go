package game

import (
	"fmt"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	a "github.com/thuyencode/raylib-go-introduction/pkg/animation"
	c "github.com/thuyencode/raylib-go-introduction/pkg/configs"
	e "github.com/thuyencode/raylib-go-introduction/pkg/entity"
	t "github.com/thuyencode/raylib-go-introduction/pkg/timer"
	u "github.com/thuyencode/raylib-go-introduction/pkg/util"
)

type Game struct {
	ShouldDisplayDebugInfo bool
	Player                 e.Player
	Timer                  t.Timer
	staticTextures         map[string]rl.Texture2D
	animationTextures      map[string][]rl.Texture2D
	font                   rl.Font
	sounds                 map[string]rl.Sound
	music                  rl.Music
	stars                  []*e.Star
	lasers                 []*e.Laser
	meteors                []*e.Meteor
	explosions             []*a.ExplosionAnimation
	ended                  bool
	scores                 int
}

func loadStaticTexture() map[string]rl.Texture2D {
	textures := make(map[string]rl.Texture2D)
	textures["player"] = rl.LoadTexture(path.Join("assets", "spaceship.png"))
	textures["star"] = rl.LoadTexture(path.Join("assets", "star.png"))
	textures["laser"] = rl.LoadTexture(path.Join("assets", "laser.png"))
	textures["meteor"] = rl.LoadTexture(path.Join("assets", "meteor.png"))

	return textures
}

func loadAnimationTextures() map[string][]rl.Texture2D {
	textures := make(map[string][]rl.Texture2D)

	for index := range 28 {
		textures["explosion"] = append(
			textures["explosion"],
			rl.LoadTexture(path.Join(
				"assets",
				"animation",
				"explosion",
				fmt.Sprintf("%d.png", index+1),
			)))
	}

	return textures
}

func loadFont() rl.Font {
	return rl.LoadFontEx(path.Join("assets", "Zero Hour.otf"), c.FONT_SIZE, nil, 0)
}

func loadSounds() map[string]rl.Sound {
	sounds := make(map[string]rl.Sound)
	sounds["explosion"] = rl.LoadSound(path.Join("assets", "audio", "explosion.wav"))
	sounds["laser"] = rl.LoadSound(path.Join("assets", "audio", "laser.wav"))

	return sounds
}

func loadMusic() rl.Music {
	return rl.LoadMusicStream(path.Join("assets", "audio", "music.wav"))
}

func generateStars() []*e.Star {
	var stars []*e.Star
	for range 30 {
		stars = append(stars, e.NewStar(
			rl.NewVector2(
				u.RandFloat32(0, float32(c.WINDOW_WIDTH)),
				u.RandFloat32(0, float32(c.WINDOW_HEIGHT))),
			u.RandFloat32(0.5, 1.2),
		))
	}

	return stars
}

func NewGame(shouldDisplayFPS bool, logLevel rl.TraceLogLevel) *Game {
	rl.SetTraceLogLevel(logLevel)
	rl.InitWindow(c.WINDOW_WIDTH, c.WINDOW_HEIGHT, "2D Game")
	rl.InitAudioDevice()

	return &Game{
		ShouldDisplayDebugInfo: shouldDisplayFPS,
		staticTextures:         loadStaticTexture(),
		animationTextures:      loadAnimationTextures(),
		font:                   loadFont(),
		sounds:                 loadSounds(),
		music:                  loadMusic(),
		stars:                  generateStars(),
	}
}

func (game *Game) SetPlayer(pos rl.Vector2, shootLaserCallback func(pos rl.Vector2)) {
	playerTexture, exist := game.staticTextures["player"]

	if !exist {
		panic("\"player\" asset not found!")
	}

	game.Player = *e.NewPlayer(playerTexture, pos, shootLaserCallback)

}

func (game *Game) SetTimer(duration float64, repeat bool, callback func()) {
	game.Timer = *t.NewTimer(duration, repeat, callback)
}

func (game *Game) DrawStars() {
	starTexture, exist := game.staticTextures["star"]

	if !exist {
		panic("\"star\" asset not found")
	}

	for _, star := range game.stars {
		rl.DrawTextureEx(starTexture, star.Pos, 0, star.Scale, rl.White)
	}
}

func (game *Game) CreateLaser(pos rl.Vector2) {
	laserTexture, laserTextureExist := game.staticTextures["laser"]
	if !laserTextureExist {
		panic("\"laser.png\" asset not found")
	}

	laserSound, laserSoundExist := game.sounds["laser"]
	if !laserSoundExist {
		panic("\"laser.wav\" asset not found")
	}

	adjustedPos := rl.NewVector2(pos.X-float32(laserTexture.Width)/2, pos.Y)
	game.lasers = append(game.lasers, e.NewLaser(laserTexture, adjustedPos))

	rl.PlaySound(laserSound)
}

func (game *Game) CreateMeteor() {
	meteorTexture, exist := game.staticTextures["meteor"]

	if !exist {
		panic("\"meteor\" asset not found")
	}

	game.meteors = append(game.meteors, e.NewMeteor(meteorTexture))
}

func (game *Game) DiscardSprites() {
	var nonDiscardedLasers []*e.Laser
	for _, laser := range game.lasers {
		if !laser.ShouldBeDiscarded {
			nonDiscardedLasers = append(nonDiscardedLasers, laser)
		}
	}

	var nonDiscardedMeteors []*e.Meteor
	for _, meteor := range game.meteors {
		if !meteor.ShouldBeDiscarded {
			nonDiscardedMeteors = append(nonDiscardedMeteors, meteor)
		}
	}

	var nonDiscardedExplosions []*a.ExplosionAnimation
	for _, explosion := range game.explosions {
		if !explosion.ShouldBeDiscarded {
			nonDiscardedExplosions = append(nonDiscardedExplosions, explosion)
		}
	}

	game.lasers = nonDiscardedLasers
	game.meteors = nonDiscardedMeteors
	game.explosions = nonDiscardedExplosions
}

func (game *Game) CheckCollisions() {
	explosionTextures, exist := game.animationTextures["explosion"]

	if !exist {
		panic("\"explosion\" asset not found")
	}

	for _, laser := range game.lasers {
		for _, meteor := range game.meteors {
			if rl.CheckCollisionCircleRec(
				meteor.GetCenterPos(), meteor.CollisionRadius,
				laser.GetRect(),
			) {
				laser.ShouldBeDiscarded, meteor.ShouldBeDiscarded = true, true

				animationPos := rl.NewVector2(
					laser.Pos.X-float32(laser.Texture.Width/2),
					laser.Pos.Y)
				game.explosions = append(game.explosions,
					a.NewExplosionAnimation(explosionTextures, animationPos))

				game.scores++

				rl.PlaySound(game.sounds["explosion"])
			}
		}
	}

	for _, meteor := range game.meteors {
		if rl.CheckCollisionCircles(
			game.Player.GetCenterPos(), game.Player.CollisionRadius,
			meteor.GetCenterPos(), meteor.CollisionRadius,
		) {
			game.ended = true
		}
	}
}

func (game *Game) Restart() {
	game.ended = false
	game.scores = 0
	game.meteors = nil
	game.lasers = nil
	game.explosions = nil
	game.Player.MoveToCenterOfWindow()
	game.Timer.Activate()
	rl.PlayMusicStream(game.music)
}

func (game *Game) Input() {
	if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) {
		if game.ended {
			game.Restart()
		}
	}
}

func (game *Game) Update() {
	rl.UpdateMusicStream(game.music)

	if game.ended {
		rl.PauseMusicStream(game.music)
		game.Timer.Deactivate()
		return
	}

	dt := rl.GetFrameTime()

	game.Player.Update(dt)
	game.DiscardSprites()
	game.Timer.Update()

	for _, laser := range game.lasers {
		laser.Update(dt)
	}
	for _, meteor := range game.meteors {
		meteor.Update(dt)
	}
	for _, explosion := range game.explosions {
		explosion.Update(dt)
	}

	game.CheckCollisions()
}

func (game *Game) DrawScore() {
	score := fmt.Sprintf("%d", game.scores)
	textSize := rl.MeasureTextEx(game.font, score, c.FONT_SIZE, 0)
	centeredPos := rl.NewVector2(c.WINDOW_WIDTH/2-textSize.X/2, 50)

	rl.DrawTextEx(game.font, score, centeredPos, c.FONT_SIZE, 0, rl.White)
}

func (game *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	if game.ShouldDisplayDebugInfo {
		rl.DrawFPS(10, 10)
		rl.DrawText(
			fmt.Sprintf("%d LASERS", len(game.lasers)),
			10, 35, 20, rl.DarkGreen)
		rl.DrawText(
			fmt.Sprintf("%d METEORS", len(game.meteors)),
			10, 60, 20, rl.DarkGreen)
	}

	game.DrawStars()
	game.DrawScore()
	game.Player.Draw()

	for _, laser := range game.lasers {
		laser.Draw()
	}
	for _, meteor := range game.meteors {
		meteor.Draw()
	}
	for _, explosion := range game.explosions {
		explosion.Draw()
	}

	rl.EndDrawing()

}

func (game *Game) Run() {
	rl.PlayMusicStream(game.music)

	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()
	defer rl.UnloadMusicStream(game.music)

	game.Player.MoveToCenterOfWindow()
	game.Timer.Activate()

	for !rl.WindowShouldClose() {
		targetFPS := int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
		rl.SetTargetFPS(targetFPS)

		game.Input()
		game.Update()
		game.Draw()
	}
}
