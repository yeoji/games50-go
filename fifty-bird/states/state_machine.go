package states

import (
	"games50-go/internal/assets"

	"github.com/hajimehoshi/ebiten"
)

type StateMachine struct {
	Current State
	Assets  *assets.Assets
}

type State interface {
	enter()
	update(screen *ebiten.Image, stateMachine *StateMachine)
	render(screen *ebiten.Image, assets *assets.Assets)
	exit()
}

func (sm *StateMachine) Change(state State) {
	sm.Current.exit()
	sm.Current = state
	sm.Current.enter()
}

func (sm *StateMachine) Update(screen *ebiten.Image) {
	sm.Current.update(screen, sm)
}

func (sm *StateMachine) Render(screen *ebiten.Image) {
	sm.Current.render(screen, sm.Assets)
}
