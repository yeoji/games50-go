package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const SCREEN_WIDTH = 432
const SCREEN_HEIGHT = 243

const START_STATE = "start"
const SERVE_STATE = "serve"
const PLAY_STATE = "play"
const DONE_STATE = "done"

type Game struct {
	State         string
	Player1       *Player
	Player2       *Player
	Ball          Ball
	Assets        Assets
	ServingPlayer *Player
	WinningPlayer *Player
}

type Assets struct {
	Fonts map[string]font.Face
}

func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.State {
		case START_STATE:
			g.State = SERVE_STATE
			break
		case SERVE_STATE:
			g.State = PLAY_STATE
			g.Ball.serve(g.ServingPlayer)
			break
		case DONE_STATE:
			g.State = SERVE_STATE
			g.Ball.reset()
			g.Player1.reset()
			g.Player2.reset()
			break
		}
	}

	if g.State == PLAY_STATE {
		if g.Ball.collides(g.Player1.Paddle) || g.Ball.collides(g.Player2.Paddle) {
			g.Ball.successfullyReturned()
		}

		playerNo, scored := g.Ball.scored()
		if scored {
			if playerNo == PLAYER_1 {
				g.updatePlayerScore(g.Player1)
				g.ServingPlayer = g.Player2
			} else {
				g.updatePlayerScore(g.Player2)
				g.ServingPlayer = g.Player1
			}
		}
		g.Ball.update()
	}

	g.Player1.update()
	g.Player2.update()

	return nil
}

func (g *Game) updatePlayerScore(player *Player) {
	player.Score++

	if player.Score == 10 {
		g.State = DONE_STATE
		g.WinningPlayer = player
	} else {
		g.State = SERVE_STATE
		g.Ball.reset()
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 45, 52, 255})

	switch g.State {
	case START_STATE:
		text.Draw(screen, "Welcome to Pong!", g.Assets.Fonts["smallFont"], SCREEN_WIDTH/2-37, 10, color.White)
		text.Draw(screen, "Press Enter to begin!", g.Assets.Fonts["smallFont"], SCREEN_WIDTH/2-46, 20, color.White)
		break
	case SERVE_STATE:
		text.Draw(screen, fmt.Sprintf("Player %d's serve!", g.ServingPlayer.PlayerNo), g.Assets.Fonts["smallFont"], SCREEN_WIDTH/2-35, 10, color.White)
		text.Draw(screen, "Press Enter to serve!", g.Assets.Fonts["smallFont"], SCREEN_WIDTH/2-46, 20, color.White)
		break
	case DONE_STATE:
		text.Draw(screen, fmt.Sprintf("Player %d wins!", g.WinningPlayer.PlayerNo), g.Assets.Fonts["largeFont"], SCREEN_WIDTH/2-55, 20, color.White)
		text.Draw(screen, "Press Enter to restart!", g.Assets.Fonts["smallFont"], SCREEN_WIDTH/2-49, 30, color.White)
		break
	}

	g.displayScore(screen)

	g.Player1.render(screen)
	g.Player2.render(screen)
	g.Ball.render(screen)

	g.displayFPS(screen)
}

func (g *Game) displayScore(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%d", g.Player1.Score), g.Assets.Fonts["scoreFont"], SCREEN_WIDTH/2-50, SCREEN_HEIGHT/3, color.White)
	text.Draw(screen, fmt.Sprintf("%d", g.Player2.Score), g.Assets.Fonts["scoreFont"], SCREEN_WIDTH/2+30, SCREEN_HEIGHT/3, color.White)
}

func (g *Game) displayFPS(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("FPS: %d", int(math.Ceil(ebiten.CurrentTPS()))), g.Assets.Fonts["smallFont"], 10, 10, color.RGBA{0, 255, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func loadAssets() Assets {
	assets := Assets{
		Fonts: make(map[string]font.Face),
	}

	fontData, err := ioutil.ReadFile("assets/fonts/font.ttf")
	if err != nil {
		log.Fatalf("Error reading font file: %v", err)
	}
	font, err := truetype.Parse(fontData)
	if err != nil {
		log.Fatalf("Error parsing font: %v", err)
	}

	assets.Fonts["smallFont"] = truetype.NewFace(font, &truetype.Options{
		Size: 8,
	})
	assets.Fonts["largeFont"] = truetype.NewFace(font, &truetype.Options{
		Size: 16,
	})
	assets.Fonts["scoreFont"] = truetype.NewFace(font, &truetype.Options{
		Size: 32,
	})

	return assets
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Pong")

	assets := loadAssets()

	player1 := Player{
		PlayerNo: PLAYER_1,
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
		PlayerNo: PLAYER_2,
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
		State:         START_STATE,
		Player1:       &player1,
		Player2:       &player2,
		Ball:          ball,
		Assets:        assets,
		ServingPlayer: &player1,
	}); err != nil {
		log.Fatal(err)
	}
}
