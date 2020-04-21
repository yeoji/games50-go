package main

import "github.com/hajimehoshi/ebiten"

const PLAYER_1 = 1
const PLAYER_2 = 2

type Player struct {
	PlayerNo int
	Controls Controls
	Paddle   *Paddle
	Score    int
}

type Controls struct {
	Up   ebiten.Key
	Down ebiten.Key
}

func (p *Player) reset() {
	p.Score = 0
}

func (p *Player) update() {
	if ebiten.IsKeyPressed(p.Controls.Up) {
		p.Paddle.moveUp()
	} else if ebiten.IsKeyPressed(p.Controls.Down) {
		p.Paddle.moveDown()
	}
}

func (p *Player) render(screen *ebiten.Image) {
	p.Paddle.render(screen)
}
