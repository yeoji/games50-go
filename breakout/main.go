package main

import (
	assetsRepo "games50-go/breakout/assets"
	"games50-go/breakout/assets/graphics"
	"games50-go/breakout/constants"
	"games50-go/breakout/states"
	"games50-go/internal/assets"
	statesutil "games50-go/internal/states"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

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
	return constants.ScreenWidth, constants.ScreenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

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
