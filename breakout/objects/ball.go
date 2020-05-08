package objects

import (
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/breakout/objects/paddles"
	"games50-go/internal/utils"

	"github.com/hajimehoshi/ebiten"
)

type colour int

const (
	Blue colour = iota
	Green
	Red
	Purple
	Yellow
	Grey
	Gold
)

func (c colour) string() string {
	return []string{"blue", "green", "red", "purple", "yellow", "grey", "gold"}[c]
}

type Ball struct {
	x      float64
	y      float64
	colour colour
}

func NewBall() Ball {
	return Ball{
		colour: colour(utils.RandomNumInRange(0, 6)),
	}
}

func (b *Ball) FollowPaddle(paddle *paddles.Paddle) {
	paddleX, paddleY := paddle.Position()

	b.x = paddleX + float64(paddle.Width()/2) - constants.BallWidth/2
	b.y = paddleY - constants.BallHeight
}

func (b *Ball) Render(screen *ebiten.Image) {
	ballOptions := &ebiten.DrawImageOptions{}
	ballOptions.GeoM.Translate(b.x, b.y)

	screen.DrawImage(assets.GetSprite("balls", b.colour.string()), ballOptions)
}
