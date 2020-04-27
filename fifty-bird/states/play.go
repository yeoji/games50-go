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

const GroundHeight = 16

type PlayState struct {
	Bird         objects.Bird
	PipePairs    []*objects.PipePair
	timer        float64
	pipeInterval int
	score        int
}

func (s *PlayState) enter() {
	s.pipeInterval = utils.RandomNumInRange(1, 3)
}

func (s *PlayState) update(screen *ebiten.Image, stateMachine *StateMachine) {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		stateMachine.Assets.Sounds["pause"].Play()
		stateMachine.Assets.Sounds["pause"].Rewind()
		stateMachine.Change(&PauseState{
			currentPlayState: s,
		})
	}

	_, screenHeight := screen.Size()
	s.timer += 1 / ebiten.CurrentTPS()
	if int(s.timer) > s.pipeInterval {
		var lastY = -1
		if len(s.PipePairs) > 0 {
			lastY = s.PipePairs[len(s.PipePairs)-1].Bottom.BoundingBox().Min.Y
		}
		s.PipePairs = append(s.PipePairs, objects.NewPipePair(screen, lastY))
		s.timer = 0
		s.pipeInterval = utils.RandomNumInRange(1, 3)
	}

	s.Bird.Update(stateMachine.Assets)
	if s.Bird.HasHitTheGround(screenHeight-GroundHeight) || s.Bird.HasHitAnyPipes(s.PipePairs) {
		stateMachine.Assets.Sounds["explosion"].Play()
		stateMachine.Assets.Sounds["hurt"].Play()

		stateMachine.Assets.Sounds["explosion"].Rewind()
		stateMachine.Assets.Sounds["hurt"].Rewind()

		stateMachine.Change(&ScoreState{
			score: s.score,
		})
	}

	for _, pipePair := range s.PipePairs {
		pipePair.Update(screen)
		if !pipePair.Scored && s.Bird.HasPassedPipes(pipePair) {
			stateMachine.Assets.Sounds["score"].Play()
			stateMachine.Assets.Sounds["score"].Rewind()
			pipePair.Scored = true
			s.score++
		}
	}
	if len(s.PipePairs) > 0 && s.PipePairs[0].IsOutOfScreen() {
		s.PipePairs = s.PipePairs[1:]
	}
}

func (s *PlayState) render(screen *ebiten.Image, assets *assets.Assets) {
	s.Bird.Render(screen)
	for _, pipePair := range s.PipePairs {
		pipePair.Render(screen)
	}

	utils.DrawText(screen, fmt.Sprintf("Score: %d", s.score), 8, 8, utils.TextOptions{
		Font:  assets.Fonts["flappyFont"],
		Color: color.White,
	})
}

func (s *PlayState) exit() {
	// do nothing
}
