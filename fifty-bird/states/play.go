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
	Bird  objects.Bird
	score int
}

func (s *PlayState) enter() {
	// do nothin
}

func (s *PlayState) update(screen *ebiten.Image, stateMachine *StateMachine) {
	s.Bird.Update()
}

func (s *PlayState) render(screen *ebiten.Image, assets *assets.Assets) {
	utils.DrawText(screen, fmt.Sprintf("Score: %d", s.score), 8, 8, utils.TextOptions{
		Font:  assets.Fonts["flappyFont"],
		Color: color.White,
	})

	s.Bird.Render(screen)
}

func (s *PlayState) exit() {
	// do nothing
}
