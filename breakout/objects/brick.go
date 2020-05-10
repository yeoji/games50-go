package objects

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/internal/particles"
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type brickTier int

const (
	Basic brickTier = iota
	Extra
	Super
	Ultra
)

func (t brickTier) string() string {
	return []string{"basic", "extra", "super", "ultra"}[t]
}

type Brick struct {
	x        float64
	y        float64
	tier     brickTier
	colour   colour
	InPlay   bool
	Locked   bool
	pEmitter *particles.ParticleEmitter
}

func NewBrick(x float64, y float64, tier brickTier, colour colour) Brick {
	pEmitter := particles.ParticleEmitter{
		ParticleImage: assets.GetSprite("particles", "brick-explode"),
		Config: particles.ParticleEmitterConfig{
			MaxParticles: 64,
			EmitterLife:  600 * time.Millisecond,
			Lifetime:     particles.Range{Min: 0.5, Max: 1},
			Acceleration: particles.Acceleration{MinX: -15, MinY: 0, MaxX: 15, MaxY: 80},
			Spawn: particles.Spawn{
				SpawnType: particles.SpawnTypeRect,
				SpawnRect: particles.SpawnRect{
					Height: 40,
					Width:  30,
					Offset: particles.Position{X: -10, Y: -20},
				},
				Frequency: 10 * time.Millisecond,
				Position:  particles.Position{X: x + 16, Y: y + 8},
			},
			Colours: []color.Color{color.RGBA{106, 190, 47, 110}, color.RGBA{106, 190, 47, 0}},
		},
	}

	return Brick{
		x:        x,
		y:        y,
		tier:     tier,
		colour:   colour,
		InPlay:   true,
		pEmitter: &pEmitter,
	}
}

func (b *Brick) Update() {
	b.pEmitter.Update()
}

func (b *Brick) Hit() {
	assets.PlaySound("brick_hit")

	b.pEmitter.Emit()

	if b.Locked {
		b.Locked = false
		return
	}

	if int(b.colour) > int(Blue) {
		b.colour--
	} else {
		if int(b.tier) > int(Basic) {
			b.tier--
			b.colour = Yellow
		} else {
			b.InPlay = false
			assets.PlaySound("brick_destroyed")
		}
	}
}

func (b *Brick) Score() int {
	if b.Locked {
		return 2000
	}
	return int(b.tier)*200 + (int(b.colour)+1)*25
}

func (b *Brick) Render(screen *ebiten.Image) {
	if b.InPlay {
		brickOptions := &ebiten.DrawImageOptions{}
		brickOptions.GeoM.Translate(b.x, b.y)

		if b.Locked {
			screen.DrawImage(assets.GetSprite("bricks", "locked"), brickOptions)
		} else {
			screen.DrawImage(assets.GetSprite(fmt.Sprintf("bricks-%s", b.colour.string()), b.tier.string()), brickOptions)
		}
	}

	b.pEmitter.Render(screen)
}

func (b *Brick) BoundingBox() image.Rectangle {
	return image.Rect(int(b.x), int(b.y), int(b.x)+constants.BrickWidth, int(b.y)+constants.BrickHeight)
}
