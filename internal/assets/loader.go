package assets

import (
	"image"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"golang.org/x/image/font"
)

const SAMPLE_RATE = 44100

type Assets struct {
	Fonts  map[string]font.Face
	Sounds map[string]*audio.Player
	Images map[string]*ebiten.Image
}

type FontSizeConfig map[string]float64
type FontLoaderConfig struct {
	File      string
	FontSizes FontSizeConfig
}

// map of sound name and sound file
type SoundLoaderConfig map[string]string

type ImageLoaderConfig map[string]string

func LoadAssets(fonts []FontLoaderConfig, sounds SoundLoaderConfig, images ImageLoaderConfig) Assets {
	audioContext, err := audio.NewContext(SAMPLE_RATE)
	if err != nil {
		log.Fatalf("Could not initialize audio context: %v", err)
	}

	assets := Assets{
		Fonts:  loadFonts(fonts),
		Sounds: loadSounds(audioContext, sounds),
		Images: loadImages(images),
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

func loadSounds(audioContext *audio.Context, soundsToLoad SoundLoaderConfig) map[string]*audio.Player {
	sounds := make(map[string]*audio.Player)

	for name, file := range soundsToLoad {
		sound, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading sounds asset: %v", err)
		}
		sounds[name], _ = audio.NewPlayerFromBytes(audioContext, sound)
	}

	return sounds
}

func loadImages(imagesToLoad ImageLoaderConfig) map[string]*ebiten.Image {
	images := make(map[string]*ebiten.Image)

	for name, file := range imagesToLoad {
		file, err := os.Open(file)
		defer file.Close()
		if err != nil {
			log.Fatalf("Error loading image asset: %v", err)
		}

		decoded, _, err := image.Decode(file)
		if err != nil {
			log.Fatalf("Error loading image asset: %v", err)
		}

		images[name], _ = ebiten.NewImageFromImage(decoded, ebiten.FilterDefault)
	}

	return images
}
