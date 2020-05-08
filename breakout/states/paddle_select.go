package states

import (
	"games50-go/breakout/assets"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type PaddleColour int

const (
	Blue PaddleColour = iota
	Green
	Red
	Purple
)

func (c PaddleColour) spriteGroup() string {
	return []string{"paddles-blue", "paddles-green", "paddles-red", "paddles-purple"}[c]
}

type PaddleSelectState struct {
	currentPaddle PaddleColour
}

func (s *PaddleSelectState) Enter() {
	s.currentPaddle = Blue
}

func (s *PaddleSelectState) Update(screen *ebiten.Image) states.State {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && s.currentPaddle != Purple {
		s.currentPaddle++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && s.currentPaddle != Blue {
		s.currentPaddle--
	}
	return nil
}

func (s *PaddleSelectState) Render(screen *ebiten.Image) {
	_, screenHeight := screen.Size()

	utils.DrawText(screen, "Select your paddle with left and right!", 0, screenHeight/4, utils.TextOptions{
		Font:            assets.GetFont("medium"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})
	utils.DrawText(screen, "(Press Enter to continue!)", 0, screenHeight/3, utils.TextOptions{
		Font:            assets.GetFont("small"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})

	s.drawArrows(screen)
	s.drawPaddle(screen)
}

func (s *PaddleSelectState) drawArrows(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()

	arrowLeftOptions := &ebiten.DrawImageOptions{}
	if s.currentPaddle == Blue {
		arrowLeftOptions.ColorM.Scale(0.5, 0.5, 0.5, 0.5) // mute the colour to indicate we can't go left anymore
	}
	arrowLeftOptions.GeoM.Translate(float64(screenWidth)/4-24, float64(screenHeight-screenHeight/3))
	screen.DrawImage(assets.GetSprite("arrows", "left"), arrowLeftOptions)

	arrowRightOptions := &ebiten.DrawImageOptions{}
	if s.currentPaddle == Purple {
		arrowRightOptions.ColorM.Scale(0.5, 0.5, 0.5, 0.5) // mute the colour to indicate we can't go right anymore
	}
	arrowRightOptions.GeoM.Translate(float64(screenWidth-screenWidth/4), float64(screenHeight-screenHeight/3))
	screen.DrawImage(assets.GetSprite("arrows", "right"), arrowRightOptions)
}

func (s *PaddleSelectState) drawPaddle(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()

	paddleOptions := &ebiten.DrawImageOptions{}
	paddleOptions.GeoM.Translate(float64(screenWidth)/2-32, float64(screenHeight-screenHeight/3))
	screen.DrawImage(assets.GetSprite(s.currentPaddle.spriteGroup(), "small"), paddleOptions)
}

func (s *PaddleSelectState) Exit() {
	// do nothing
}
