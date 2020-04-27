package assets

import (
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/audio"
	"golang.org/x/image/font"
)

type Assets struct {
	Fonts  map[string]font.Face
	Sounds map[string]*audio.Player
}

type FontSizeConfig map[string]float64
type FontLoaderConfig struct {
	File      string
	FontSizes FontSizeConfig
}

func LoadAssets(fonts []FontLoaderConfig, sounds SoundLoaderConfig) Assets {
	audioContext := getAudioContext()

	assets := Assets{
		Fonts:  loadFonts(fonts),
		Sounds: loadSounds(audioContext, sounds),
	}

	return assets
}

// only loads TTF fonts
func loadFonts(fontsToLoad []FontLoaderConfig) map[string]font.Face {
	fonts := make(map[string]font.Face)

	for _, fontToLoad := range fontsToLoad {
		fontData, err := ioutil.ReadFile(fontToLoad.File)
		if err != nil {
			log.Fatalf("Error reading font file: %v", err)
		}
		font, err := truetype.Parse(fontData)
		if err != nil {
			log.Fatalf("Error parsing font: %v", err)
		}

		for name, size := range fontToLoad.FontSizes {
			fonts[name] = truetype.NewFace(font, &truetype.Options{
				Size: size,
			})
		}
	}

	return fonts
}
