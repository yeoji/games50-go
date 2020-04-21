package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Paddle struct {
	width  float64
	height float64
	x      float64
	y      float64
}

const PADDLE_SPEED = 200

func (p *Paddle) moveUp() {
	p.y -= (PADDLE_SPEED * 1 / ebiten.CurrentTPS())
	if p.y < 0 {
		p.y = 0
	}
}

func (p *Paddle) moveDown() {
	p.y += (PADDLE_SPEED * 1 / ebiten.CurrentTPS())
	if p.y > SCREEN_HEIGHT-p.height {
		p.y = SCREEN_HEIGHT - p.height
	}
}

func (p *Paddle) render(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.x, p.y, p.width, p.height, color.White)
}
