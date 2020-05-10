package particles

import (
	"games50-go/internal/utils"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Particle struct {
	image         *ebiten.Image
	remainingLife float64
	lifetime      float64
	position      Position
	accelerationX float64
	accelerationY float64
	dy            float64
	dx            float64
	activeColour  color.Color
	colours       []color.Color
}

func NewParticle(image *ebiten.Image, lifetime float64, position Position, acceleration Acceleration, colours []color.Color) *Particle {
	return &Particle{
		image:         image,
		lifetime:      lifetime,
		remainingLife: lifetime,
		position:      position,
		accelerationX: utils.RandomFloatInRange(acceleration.MinX, acceleration.MaxX),
		accelerationY: utils.RandomFloatInRange(acceleration.MinY, acceleration.MaxY),
		colours:       colours,
	}
}

func (p *Particle) Update() {
	delta := math.Max(0, 1/ebiten.CurrentTPS())
	p.remainingLife -= delta

	// update the velocity of the particle
	p.dx += p.accelerationX * delta
	p.dy += p.accelerationY * delta

	// update the position of the particle
	p.position.X += p.dx * delta
	p.position.Y += p.dy * delta

	// update the colour according to the current lifetime
	if len(p.colours) > 0 {
		t := 1 - p.remainingLife/p.lifetime
		s := t * float64(len(p.colours)-1)
		colourIndex := int(s)
		k := colourIndex + 1
		if colourIndex == len(p.colours)-1 {
			k = colourIndex
		}

		// interpolate colours
		s -= float64(colourIndex)
		r, g, b, a := p.colours[colourIndex].RGBA()
		r, g, b, a = r/0x101, g/0x101, b/0x101, a/0x101
		newR := float64(r) * (1 - s)
		newG := float64(g) * (1 - s)
		newB := float64(b) * (1 - s)
		newA := float64(a) * (1 - s)

		r, g, b, a = p.colours[k].RGBA()
		r, g, b, a = r/0x101, g/0x101, b/0x101, a/0x101
		nextR := float64(r) * s
		nextG := float64(g) * s
		nextB := float64(b) * s
		nextA := float64(a) * s

		addedR := newR + nextR
		addedG := newG + nextG
		addedB := newB + nextB
		addedA := newA + nextA

		p.activeColour = color.RGBA{R: uint8(addedR), G: uint8(addedG), B: uint8(addedB), A: uint8(addedA)}
	}
}

func (p *Particle) OutOfLife() bool {
	return p.remainingLife <= 0
}

func (p *Particle) Render(screen *ebiten.Image) {
	particleOptions := &ebiten.DrawImageOptions{}
	particleOptions.GeoM.Translate(p.position.X, p.position.Y)

	if p.activeColour != nil {
		r, g, b, a := p.activeColour.RGBA()
		r, g, b, a = r/0x101, g/0x101, b/0x101, a/0x101

		particleOptions.ColorM.Scale(float64(r)/0xff, float64(g)/0xff, float64(b)/0xff, float64(a)/0xff)
	}

	screen.DrawImage(p.image, particleOptions)
}
