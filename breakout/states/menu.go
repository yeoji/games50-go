package states

import (
	"games50-go/breakout/assets"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const StartOption = 0
const HighScoresOption = 1

type MenuState struct {
	highlighted int // the menu item that is highlighted
}

func (s *MenuState) Enter() {
	// do nothing
}

func (s *MenuState) Update(screen *ebiten.Image) states.State {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.highlighted = (s.highlighted + 1) % 2
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch s.highlighted {
		case StartOption:
			return &PaddleSelectState{}
		case HighScoresOption:
			break
		}
	}

	return nil
}

func (s *MenuState) Render(screen *ebiten.Image) {
	_, screenHeight := screen.Size()

	utils.DrawText(screen, "BREAKOUT", 0, screenHeight/3, utils.TextOptions{
		Font:            assets.GetFont("large"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	var startOptionColour color.Color = color.White
	if s.highlighted == StartOption {
		startOptionColour = color.RGBA{103, 255, 255, 255}
	}
	utils.DrawText(screen, "START", 0, screenHeight/2+70, utils.TextOptions{
		Font:            assets.GetFont("medium"),
		Color:           startOptionColour,
		HorizontalAlign: utils.CenterAlign,
	})

	var highScoresOptionColour color.Color = color.White
	if s.highlighted == HighScoresOption {
		highScoresOptionColour = color.RGBA{103, 255, 255, 255}
	}
	utils.DrawText(screen, "HIGH SCORES", 0, screenHeight/2+90, utils.TextOptions{
		Font:            assets.GetFont("medium"),
		Color:           highScoresOptionColour,
		HorizontalAlign: utils.CenterAlign,
	})
}

func (s *MenuState) Exit() {
	// do nothing
}
