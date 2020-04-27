package states

import (
	"fmt"
	"games50-go/fifty-bird/objects"
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type ScoreState struct {
	score int
}

func (s *ScoreState) enter() {
	// do nothing
}

func (s *ScoreState) update(screen *ebiten.Image, stateMachine *StateMachine) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		stateMachine.Change(&PlayState{
			Bird: objects.NewBird(screen),
		})
	}
}

func (s *ScoreState) render(screen *ebiten.Image, assets *assets.Assets) {
	utils.DrawText(screen, "Oof! You Lost!", 0, 64, utils.TextOptions{
		Font:            assets.Fonts["flappyFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	utils.DrawText(screen, fmt.Sprintf("Score: %d", s.score), 0, 100, utils.TextOptions{
		Font:            assets.Fonts["mediumFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	utils.DrawText(screen, "Press Enter to play again!", 0, 160, utils.TextOptions{
		Font:            assets.Fonts["mediumFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})
}

func (s *ScoreState) exit() {
	// do nothing
}
