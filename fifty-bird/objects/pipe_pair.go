package objects

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type PipePair struct {
	Top    Pipe
	Bottom Pipe
}

func NewPipePair(screen *ebiten.Image) *PipePair {
	screenWidth, screenHeight := screen.Size()

	return &PipePair{
		Bottom: NewPipe(float64(screenWidth), float64(screenHeight/2)),
		Top:    NewPipe(float64(screenWidth), float64(screenHeight/2)),
	}
}

func (p *PipePair) Update(screen *ebiten.Image) {
	p.Top.Update()
	p.Bottom.Update()
}

func (p *PipePair) Render(screen *ebiten.Image) {
	p.Top.Render(screen, topPipeDrawingOptions(p.Top.width))
	p.Bottom.Render(screen, &ebiten.DrawImageOptions{})
}

func (p *PipePair) IsOutOfScreen() bool {
	return p.Top.x < float64(-p.Top.width)
}

func topPipeDrawingOptions(pipeWidth int) *ebiten.DrawImageOptions {
	topPipeOptions := &ebiten.DrawImageOptions{}
	topPipeOptions.GeoM.Rotate(math.Pi) // rotate 180 degrees in radian
	topPipeOptions.GeoM.Translate(float64(pipeWidth), -90)

	return topPipeOptions
}
