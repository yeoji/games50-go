package states

import (
	"fmt"
	"games50-go/fifty-bird/objects"
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type PlayState struct {
	Bird      objects.Bird
	PipePairs []*objects.PipePair
	timer     float64
	score     int
}

func (s *PlayState) enter() {
	s.timer = 0
}

func (s *PlayState) update(screen *ebiten.Image, stateMachine *StateMachine) {
	s.timer += 1 / ebiten.CurrentTPS()
	if s.timer > 2 {
		s.PipePairs = append(s.PipePairs, objects.NewPipePair(screen))
		s.timer = 0
	}

	s.Bird.Update()

	for _, pipePair := range s.PipePairs {
		pipePair.Update(screen)
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
