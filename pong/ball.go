package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Ball struct {
	width  float64
	height float64
	x      float64
	y      float64
	dx     float64 // horizontal velocity
	dy     float64 // vertical velocity
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (b *Ball) serve(servingPlayer *Player) {
	b.dy = float64(randomNumInRange(-50, 50))

	if servingPlayer.PlayerNo == PLAYER_1 {
		b.dx = float64(randomNumInRange(140, 200))
	} else {
		b.dx = -float64(randomNumInRange(140, 200))
	}
}

func (b *Ball) successfullyReturned() {
	// keep vertical direction of ball the same
	if b.dy < 0 {
		b.dy = -float64(randomNumInRange(10, 150))
	} else {
		b.dy = float64(randomNumInRange(10, 150))
	}

	// increase the speed of the ball
	b.dx = -b.dx * 1.03
}

func (b *Ball) collides(paddle *Paddle) bool {
	if b.x > paddle.x+paddle.width || paddle.x > b.x+b.width {
		return false
	}

	if b.y > paddle.y+paddle.height || paddle.y > b.y+b.height {
		return false
	}

	return true
}

func (b *Ball) scored() (int, bool) {
	if b.x > SCREEN_WIDTH {
		return PLAYER_1, true
	}
	if b.x+b.width < 0 {
		return PLAYER_2, true
	}
	return 0, false
}

func (b *Ball) update() {
	b.x += b.dx * 1 / ebiten.CurrentTPS()
	b.y += b.dy * 1 / ebiten.CurrentTPS()

	if b.y < 0 {
		b.y = 0
		b.dy = -b.dy
	} else if b.y > SCREEN_HEIGHT-b.height {
		b.y = SCREEN_HEIGHT - b.height
		b.dy = -b.dy
	}
}

func (b *Ball) render(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.x, b.y, b.width, b.height, color.White)
}

func randomNumInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
