package objects

import (
	"github.com/hajimehoshi/ebiten"
)

type PipePair struct {
	Top    Pipe
	Bottom Pipe
}

func NewPipePair(screen *ebiten.Image) *PipePair {
	screenWidth, screenHeight := screen.Size()

	return &PipePair{
		Bottom: NewPipe(BottomPipe, float64(screenWidth), float64(screenHeight/2)),
		Top:    NewPipe(TopPipe, float64(screenWidth), float64(screenHeight/2)),
	}
}

func (p *PipePair) Update(screen *ebiten.Image) {
	p.Top.Update()
	p.Bottom.Update()
}

func (p *PipePair) Render(screen *ebiten.Image) {
	p.Top.Render(screen)
	p.Bottom.Render(screen)
}

func (p *PipePair) IsOutOfScreen() bool {
	return p.Top.x < float64(-p.Top.width)
}
