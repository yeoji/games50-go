package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const BackgroundScrollSpeed = -30
const BackgroundLoopingPoint = 413

const GroundScrollSpeed = -60

type Scene struct {
	Background    *ebiten.Image
	Ground        *ebiten.Image
	backgroundPos float64
	groundPos     float64
	scrolling     bool
}

func (s *Scene) update() {
	if s.scrolling && ebiten.CurrentTPS() > 0 {
		s.backgroundPos = math.Mod((s.backgroundPos + BackgroundScrollSpeed*1/ebiten.CurrentTPS()), BackgroundLoopingPoint)
		s.groundPos = math.Mod((s.groundPos + GroundScrollSpeed*1/ebiten.CurrentTPS()), SCREEN_WIDTH)
	}
}

func (s *Scene) drawBackground(screen *ebiten.Image) {
	backgroundOptions := &ebiten.DrawImageOptions{}
	backgroundOptions.GeoM.Translate(s.backgroundPos, 0)
	screen.DrawImage(s.Background, backgroundOptions)
}

func (s *Scene) drawGround(screen *ebiten.Image) {
	_, groundHeight := s.Ground.Size()

	groundOptions := &ebiten.DrawImageOptions{}
	groundOptions.GeoM.Translate(s.groundPos, float64(SCREEN_HEIGHT-groundHeight))
	screen.DrawImage(s.Ground, groundOptions)
}
