package assets

import (
	"fmt"
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

// TODO should manage this better
func loadSprites() {
	sprites = make(map[string]map[string]*ebiten.Image)

	arrowsSpriteSheet := assets.LoadImage(graphics.Arrows_png)
	sprites["arrows"] = make(map[string]*ebiten.Image)
	sprites["arrows"]["left"] = arrowsSpriteSheet.SubImage(image.Rect(0, 0, 24, 24)).(*ebiten.Image)
	sprites["arrows"]["right"] = arrowsSpriteSheet.SubImage(image.Rect(24, 0, 48, 24)).(*ebiten.Image)

	breakoutSpriteSheet := assets.LoadImage(graphics.Breakout_png)
	paddleColours := []string{"blue", "green", "red", "purple"}
	yOffset := 64
	for _, colour := range paddleColours {
		// paddle height = 16 // width 32, 64, 96, 128
		spriteGroup := fmt.Sprintf("paddles-%s", colour)
		sprites[spriteGroup] = make(map[string]*ebiten.Image)

		sprites[spriteGroup]["smallest"] = breakoutSpriteSheet.SubImage(image.Rect(0, yOffset, 32, yOffset+16)).(*ebiten.Image)
		sprites[spriteGroup]["small"] = breakoutSpriteSheet.SubImage(image.Rect(32, yOffset, 96, yOffset+16)).(*ebiten.Image)
		sprites[spriteGroup]["large"] = breakoutSpriteSheet.SubImage(image.Rect(96, yOffset, 192, yOffset+16)).(*ebiten.Image)
		yOffset += 16
		sprites[spriteGroup]["largest"] = breakoutSpriteSheet.SubImage(image.Rect(0, yOffset, 128, yOffset+16)).(*ebiten.Image)
		yOffset += 16
	}

	ballColours := []string{"blue", "green", "red", "purple", "yellow", "grey", "gold"}
	// x-offset 96 y-offset 48 x-finish 128
	sprites["balls"] = make(map[string]*ebiten.Image)
	xStart := 96
	xOffset := 96
	xLimit := 128
	yOffset = 48
	for _, colour := range ballColours {
		sprites["balls"][colour] = breakoutSpriteSheet.SubImage(image.Rect(xOffset, yOffset, xOffset+constants.BallWidth, yOffset+constants.BallHeight)).(*ebiten.Image)
		xOffset += constants.BallWidth
		if xOffset == xLimit {
			yOffset += constants.BallHeight
			xOffset = xStart
		}
	}

	brickTiers := []string{"basic", "extra", "super", "ultra"}
	brickColours := []string{"blue", "green", "red", "purple", "yellow"}
	xStart = 0
	xOffset = 0
	xLimit = 192
	yOffset = 0
	for _, colour := range brickColours {
		spriteGroup := fmt.Sprintf("bricks-%s", colour)
		sprites[spriteGroup] = make(map[string]*ebiten.Image)

		for _, tier := range brickTiers {
			sprites[spriteGroup][tier] = breakoutSpriteSheet.SubImage(image.Rect(xOffset, yOffset, xOffset+constants.BrickWidth, yOffset+constants.BrickHeight)).(*ebiten.Image)
			xOffset += constants.BrickWidth
			if xOffset == xLimit {
				yOffset += constants.BrickHeight
				xOffset = xStart
			}
		}
	}
	sprites["bricks"] = make(map[string]*ebiten.Image)
	sprites["bricks"]["locked"] = breakoutSpriteSheet.SubImage(image.Rect(160, 48, 160+constants.BrickWidth, 48+constants.BrickHeight)).(*ebiten.Image)

	sprites["hearts"] = make(map[string]*ebiten.Image)
	sprites["hearts"]["full"] = breakoutSpriteSheet.SubImage(image.Rect(128, 48, 128+constants.HeartWidth, 48+constants.HeartHeight)).(*ebiten.Image)
	sprites["hearts"]["empty"] = breakoutSpriteSheet.SubImage(image.Rect(138, 48, 138+constants.HeartWidth, 48+constants.HeartHeight)).(*ebiten.Image)

	sprites["powerups"] = make(map[string]*ebiten.Image)
	sprites["powerups"]["extra-balls"] = breakoutSpriteSheet.SubImage(image.Rect(128, 192, 128+constants.PowerupWidth, 192+constants.PowerupHeight)).(*ebiten.Image)
	sprites["powerups"]["key"] = breakoutSpriteSheet.SubImage(image.Rect(144, 192, 144+constants.PowerupWidth, 192+constants.PowerupHeight)).(*ebiten.Image)
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
