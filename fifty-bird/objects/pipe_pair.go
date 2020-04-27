package objects

import (
	"games50-go/internal/utils"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type PipePair struct {
	Top    Pipe
	Bottom Pipe
	Scored bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewPipePair(screen *ebiten.Image, lastY int) *PipePair {
	screenWidth, screenHeight := screen.Size()

	var y = float64(screenHeight/2) + float64(utils.RandomNumInRange(-50, 50))
	if lastY > 0 {
		// make sure the next Y is within a range
		if lastY > int(screenHeight/2) {
			y = float64(lastY+utils.RandomNumInRange(-40, 20)) - 10
		} else {
			y = float64(lastY+utils.RandomNumInRange(-20, 40)) + 10
		}
	}

	return &PipePair{
		Bottom: NewPipe(BottomPipe, float64(screenWidth), y),
		Top:    NewPipe(TopPipe, float64(screenWidth), y),
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
