package states

import (
	"fmt"
	"games50-go/fifty-bird/objects"
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type CountdownState struct {
	countdown int
}

func (s *CountdownState) enter() {
	ticker := time.Tick(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker:
				s.countdown--
				if s.countdown == 0 {
					return
				}
			}
		}
	}()
}

func (s *CountdownState) update(screen *ebiten.Image, stateMachine *StateMachine) {
	if s.countdown == 0 {
		stateMachine.Change(&PlayState{
			Bird: objects.NewBird(screen),
		})
	}
}

func (s *CountdownState) render(screen *ebiten.Image, assets *assets.Assets) {
	utils.DrawText(screen, fmt.Sprintf("%d", s.countdown), 0, 0, utils.TextOptions{
		Font:            assets.Fonts["hugeFont"],
		Color:           color.White,
		HorizontalAlign: utils.CenterAlign,
		VerticalAlign:   utils.CenterAlign,
	})
}

func (s *CountdownState) exit() {
	// do nothing
}
