package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Ball struct {
	width  float64
	height float64
	x      float64
	y      float64
}

func (b *Ball) render(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.x, b.y, b.width, b.height, color.White)
}
