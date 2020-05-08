package states

import (
	"games50-go/breakout/objects"
	"games50-go/breakout/objects/paddles"
	"games50-go/internal/states"
	"games50-go/internal/utils"

	"github.com/hajimehoshi/ebiten"
)

type PlayState struct {
	paddle paddles.Paddle
	balls  []*objects.Ball
	level  objects.Level
	score  int
	health int
}

func (s *PlayState) Enter() {
	s.balls[0].Serve()
}

func (s *PlayState) Update(screen *ebiten.Image) states.State {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.paddle.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.paddle.MoveRight()
	}

	for _, ball := range s.balls {
		ball.Move()

		if utils.Collides(ball, &s.paddle) {
			ball.BounceOffPaddle(&s.paddle)
		}

		s.checkBrickCollision(ball)
		if s.allBricksCleared() {
			return &ServeState{
				paddle: s.paddle,
				level:  objects.NewLevel(s.level.Number + 1),
				score:  s.score,
				health: s.health,
			}
		}

		if ball.IsOutOfScreen() {
			s.health--
			if s.health == 0 {
				return &GameOverState{
					score: s.score,
				}
			}
			return &ServeState{
				paddle: s.paddle,
				level:  s.level,
				score:  s.score,
				health: s.health,
			}
		}
	}

	return nil
}

func (s *PlayState) checkBrickCollision(ball *objects.Ball) {
	for _, brick := range s.level.Bricks {
		if utils.Collides(ball, brick) {
			ball.HitBrick(brick)

			s.score += brick.Score()
			break
		}
	}
}

func (s *PlayState) allBricksCleared() bool {
	for _, brick := range s.level.Bricks {
		if brick.InPlay {
			return false
		}
	}
	return true
}

func (s *PlayState) Render(screen *ebiten.Image) {
	s.paddle.Render(screen)

	for _, ball := range s.balls {
		ball.Render(screen)
	}

	for _, brick := range s.level.Bricks {
		brick.Render(screen)
	}

	renderScore(s.score, screen)
	renderHealth(s.health, screen)
}

func (s *PlayState) Exit() {
	// do nothing
}
