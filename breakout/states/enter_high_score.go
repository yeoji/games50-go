package states

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/breakout/data"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type EnterHighScoreState struct {
	score       int
	name        []rune
	highlighted int
}

func (s *EnterHighScoreState) Enter() {
	s.name = []rune{'A', 'A', 'A'}
}

func (s *EnterHighScoreState) Update(screen *ebiten.Image) states.State {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && s.highlighted < 2 {
		s.highlighted++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && s.highlighted > 0 {
		s.highlighted--
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && s.name[s.highlighted] < rune('Z') {
		s.name[s.highlighted]++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && s.name[s.highlighted] > rune('A') {
		s.name[s.highlighted]--
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		data.UpdateHighScores(string(s.name), s.score)

		return &HighScoresState{}
	}
	return nil
}

func (s *EnterHighScoreState) Render(screen *ebiten.Image) {
	utils.DrawText(screen, fmt.Sprintf("Your score: %d", s.score), 0, 30, utils.TextOptions{
		Font:            assets.GetFont("medium"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	s.drawNameChar(0, screen)
	s.drawNameChar(1, screen)
	s.drawNameChar(2, screen)

	utils.DrawText(screen, "Press Enter to confirm!", 0, constants.ScreenHeight-18, utils.TextOptions{
		Font:            assets.GetFont("small"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})
}

func (s *EnterHighScoreState) drawNameChar(charPosition int, screen *ebiten.Image) {
	unhighlightedTextOption := utils.TextOptions{
		Font:          assets.GetFont("large"),
		Color:         color.White,
		VerticalAlign: utils.CenterAlign,
	}
	highlightedCharOption := utils.TextOptions{
		Font:          assets.GetFont("large"),
		Color:         color.RGBA{103, 255, 255, 255},
		VerticalAlign: utils.CenterAlign,
	}

	if s.highlighted == charPosition {
		utils.DrawText(screen, string(s.name[charPosition]), constants.ScreenWidth/2-28+charPosition*24, 10, highlightedCharOption)
	} else {
		utils.DrawText(screen, string(s.name[charPosition]), constants.ScreenWidth/2-28+charPosition*24, 10, unhighlightedTextOption)
	}
}

func (s *EnterHighScoreState) Exit() {
	// do nothing
}
