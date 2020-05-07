package main

import (
	"games50-go/internal/states"
	"math"

	"github.com/hajimehoshi/ebiten"
)

type World struct {
	Background   *ebiten.Image
	StateMachine *states.StateMachine
}

func (w *World) update(screen *ebiten.Image) {
	w.StateMachine.Update(screen)
}

func (w *World) render(screen *ebiten.Image) {
	w.drawBackground(screen)

	w.StateMachine.Render(screen)
}

func (w *World) drawBackground(screen *ebiten.Image) {
	backgroundOptions := &ebiten.DrawImageOptions{}
	backgroundOptions.GeoM.Scale(math.Ceil(SCREEN_WIDTH/float64(w.Background.Bounds().Dx())), math.Ceil((SCREEN_HEIGHT)/float64(w.Background.Bounds().Dy())))
	screen.DrawImage(w.Background, backgroundOptions)
}
