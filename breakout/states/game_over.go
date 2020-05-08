package states

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type GameOverState struct {
	score int
}

func (s *GameOverState) Enter() {
	// do nothing
}

func (s *GameOverState) Update(screen *ebiten.Image) states.State {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return &MenuState{}
	}
	// TODO high scores

	return nil
}

func (s *GameOverState) Render(screen *ebiten.Image) {
	utils.DrawText(screen, "GAME OVER", 0, constants.ScreenHeight/3, utils.TextOptions{
		Font:            assets.GetFont("large"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	mediumFontCenterAlign := utils.TextOptions{
		Font:            assets.GetFont("medium"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	}
	utils.DrawText(screen, fmt.Sprintf("Final Score: %d", s.score), 0, constants.ScreenHeight/2, mediumFontCenterAlign)
	utils.DrawText(screen, "Press Enter!", 0, constants.ScreenHeight-constants.ScreenHeight/4, mediumFontCenterAlign)
}

func (s *GameOverState) Exit() {
	// do nothing
}
