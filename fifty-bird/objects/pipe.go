package objects

import (
	"games50-go/internal/assets"
	"image"

	"github.com/hajimehoshi/ebiten"
)

const PipeScrollingSpeed = -60

type Pipe struct {
	image  *ebiten.Image
	x      float64
	y      float64
	width  int
	height int
}

func NewPipe(x float64, y float64) Pipe {
	pipeImage := assets.LoadImage("assets/art/pipe.png")
	pipeWidth, pipeHeight := pipeImage.Size()

	return Pipe{
		image:  pipeImage,
		x:      x,
		y:      y,
		width:  pipeWidth,
		height: pipeHeight,
	}
}

func (p *Pipe) Update() {
	p.x += PipeScrollingSpeed * 1 / ebiten.CurrentTPS()
}

func (p *Pipe) Render(screen *ebiten.Image, pipeOptions *ebiten.DrawImageOptions) {
	pipeOptions.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, pipeOptions)
}

func (p *Pipe) BoundingBox() image.Rectangle {
	return image.Rect(int(p.x), int(p.y), int(p.x)+p.width, int(p.y)+p.height)
}
