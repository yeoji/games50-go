package objects

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/internal/particles"
	"image"
	"image/color"

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

var paletteColour = map[colour]color.Color{
	Blue:   color.RGBA{99, 155, 255, 255},
	Green:  color.RGBA{106, 190, 47, 255},
	Red:    color.RGBA{217, 87, 99, 255},
	Purple: color.RGBA{215, 123, 186, 255},
	Yellow: color.RGBA{251, 242, 54, 255},
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
	b := Brick{
		x:      x,
		y:      y,
		tier:   tier,
		colour: colour,
		InPlay: true,
	}
	b.newParticleEmitter()
	return b
}

func (b *Brick) Update() {
	b.pEmitter.Update()
}

func (b *Brick) Hit() {
	assets.PlaySound("brick_hit")

	paletteR, paletteG, paletteB, _ := paletteColour[b.colour].RGBA()
	paletteR, paletteG, paletteB = paletteR/0x101, paletteG/0x101, paletteB/0x101
	b.pEmitter.SetColours([]color.Color{
		color.RGBA{uint8(paletteR), uint8(paletteG), uint8(paletteB), uint8(55 * (int(b.tier) + 1))},
		color.RGBA{uint8(paletteR), uint8(paletteG), uint8(paletteB), 0},
	})
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
}

func (b *Brick) RenderParticles(screen *ebiten.Image) {
	b.pEmitter.Render(screen)
}

func (b *Brick) BoundingBox() image.Rectangle {
	return image.Rect(int(b.x), int(b.y), int(b.x)+constants.BrickWidth, int(b.y)+constants.BrickHeight)
}

func (b *Brick) newParticleEmitter() {
	paletteR, paletteG, paletteB, _ := paletteColour[b.colour].RGBA()
	paletteR, paletteG, paletteB = paletteR/0x101, paletteG/0x101, paletteB/0x101

	b.pEmitter = &particles.ParticleEmitter{
		ParticleImage: assets.GetSprite("particles", "brick-explode"),
		Config: particles.ParticleEmitterConfig{
			MaxParticles: 64,
			Lifetime:     particles.Range{Min: 0.5, Max: 1},
			Acceleration: particles.Acceleration{MinX: -70, MinY: 40, MaxX: 70, MaxY: 80},
			Spawn: particles.Spawn{
				SpawnType: particles.SpawnTypeRect,
				SpawnRect: particles.SpawnRect{
					Height: 30,
					Width:  40,
					Offset: particles.Position{X: -20, Y: -15},
				},
				Position: particles.Position{X: b.x + 16, Y: b.y + 8},
			},
		},
	}
}
