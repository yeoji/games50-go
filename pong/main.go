package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
)

const SCREEN_WIDTH = 432
const SCREEN_HEIGHT = 243

const START_STATE = "start"
const SERVE_STATE = "serve"
const PLAY_STATE = "play"
const DONE_STATE = "done"

var assets Assets
var numPlayers = 1

type Game struct {
	State         string
	Player1       *Player
	Player2       *Player
	Ball          Ball
	ServingPlayer *Player
	WinningPlayer *Player
}

func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.State = START_STATE
		g.Ball.reset()
		g.Player1.reset()
		g.Player2.reset()
		g.ServingPlayer = g.Player1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.State {
		case START_STATE:
			g.State = SERVE_STATE
			if numPlayers == 1 {
				g.Player2.AI = true
			} else {
				g.Player2.AI = false
			}
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

	if g.State == START_STATE {
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && numPlayers == 2 {
			numPlayers = 1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) && numPlayers == 1 {
			numPlayers = 2
		}
	}

	if g.State == PLAY_STATE {
		if g.Ball.collides(g.Player1.Paddle) {
			assets.Sounds["paddleHit"].Play()
			g.Ball.successfullyReturned(g.Player1.Paddle)
			assets.Sounds["paddleHit"].Rewind()
		}
		if g.Ball.collides(g.Player2.Paddle) {
			assets.Sounds["paddleHit"].Play()
			g.Ball.successfullyReturned(g.Player2.Paddle)
			assets.Sounds["paddleHit"].Rewind()
		}

		playerNo, scored := g.Ball.scored()
		if scored {
			assets.Sounds["score"].Play()
			if playerNo == PLAYER_1 {
				g.updatePlayerScore(g.Player1)
				g.ServingPlayer = g.Player2
			} else {
				g.updatePlayerScore(g.Player2)
				g.ServingPlayer = g.Player1
			}
			assets.Sounds["score"].Rewind()
		}
		g.Ball.update()
	}

	g.Player1.update(&g.Ball)
	g.Player2.update(&g.Ball)

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
		text.Draw(screen, "Welcome to Pong!", assets.Fonts["largeFont"], SCREEN_WIDTH/2-75, 30, color.White)

		text.Draw(screen, "Players", assets.Fonts["largeFont"], SCREEN_WIDTH/2-34, SCREEN_HEIGHT/2, color.White)
		text.Draw(screen, fmt.Sprintf("< %d >", numPlayers), assets.Fonts["largeFont"], SCREEN_WIDTH/2-23, SCREEN_HEIGHT/2+20, color.White)

		text.Draw(screen, "Select number of players and press Enter!", assets.Fonts["smallFont"], SCREEN_WIDTH/2-93, SCREEN_HEIGHT/2+50, color.White)
		break
	case SERVE_STATE:
		text.Draw(screen, fmt.Sprintf("Player %d's serve!", g.ServingPlayer.PlayerNo), assets.Fonts["smallFont"], SCREEN_WIDTH/2-35, 10, color.White)
		text.Draw(screen, "Press Enter to serve!", assets.Fonts["smallFont"], SCREEN_WIDTH/2-46, 20, color.White)
		break
	case DONE_STATE:
		text.Draw(screen, fmt.Sprintf("Player %d wins!", g.WinningPlayer.PlayerNo), assets.Fonts["largeFont"], SCREEN_WIDTH/2-55, 20, color.White)
		text.Draw(screen, "Press Enter to restart!", assets.Fonts["smallFont"], SCREEN_WIDTH/2-49, 30, color.White)
		break
	}

	if g.State != START_STATE {
		g.displayScore(screen)

		g.Player1.render(screen)
		g.Player2.render(screen)
		g.Ball.render(screen)

		g.displayFPS(screen)
	}
}

func (g *Game) displayScore(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%d", g.Player1.Score), assets.Fonts["scoreFont"], SCREEN_WIDTH/2-50, SCREEN_HEIGHT/3, color.White)
	text.Draw(screen, fmt.Sprintf("%d", g.Player2.Score), assets.Fonts["scoreFont"], SCREEN_WIDTH/2+30, SCREEN_HEIGHT/3, color.White)
}

func (g *Game) displayFPS(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("FPS: %d", int(math.Ceil(ebiten.CurrentTPS()))), assets.Fonts["smallFont"], 10, 10, color.RGBA{0, 255, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Pong")

	assets = loadAssets()

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
		ServingPlayer: &player1,
	}); err != nil {
		log.Fatal(err)
	}
}
