package main

import (
	"games50-go/fifty-bird/states"
	"games50-go/internal/assets"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const SCREEN_WIDTH = 512
const SCREEN_HEIGHT = 288

type Game struct {
	stateMachine *states.StateMachine
	scene        *Scene
	assets       *assets.Assets
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.scene.update()
	g.stateMachine.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.render(screen)
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
				"hugeFont":   56,
			},
		},
	}, assets.SoundLoaderConfig{}, assets.ImageLoaderConfig{
		"background": "assets/art/background.png",
		"ground":     "assets/art/ground.png",
	})

	if err := ebiten.RunGame(&Game{
		assets: &assets,
		scene: &Scene{
			Background: assets.Images["background"],
			Ground:     assets.Images["ground"],
			scrolling:  true,
		},
		stateMachine: &states.StateMachine{
			Current: &states.TitleScreenState{},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
