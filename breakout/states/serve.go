package states

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/breakout/objects"
	"games50-go/breakout/objects/paddles"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type ServeState struct {
	paddle paddles.Paddle
	ball   objects.Ball
	level  objects.Level
	score  int
	health int
}

func (s *ServeState) Enter() {
	s.ball = objects.NewBall()
	s.ball.FollowPaddle(&s.paddle)
}

func (s *ServeState) Update(screen *ebiten.Image) states.State {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.paddle.MoveLeft()
		s.ball.FollowPaddle(&s.paddle)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.paddle.MoveRight()
		s.ball.FollowPaddle(&s.paddle)
	}
	return nil
}

func (s *ServeState) Render(screen *ebiten.Image) {
	s.paddle.Render(screen)
	s.ball.Render(screen)

	for _, brick := range s.level.Bricks {
		brick.Render(screen)
	}

	renderScore(s.score, screen)
	renderHealth(s.health, screen)

	utils.DrawText(screen, fmt.Sprintf("Level %d", s.level.Number), 0, constants.ScreenHeight/3, utils.TextOptions{
		Font:            assets.GetFont("large"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
	})
	utils.DrawText(screen, "Press Enter to serve!", 0, 0, utils.TextOptions{
		Font:            assets.GetFont("medium"),
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
		VerticalAlign:   utils.CenterAlign,
	})
}

func (s *ServeState) Exit() {
	// do nothing
}
