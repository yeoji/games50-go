package objects

import (
	"games50-go/internal/assets"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
)

const GapHeight = 90
const PipeScrollingSpeed = -60
const TopPipe = "top"
const BottomPipe = "bottom"

type Pipe struct {
	location string
	image    *ebiten.Image
	x        float64
	y        float64
	width    int
	height   int
}

func NewPipe(location string, x float64, y float64) Pipe {
	pipeImage := assets.LoadImage("assets/art/pipe.png")
	pipeWidth, pipeHeight := pipeImage.Size()

	return Pipe{
		location: location,
		image:    pipeImage,
		x:        x,
		y:        y,
		width:    pipeWidth,
		height:   pipeHeight,
	}
}

func (p *Pipe) Update() {
	p.x += PipeScrollingSpeed * 1 / ebiten.CurrentTPS()
}

func (p *Pipe) Render(screen *ebiten.Image) {
	pipeOptions := &ebiten.DrawImageOptions{}
	if p.location == TopPipe {
		pipeOptions.GeoM.Rotate(math.Pi) // rotate 180 degrees in radian
		pipeOptions.GeoM.Translate(float64(p.width), -GapHeight)
	}
	pipeOptions.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, pipeOptions)
}

func (p *Pipe) BoundingBox() image.Rectangle {
	if p.location == TopPipe {
		// make y an insanely high number so that the bird can't avoid pipes by going into the sky
		return image.Rect(int(p.x), -2000, int(p.x)+p.width, int(p.y)-GapHeight)
	}
	return image.Rect(int(p.x), int(p.y), int(p.x)+p.width, int(p.y)+p.height)
}
