package main

import (
	"games50-go/fifty-bird/assets/art"
	"games50-go/fifty-bird/assets/fonts"
	"games50-go/fifty-bird/assets/sounds"
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
			FontData: fonts.Flappy_ttf,
			FontSizes: assets.FontSizeConfig{
				"mediumFont": 14,
				"flappyFont": 28,
				"hugeFont":   56,
			},
		},
	}, assets.SoundLoaderConfig{
		"pause":     sounds.Pause_wav,
		"jump":      sounds.Jump_wav,
		"score":     sounds.Score_wav,
		"explosion": sounds.Explosion_wav,
		"hurt":      sounds.Hurt_wav,
	})

	bgm := assets.NewLoopingAudio(sounds.MariosWay_mp3)
	bgm.Play()

	if err := ebiten.RunGame(&Game{
		scene: &Scene{
			Background: assets.LoadImage(art.Background_png),
			Ground:     assets.LoadImage(art.Ground_png),
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
