package states

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/breakout/data"
	"games50-go/breakout/objects"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type HighScoresState struct {
	highScores []*objects.HighScore
}

func (s *HighScoresState) Enter() {
	s.highScores = data.GetHighScores()
}

func (s *HighScoresState) Update(screen *ebiten.Image) states.State {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return &MenuState{}
	}
	return nil
}

func (s *HighScoresState) Render(screen *ebiten.Image) {
	utils.DrawText(screen, "High Scores", 0, 20, utils.TextOptions{
		Font:            assets.GetFont("large"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	mediumFontRightAlign := utils.TextOptions{
		Font:            assets.GetFont("medium"),
		Color:           color.White,
		HorizontalAlign: utils.RightAlign,
	}

	rank := 1
	// draw the high scores so far
	for _, highScore := range s.highScores {
		utils.DrawText(screen, fmt.Sprintf("%d.", rank), constants.ScreenWidth/4, 60+rank*13, utils.TextOptions{
			Font:  assets.GetFont("medium"),
			Color: color.White,
		})

		utils.DrawText(screen, highScore.Name, constants.ScreenWidth/4+38+50, 60+rank*13, mediumFontRightAlign)
		utils.DrawText(screen, fmt.Sprintf("%d", highScore.Score), constants.ScreenWidth/2+100, 60+rank*13, mediumFontRightAlign)

		rank++
	}
	// draw the empty scores
	for rank <= 10 {
		utils.DrawText(screen, fmt.Sprintf("%d.", rank), constants.ScreenWidth/4, 60+rank*13, utils.TextOptions{
			Font:  assets.GetFont("medium"),
			Color: color.White,
		})

		utils.DrawText(screen, "---", constants.ScreenWidth/4+38+50, 70+rank*13, mediumFontRightAlign)
		utils.DrawText(screen, "---", constants.ScreenWidth/2+100, 70+rank*13, mediumFontRightAlign)

		rank++
	}

	utils.DrawText(screen, "Press Escape to return to the main menu!", 0, constants.ScreenHeight-18, utils.TextOptions{
		Font:            assets.GetFont("small"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})
}

func (s *HighScoresState) Exit() {
	// do nothing
}
