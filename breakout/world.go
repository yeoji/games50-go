package main

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"
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
	w.renderFPS(screen)

	w.StateMachine.Render(screen)
}

func (w *World) drawBackground(screen *ebiten.Image) {
	backgroundOptions := &ebiten.DrawImageOptions{}
	backgroundOptions.GeoM.Scale(math.Ceil(constants.ScreenWidth/float64(w.Background.Bounds().Dx())), math.Ceil((constants.ScreenHeight)/float64(w.Background.Bounds().Dy())))
	screen.DrawImage(w.Background, backgroundOptions)
}

func (w *World) renderFPS(screen *ebiten.Image) {
	utils.DrawText(screen, fmt.Sprintf("FPS: %d", int(math.Ceil(ebiten.CurrentTPS()))), 5, 5, utils.TextOptions{
		Font:  assets.GetFont("small"),
		Color: color.RGBA{0, 255, 0, 255},
	})
}
