package main

import (
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/audio"
	"golang.org/x/image/font"
)

const SAMPLE_RATE = 44100

type Assets struct {
	Fonts  map[string]font.Face
	Sounds map[string]*audio.Player
}

func loadAssets() Assets {
	audioContext, err := audio.NewContext(SAMPLE_RATE)
	if err != nil {
		log.Fatalf("Could not initialize audio context: %v", err)
	}

	assets := Assets{
		Fonts:  loadFonts(),
		Sounds: loadSounds(audioContext),
	}

	return assets
}

func loadFonts() map[string]font.Face {
	fonts := make(map[string]font.Face)

	fontData, err := ioutil.ReadFile("assets/fonts/font.ttf")
	if err != nil {
		log.Fatalf("Error reading font file: %v", err)
	}
	font, err := truetype.Parse(fontData)
	if err != nil {
		log.Fatalf("Error parsing font: %v", err)
	}

	fonts["smallFont"] = truetype.NewFace(font, &truetype.Options{
		Size: 8,
	})
	fonts["largeFont"] = truetype.NewFace(font, &truetype.Options{
		Size: 16,
	})
	fonts["scoreFont"] = truetype.NewFace(font, &truetype.Options{
		Size: 32,
	})

	return fonts
}

func loadSounds(audioContext *audio.Context) map[string]*audio.Player {
	sounds := make(map[string]*audio.Player)

	paddleHit, err := ioutil.ReadFile("assets/sounds/paddle_hit.wav")
	if err != nil {
		log.Fatalf("Error reading sounds asset: %v", err)
	}
	sounds["paddleHit"], _ = audio.NewPlayerFromBytes(audioContext, paddleHit)

	score, err := ioutil.ReadFile("assets/sounds/score.wav")
	if err != nil {
		log.Fatalf("Error reading sounds asset: %v", err)
	}
	sounds["score"], _ = audio.NewPlayerFromBytes(audioContext, score)

	wallHit, err := ioutil.ReadFile("assets/sounds/wall_hit.wav")
	if err != nil {
		log.Fatalf("Error reading sounds asset: %v", err)
	}
	sounds["wallHit"], _ = audio.NewPlayerFromBytes(audioContext, wallHit)

	return sounds
}
