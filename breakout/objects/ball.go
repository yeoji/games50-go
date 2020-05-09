package objects

import (
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/breakout/objects/paddles"
	"games50-go/internal/utils"
	"image"
	"math"

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
	dx     float64
	dy     float64
	colour colour
}

func NewBall() *Ball {
	return &Ball{
		colour: colour(utils.RandomNumInRange(0, 6)),
	}
}

func (b *Ball) Move() {
	b.x += b.dx * (1 / ebiten.CurrentTPS())
	if b.x < 0 {
		b.x = 0
		b.dx = -b.dx
	}
	if b.x+constants.BallWidth > constants.ScreenWidth {
		b.x = constants.ScreenWidth - constants.BallWidth
		b.dx = -b.dx
	}

	b.y += b.dy * (1 / ebiten.CurrentTPS())
	if b.y < 0 {
		b.y = 0
		b.dy = -b.dy
	}
}

func (b *Ball) IsOutOfScreen() bool {
	return b.y > constants.ScreenHeight
}

func (b *Ball) Serve() {
	b.dx = float64(utils.RandomNumInRange(-200, 200))
	b.dy = float64(utils.RandomNumInRange(-60, -50))
}

func (b *Ball) HitBrick(brick *Brick, keyPowerup bool) int {
	if !brick.InPlay {
		return 0
	}

	brickScore := 0
	if !brick.Locked || keyPowerup {
		brickScore = brick.Score()
		brick.Hit()
	}

	// offset the check by a couple of pixels so that flush corner hits register as Y flips, not X flips
	leftCornerOffset := 6
	rightCornerOffset := 2

	if b.x+float64(leftCornerOffset) < brick.x && b.dx > 0 {
		// left edge hit
		b.dx = -b.dx
		b.x = brick.x - constants.BallWidth
	} else if b.x+float64(rightCornerOffset) > brick.x+constants.BrickWidth && b.dx < 0 {
		// right edge hit
		b.dx = -b.dx
		b.x = brick.x + constants.BrickWidth
	} else if b.y < brick.y {
		// top edge hit
		b.dy = -b.dy
		b.y = brick.y - constants.BallHeight
	} else {
		// bottom edge hit
		b.dy = -b.dy
		b.y = brick.y + constants.BrickHeight
	}

	if math.Abs(b.dy) < 150 {
		b.dy = b.dy * 1.02
	}

	return brickScore
}

func (b *Ball) BounceOffPaddle(paddle *paddles.Paddle) {
	paddleX, paddleY := paddle.Position()
	b.y = paddleY - constants.BallHeight
	b.dy = -b.dy

	// if the ball hits the paddle on the left side while the paddle is moving left
	if b.x < paddleX+float64(paddle.Width()/2) && paddle.Direction() < 0 {
		b.dx = -50 + -(8 * (paddleX + float64(paddle.Width()/2) - b.x))
	} else if b.x > paddleX+float64(paddle.Width()/2) && paddle.Direction() > 0 {
		// else if ball hits the paddle on the right side while the paddle is moving right
		b.dx = 50 + (8 * math.Abs(paddleX+float64(paddle.Width()/2)-b.x))
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

func (b *Ball) BoundingBox() image.Rectangle {
	return image.Rect(int(b.x), int(b.y), int(b.x)+constants.BallWidth, int(b.y)+constants.BallHeight)
}
