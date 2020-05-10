package assets

import (
	"games50-go/breakout/assets/fonts"
	"games50-go/breakout/assets/graphics"
	"games50-go/breakout/assets/sounds"
	"games50-go/breakout/constants"
	"games50-go/internal/assets"
	"image"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

var loadedAssets assets.Assets
var sprites map[string]map[string]*ebiten.Image

func LoadAssets() {
	loadedAssets = assets.LoadAssets([]assets.FontLoaderConfig{
		{
			FontData: fonts.Font_ttf,
			FontSizes: assets.FontSizeConfig{
				"small":  8,
				"medium": 16,
				"large":  32,
			},
		},
	}, assets.SoundLoaderConfig{
		"brick_destroyed": sounds.Brick_hit_1_wav,
		"brick_hit":       sounds.Brick_hit_2_wav,
		"confirm":         sounds.Confirm_wav,
		"high_score":      sounds.High_score_wav,
		"hurt":            sounds.Hurt_wav,
		"no_select":       sounds.No_select_wav,
		"paddle_hit":      sounds.Paddle_hit_wav,
		"pause":           sounds.Pause_wav,
		"recover":         sounds.Recover_wav,
		"select":          sounds.Select_wav,
		"victory":         sounds.Victory_wav,
		"wall_hit":        sounds.Wall_hit_wav,
	})

	loadSprites()
}

func loadSprites() {
	sprites = make(map[string]map[string]*ebiten.Image)

	arrowsSpriteSheet := assets.LoadImage(graphics.Arrows_png)
	arrowSprites := assets.CutSpritesheet(arrowsSpriteSheet, assets.CutSpritesheetConfig{
		XEnd: 48,
		YEnd: 24,
		Sprites: []assets.SpriteConfig{{
			Width:  24,
			Height: 24,
		}},
	})
	assets.GroupSprites(arrowSprites, sprites, []string{"arrows"}, []string{"left", "right"})

	breakoutSpriteSheet := assets.LoadImage(graphics.Breakout_png)
	paddleColours := []string{"paddles-blue", "paddles-green", "paddles-red", "paddles-purple"}
	paddleSprites := assets.CutSpritesheet(breakoutSpriteSheet, assets.CutSpritesheetConfig{
		YOffset: 64,
		XLimit:  128,
		XEnd:    128,
		YEnd:    192,
		Sprites: []assets.SpriteConfig{{
			Width:  32,
			Height: constants.PaddleHeight,
		}, {
			Width:  64,
			Height: constants.PaddleHeight,
		}, {
			Width:  96,
			Height: constants.PaddleHeight,
		}, {
			Width:  128,
			Height: constants.PaddleHeight,
		}},
	})
	assets.GroupSprites(paddleSprites, sprites, paddleColours, []string{"smallest", "small", "large", "largest"})

	ballColours := []string{"blue", "green", "red", "purple", "yellow", "grey", "gold"}
	ballSprites := assets.CutSpritesheet(breakoutSpriteSheet, assets.CutSpritesheetConfig{
		XOffset: 96,
		YOffset: 48,
		XLimit:  128,
		XEnd:    120,
		YEnd:    64,
		Sprites: []assets.SpriteConfig{{
			Width:  constants.BallWidth,
			Height: constants.BallHeight,
		}},
	})
	assets.GroupSprites(ballSprites, sprites, []string{"balls"}, ballColours)

	brickTiers := []string{"basic", "extra", "super", "ultra"}
	brickColours := []string{"bricks-blue", "bricks-green", "bricks-red", "bricks-purple", "bricks-yellow"}
	brickSprites := assets.CutSpritesheet(breakoutSpriteSheet, assets.CutSpritesheetConfig{
		XEnd: 64,
		YEnd: 64,
		Sprites: []assets.SpriteConfig{{
			Width:  constants.BrickWidth,
			Height: constants.BrickHeight,
		}},
	})
	assets.GroupSprites(brickSprites, sprites, brickColours, brickTiers)

	sprites["bricks"] = make(map[string]*ebiten.Image)
	sprites["bricks"]["locked"] = breakoutSpriteSheet.SubImage(image.Rect(160, 48, 160+constants.BrickWidth, 48+constants.BrickHeight)).(*ebiten.Image)

	heartSprites := assets.CutSpritesheet(breakoutSpriteSheet, assets.CutSpritesheetConfig{
		XOffset: 128,
		YOffset: 48,
		XEnd:    148,
		YEnd:    58,
		Sprites: []assets.SpriteConfig{{
			Width:  constants.HeartWidth,
			Height: constants.HeartHeight,
		}},
	})
	assets.GroupSprites(heartSprites, sprites, []string{"hearts"}, []string{"full", "empty"})

	powerupSprites := assets.CutSpritesheet(breakoutSpriteSheet, assets.CutSpritesheetConfig{
		XOffset: 128,
		YOffset: 192,
		XEnd:    160,
		YEnd:    208,
		Sprites: []assets.SpriteConfig{{
			Width:  constants.PowerupWidth,
			Height: constants.PowerupHeight,
		}},
	})
	assets.GroupSprites(powerupSprites, sprites, []string{"powerups"}, []string{"extra-balls", "key"})

	sprites["particles"] = make(map[string]*ebiten.Image)
	sprites["particles"]["brick-explode"] = assets.LoadImage(graphics.Particle_png)
}

func GetFont(name string) font.Face {
	return loadedAssets.Fonts[name]
}

func PlaySound(name string) {
	sound := loadedAssets.Sounds[name]
	sound.Play()
	sound.Rewind()
}

func GetSprite(group string, name string) *ebiten.Image {
	return sprites[group][name]
}
