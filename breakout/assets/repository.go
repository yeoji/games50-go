package assets

import (
	"fmt"
	"games50-go/breakout/assets/graphics"
	"games50-go/internal/assets"
	"games50-go/pong/assets/fonts"
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
	}, assets.SoundLoaderConfig{})

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
}

func GetFont(name string) font.Face {
	return loadedAssets.Fonts[name]
}

func GetSprite(group string, name string) *ebiten.Image {
	return sprites[group][name]
}
