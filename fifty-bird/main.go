package main

import (
	"games50-go/fifty-bird/states"
	"games50-go/internal/assets"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const SCREEN_WIDTH = 512
const SCREEN_HEIGHT = 288

type Game struct {
	stateMachine *states.StateMachine
	assets       *assets.Assets
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.stateMachine.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 45, 52, 255})
	g.stateMachine.Render(screen, g.assets)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Fifty Bird")

	assets := assets.LoadAssets([]assets.FontLoaderConfig{
		{
			File: "assets/fonts/flappy.ttf",
			FontSizes: assets.FontSizeConfig{
				"mediumFont": 14,
				"flappyFont": 28,
			},
		},
	}, assets.SoundLoaderConfig{})

	if err := ebiten.RunGame(&Game{
		assets: &assets,
		stateMachine: &states.StateMachine{
			Current: &states.TitleScreenState{},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
