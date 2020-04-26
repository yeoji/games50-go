package states

import (
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type TitleScreenState struct{}

func (s *TitleScreenState) enter() {
	// do nothing
}

func (s *TitleScreenState) update() {
	// do nothing
}

func (s *TitleScreenState) render(screen *ebiten.Image, assets *assets.Assets) {
	utils.DrawText(screen, "Fifty Bird", 0, 64, utils.TextOptions{
		Font:            assets.Fonts["flappyFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	utils.DrawText(screen, "Press Enter", 0, 100, utils.TextOptions{
		Font:            assets.Fonts["mediumFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})
}

func (s *TitleScreenState) exit() {
	// do nothing
}
