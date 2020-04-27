package states

import (
	"games50-go/fifty-bird/assets/art"
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type PauseState struct {
	pauseIcon        *ebiten.Image
	currentPlayState *PlayState
}

func (s *PauseState) enter() {
	s.pauseIcon = assets.LoadImage(art.Pause_png)
}

func (s *PauseState) update(screen *ebiten.Image, stateMachine *StateMachine) {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		// resume the game
		stateMachine.Change(s.currentPlayState)
	}
}

func (s *PauseState) render(screen *ebiten.Image, assets *assets.Assets) {
	utils.DrawText(screen, "Game Paused", 0, 64, utils.TextOptions{
		Font:            assets.Fonts["flappyFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	screenWidth, screenHeight := screen.Size()
	iconWidth, iconHeight := s.pauseIcon.Size()

	pauseOptions := &ebiten.DrawImageOptions{}
	pauseOptions.GeoM.Translate(float64(screenWidth/2-iconWidth/2), float64(screenHeight/2-iconHeight/2))
	screen.DrawImage(s.pauseIcon, pauseOptions)
}

func (s *PauseState) exit() {
	// do nothing
}
