package states

import (
	"github.com/hajimehoshi/ebiten"
)

type State interface {
	Enter()

	// Update returns a next state, if there is one
	Update(screen *ebiten.Image) State

	Render(screen *ebiten.Image)
	Exit()
}

type StateMachine struct {
	Current State
}

func (sm *StateMachine) Change(state State) {
	sm.Current.Exit()
	sm.Current = state
	sm.Current.Enter()
}

func (sm *StateMachine) Update(screen *ebiten.Image) {
	nextState := sm.Current.Update(screen)
	if nextState != nil {
		sm.Change(nextState)
	}
}

func (sm *StateMachine) Render(screen *ebiten.Image) {
	sm.Current.Render(screen)
}
