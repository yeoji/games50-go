package states

import (
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/breakout/objects"
	"games50-go/breakout/objects/paddles"
	"games50-go/internal/states"
	"games50-go/internal/utils"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type PlayState struct {
	paddle             paddles.Paddle
	balls              []*objects.Ball
	level              objects.Level
	score              int
	health             int
	recoverPoints      int
	paused             bool
	powerup            *objects.Powerup
	hasKeyPowerup      bool
	powerupSpawnFinish chan bool
}

func (s *PlayState) Enter() {
	s.balls[0].Serve()
	s.recoverPoints = int(5000 * math.Pow(2, float64(s.score)/5000))

	s.powerupSpawnFinish = make(chan bool, 1)
	go s.startPowerupSpawn()
}

func (s *PlayState) Update(screen *ebiten.Image) states.State {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.paused = !s.paused
		assets.PlaySound("pause")
	}
	if s.paused {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.paddle.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.paddle.MoveRight()
	}

	if s.powerup != nil {
		s.powerup.DriftDown()

		if utils.Collides(&s.paddle, s.powerup) {
			s.activatePowerup()
		}
	}

	var updatedBalls []*objects.Ball
	for i, ball := range s.balls {
		ball.Move()

		if utils.Collides(ball, &s.paddle) {
			assets.PlaySound("paddle_hit")
			ball.BounceOffPaddle(&s.paddle)
		}

		s.checkBrickCollision(ball)
		if s.allBricksCleared() {
			assets.PlaySound("victory")
			return &ServeState{
				paddle: s.paddle,
				level:  objects.NewLevel(s.level.Number + 1),
				score:  s.score,
				health: s.health,
			}
		}

		if ball.IsOutOfScreen() {
			if len(s.balls) == 1 {
				assets.PlaySound("hurt")

				s.paddle.Shrink()

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

			updatedBalls = append(s.balls[:i], s.balls[i+1:]...)
		}
	}
	if len(updatedBalls) > 0 {
		s.balls = updatedBalls
	}

	if s.score > s.recoverPoints {
		s.health = int(math.Min(constants.MaxHearts, float64(s.health)+1))
		s.recoverPoints *= 2

		assets.PlaySound("recover")
		s.paddle.Grow()
	}

	for _, brick := range s.level.Bricks {
		brick.Update()
	}

	return nil
}

func (s *PlayState) activatePowerup() {
	switch s.powerup.PowerupType {
	case objects.ExtraBallsPowerup:
		s.spawnBalls(2)
		break
	case objects.KeyPowerup:
		s.hasKeyPowerup = true
		break
	}

	s.powerup = nil
}

func (s *PlayState) spawnBalls(numBalls int) {
	for i := 0; i < numBalls; i++ {
		newBall := objects.NewBall()
		newBall.FollowPaddle(&s.paddle)
		newBall.Serve()

		s.balls = append(s.balls, newBall)
	}
}

func (s *PlayState) checkBrickCollision(ball *objects.Ball) {
	for _, brick := range s.level.Bricks {
		if utils.Collides(ball, brick) {
			score := ball.HitBrick(brick, s.hasKeyPowerup)
			s.score += score
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
	if s.paused {
		utils.DrawText(screen, "PAUSED", 0, constants.ScreenHeight/2-16, utils.TextOptions{
			Font:            assets.GetFont("large"),
			Color:           color.White,
			HorizontalAlign: utils.CenterAlign,
		})
	}

	s.paddle.Render(screen)

	for _, ball := range s.balls {
		ball.Render(screen)
	}

	if s.powerup != nil {
		s.powerup.Render(screen)
	}

	for _, brick := range s.level.Bricks {
		brick.Render(screen)
	}
	for _, brick := range s.level.Bricks {
		brick.RenderParticles(screen)
	}

	renderScore(s.score, screen)
	renderHealth(s.health, screen)
}

func (s *PlayState) startPowerupSpawn() {
	ticker := time.Tick(30 * time.Second)
	for {
		select {
		case <-s.powerupSpawnFinish:
			return
		case <-ticker:
			if !s.paused {
				spawnKeyPowerup := false
				if s.level.HasLockedBrick() && !s.hasKeyPowerup {
					spawnKeyPowerup = utils.RandomNumInRange(1, 2) == 1
				}

				if spawnKeyPowerup {
					s.powerup = objects.NewPowerup(objects.KeyPowerup)
				} else {
					s.powerup = objects.NewPowerup(objects.ExtraBallsPowerup)
				}
			}
		}
	}
}

func (s *PlayState) Exit() {
	s.powerupSpawnFinish <- true
}
