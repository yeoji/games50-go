package states

import (
	"fmt"
	"games50-go/fifty-bird/assets/art"
	"games50-go/fifty-bird/objects"
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type ScoreState struct {
	score int
	medal *ebiten.Image
}

func (s *ScoreState) enter() {
	if s.score < 10 {
		s.medal = assets.LoadImage(art.BronzeMedal_png)
	} else if s.score < 15 {
		s.medal = assets.LoadImage(art.SilverMedal_png)
	} else {
		s.medal = assets.LoadImage(art.GoldMedal_png)
	}
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

	s.drawMedal(screen)

	_, medalHeight := s.medal.Size()
	utils.DrawText(screen, "Press Enter to play again!", 0, 160+medalHeight, utils.TextOptions{
		Font:            assets.Fonts["mediumFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})
}

func (s *ScoreState) drawMedal(screen *ebiten.Image) {
	screenWidth, _ := screen.Size()
	medalWidth, _ := s.medal.Size()

	medalOptions := &ebiten.DrawImageOptions{}
	medalOptions.GeoM.Translate(float64(screenWidth/2-medalWidth/2), 130)
	screen.DrawImage(s.medal, medalOptions)
}

func (s *ScoreState) exit() {
	// do nothing
}
