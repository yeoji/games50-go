package states

import (
	"games50-go/internal/assets"

	"github.com/hajimehoshi/ebiten"
)

type StateMachine struct {
	Current State
}

type State interface {
	enter()
	update(stateMachine *StateMachine)
	render(screen *ebiten.Image, assets *assets.Assets)
	exit()
}

func (sm *StateMachine) Change(state State) {
	sm.Current.exit()
	sm.Current = state
	sm.Current.enter()
}

func (sm *StateMachine) Update() {
	sm.Current.update(sm)
}

func (sm *StateMachine) Render(screen *ebiten.Image, assets *assets.Assets) {
	sm.Current.render(screen, assets)
}
