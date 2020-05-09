package objects

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"image"

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
	InPlay bool
	Locked bool
}

func NewBrick(x float64, y float64, tier brickTier, colour colour) Brick {
	return Brick{
		x:      x,
		y:      y,
		tier:   tier,
		colour: colour,
		InPlay: true,
	}
}

func (b *Brick) Hit() {
	assets.PlaySound("brick_hit")

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

func (b *Brick) BoundingBox() image.Rectangle {
	return image.Rect(int(b.x), int(b.y), int(b.x)+constants.BrickWidth, int(b.y)+constants.BrickHeight)
}
