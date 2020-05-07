package assets

import (
	"games50-go/internal/assets"
	"games50-go/pong/assets/fonts"

	"golang.org/x/image/font"
)

var loadedAssets assets.Assets

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
}

func GetFont(name string) font.Face {
	return loadedAssets.Fonts[name]
}
