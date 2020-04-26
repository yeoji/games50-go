package states

import (
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type PlayState struct {
}

func (s *PlayState) enter() {
	// do nothing
}

func (s *PlayState) update(stateMachine *StateMachine) {
	// do nothing
}

func (s *PlayState) render(screen *ebiten.Image, assets *assets.Assets) {
	utils.DrawText(screen, "Play", 0, 0, utils.TextOptions{
		Font:            assets.Fonts["hugeFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
		VerticalAlign:   utils.CenterAlign,
	})
}

func (s *PlayState) exit() {
	// do nothing
}
