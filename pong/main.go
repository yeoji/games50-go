package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const SCREEN_WIDTH = 432
const SCREEN_HEIGHT = 243

type Game struct {
	Player1 Player
	Player2 Player
	Ball    Ball
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.Player1.update()
	g.Player2.update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 45, 52, 255})

	g.Player1.render(screen)
	g.Player2.render(screen)
	g.Ball.render(screen)

	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Pong")

	player1 := Player{
		Controls: Controls{
			Up:   ebiten.KeyW,
			Down: ebiten.KeyS,
		},
		Paddle: &Paddle{
			x:      10,
			y:      30,
			width:  5,
			height: 20,
		},
	}

	player2 := Player{
		Controls: Controls{
			Up:   ebiten.KeyUp,
			Down: ebiten.KeyDown,
		},
		Paddle: &Paddle{
			x:      SCREEN_WIDTH - 15,
			y:      SCREEN_HEIGHT - 30,
			width:  5,
			height: 20,
		},
	}

	ball := Ball{
		x:      (SCREEN_WIDTH / 2) - 2,
		y:      (SCREEN_HEIGHT / 2) - 2,
		width:  4,
		height: 4,
	}

	if err := ebiten.RunGame(&Game{
		Player1: player1,
		Player2: player2,
		Ball:    ball,
	}); err != nil {
		log.Fatal(err)
	}
}
