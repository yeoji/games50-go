package objects

import (
	"games50-go/internal/assets"

	"github.com/hajimehoshi/ebiten"
)

const PipeScrollingSpeed = -60

type Pipe struct {
	image *ebiten.Image
	x     float64
	y     float64
	width int
}

func NewPipe(x float64, y float64) Pipe {
	pipeImage := assets.LoadImage("assets/art/pipe.png")
	pipeWidth, _ := pipeImage.Size()

	return Pipe{
		image: pipeImage,
		x:     x,
		y:     y,
		width: pipeWidth,
	}
}

func (p *Pipe) Update() {
	p.x += PipeScrollingSpeed * 1 / ebiten.CurrentTPS()
}

func (p *Pipe) Render(screen *ebiten.Image, pipeOptions *ebiten.DrawImageOptions) {
	pipeOptions.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, pipeOptions)
}
