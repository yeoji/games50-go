package main

import "github.com/hajimehoshi/ebiten"

const PLAYER_1 = 1
const PLAYER_2 = 2

// extra padding so we don't end up with an unbeatable AI
const AI_BUFFER_SPACE = 2

type Player struct {
	PlayerNo int
	Controls Controls
	Paddle   *Paddle
	Score    int
	AI       bool
}

type Controls struct {
	Up   ebiten.Key
	Down ebiten.Key
}

func (p *Player) reset() {
	p.Score = 0
}

func (p *Player) update(ball *Ball) {
	if p.AI && ball.dx > 0 {
		if ball.y <= p.Paddle.y-AI_BUFFER_SPACE {
			p.Paddle.moveUp()
		} else if ball.y > p.Paddle.y+p.Paddle.height+AI_BUFFER_SPACE {
			p.Paddle.moveDown()
		}
	} else if !p.AI {
		if ebiten.IsKeyPressed(p.Controls.Up) {
			p.Paddle.moveUp()
		} else if ebiten.IsKeyPressed(p.Controls.Down) {
			p.Paddle.moveDown()
		}
	}
}

func (p *Player) render(screen *ebiten.Image) {
	p.Paddle.render(screen)
}
