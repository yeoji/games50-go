package assets

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type SpriteConfig struct {
	Width  int
	Height int
}

type CutSpritesheetConfig struct {
	XOffset int
	YOffset int
	XLimit  int
	XEnd    int
	YEnd    int
	Sprites []SpriteConfig
}

func LoadImage(imageData []byte) *ebiten.Image {
	decoded, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		log.Fatalf("Error loading image asset: %v", err)
	}

	loadedImage, _ := ebiten.NewImageFromImage(decoded, ebiten.FilterDefault)
	return loadedImage
}

func CutSpritesheet(sheet *ebiten.Image, cutConfig CutSpritesheetConfig) []*ebiten.Image {
	sheetWidth, _ := sheet.Size()

	var cutSprites []*ebiten.Image

	xOffset := cutConfig.XOffset
	yOffset := cutConfig.YOffset
	spriteConfigIndex := 0

	for {
		sprite := cutConfig.Sprites[spriteConfigIndex]

		cut := sheet.SubImage(image.Rect(xOffset, yOffset, xOffset+sprite.Width, yOffset+sprite.Height)).(*ebiten.Image)
		cutSprites = append(cutSprites, cut)

		xOffset += sprite.Width
		if xOffset == cutConfig.XEnd && yOffset+sprite.Height == cutConfig.YEnd {
			break
		}

		if xOffset == sheetWidth || xOffset == cutConfig.XLimit {
			// break to new row in sprite sheet
			yOffset += sprite.Height
			xOffset = cutConfig.XOffset
		}

		// loop sprite configs
		spriteConfigIndex = (spriteConfigIndex + 1) % len(cutConfig.Sprites)
	}

	return cutSprites
}

// Groups the sprite array into different groups and names them
// Updates the groupedSprites map in place
func GroupSprites(sprites []*ebiten.Image, groupedSprites map[string]map[string]*ebiten.Image, groups []string, names []string) {
	spriteIndex := 0
	for _, group := range groups {
		groupedSprites[group] = make(map[string]*ebiten.Image)

		for _, name := range names {
			groupedSprites[group][name] = sprites[spriteIndex]

			spriteIndex++
		}
	}
}
