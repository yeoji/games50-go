package paddles

import (
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"image"

	"github.com/hajimehoshi/ebiten"
)

const PaddleSpeed = 200

type Paddle struct {
	Colour PaddleColour
	Size   PaddleSize
	x      float64
	y      float64
	dx     float64
}

func NewPaddle(colour PaddleColour, size PaddleSize) Paddle {
	paddle := Paddle{
		Colour: colour,
		Size:   size,
	}
	paddle.x = float64(constants.ScreenWidth/2 - paddle.Width()/2)
	paddle.y = float64(constants.ScreenHeight - 32)
	return paddle
}

func (p *Paddle) MoveLeft() {
	p.dx = -PaddleSpeed
	p.x += p.dx * 1 / ebiten.CurrentTPS()
	if p.x < 0 {
		p.x = 0
	}
}

func (p *Paddle) MoveRight() {
	p.dx = PaddleSpeed
	p.x += p.dx * 1 / ebiten.CurrentTPS()
	if int(p.x)+p.Width() > constants.ScreenWidth {
		p.x = float64(constants.ScreenWidth - p.Width())
	}
}

func (p *Paddle) Direction() float64 {
	return p.dx
}

func (p *Paddle) Position() (float64, float64) {
	return p.x, p.y
}

func (p *Paddle) Width() int {
	return assets.GetSprite(p.Colour.SpriteGroup(), p.Size.String()).Bounds().Dx()
}

func (p *Paddle) Render(screen *ebiten.Image) {
	paddleOptions := &ebiten.DrawImageOptions{}
	paddleOptions.GeoM.Translate(p.x, p.y)

	screen.DrawImage(assets.GetSprite(p.Colour.SpriteGroup(), p.Size.String()), paddleOptions)
}

func (p *Paddle) BoundingBox() image.Rectangle {
	return image.Rect(int(p.x), int(p.y), int(p.x)+p.Width(), int(p.y)+constants.PaddleHeight)
}
