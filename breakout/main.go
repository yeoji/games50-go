package main

import (
	assetsRepo "games50-go/breakout/assets"
	"games50-go/breakout/assets/graphics"
	"games50-go/breakout/states"
	"games50-go/internal/assets"
	statesutil "games50-go/internal/states"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const SCREEN_WIDTH = 432
const SCREEN_HEIGHT = 243

type Game struct {
	World *World
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.World.update(screen)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.World.render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Breakout")

	assetsRepo.LoadAssets()

	if err := ebiten.RunGame(&Game{
		World: &World{
			Background: assets.LoadImage(graphics.Background_png),
			StateMachine: &statesutil.StateMachine{
				Current: &states.MenuState{},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
