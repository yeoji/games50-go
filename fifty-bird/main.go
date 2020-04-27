package main

import (
	"games50-go/fifty-bird/states"
	"games50-go/internal/assets"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const SCREEN_WIDTH = 512
const SCREEN_HEIGHT = 288

type Game struct {
	stateMachine *states.StateMachine
	scene        *Scene
}

func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.scene.scrolling = !g.scene.scrolling
	}

	g.scene.update()
	g.stateMachine.Update(screen)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.drawBackground(screen)
	g.stateMachine.Render(screen)
	g.scene.drawGround(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Fifty Bird")

	loadedAssets := assets.LoadAssets([]assets.FontLoaderConfig{
		{
			File: "assets/fonts/flappy.ttf",
			FontSizes: assets.FontSizeConfig{
				"mediumFont": 14,
				"flappyFont": 28,
				"hugeFont":   56,
			},
		},
	}, assets.SoundLoaderConfig{
		"pause":     "assets/sounds/pause.wav",
		"jump":      "assets/sounds/jump.wav",
		"score":     "assets/sounds/score.wav",
		"explosion": "assets/sounds/explosion.wav",
		"hurt":      "assets/sounds/hurt.wav",
	})

	bgm := assets.NewLoopingAudio("assets/sounds/marios_way.mp3")
	bgm.Play()

	if err := ebiten.RunGame(&Game{
		scene: &Scene{
			Background: assets.LoadImage("assets/art/background.png"),
			Ground:     assets.LoadImage("assets/art/ground.png"),
			scrolling:  true,
		},
		stateMachine: &states.StateMachine{
			Current: &states.TitleScreenState{},
			Assets:  &loadedAssets,
		},
	}); err != nil {
		log.Fatal(err)
	}
}
