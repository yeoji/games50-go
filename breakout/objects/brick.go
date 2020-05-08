package objects

import (
	"fmt"
	"games50-go/breakout/assets"

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
	x      float64
	y      float64
	tier   brickTier
	colour colour
}

func NewBrick(x float64, y float64, tier brickTier, colour colour) Brick {
	return Brick{
		x:      x,
		y:      y,
		tier:   tier,
		colour: colour,
	}
}

func (b *Brick) Render(screen *ebiten.Image) {
	brickOptions := &ebiten.DrawImageOptions{}
	brickOptions.GeoM.Translate(b.x, b.y)

	screen.DrawImage(assets.GetSprite(fmt.Sprintf("bricks-%s", b.colour.string()), b.tier.string()), brickOptions)
}
