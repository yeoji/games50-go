package objects

import (
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/internal/utils"
	"image"

	"github.com/hajimehoshi/ebiten"
)

const ExtraBallsPowerup = "extra-balls"
const KeyPowerup = "key"

const PowerupFallSpeed = 20

type Powerup struct {
	PowerupType string
	x           float64
	y           float64
}

func NewPowerup(powerupType string) *Powerup {
	return &Powerup{
		PowerupType: powerupType,
		x:           float64(utils.RandomNumInRange(40, 400)),
		y:           float64(utils.RandomNumInRange(70, 90)),
	}
}

func (p *Powerup) DriftDown() {
	p.y += PowerupFallSpeed * 1 / ebiten.CurrentTPS()
}

func (p *Powerup) Render(screen *ebiten.Image) {
	powerupOptions := &ebiten.DrawImageOptions{}
	powerupOptions.GeoM.Translate(p.x, p.y)

	screen.DrawImage(assets.GetSprite("powerups", p.PowerupType), powerupOptions)
}

func (p *Powerup) BoundingBox() image.Rectangle {
	return image.Rect(int(p.x), int(p.y), int(p.x)+constants.PowerupWidth, int(p.y)+constants.PowerupHeight)
}
