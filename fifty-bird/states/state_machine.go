package states

import "github.com/hajimehoshi/ebiten"

type StateMachine struct {
	Current State
}

type State interface {
	enter()
	update()
	render(screen *ebiten.Image)
	exit()
}

func (sm *StateMachine) Change(state State) {
	sm.Current.exit()
	sm.Current = state
	sm.Current.enter()
}

func (sm *StateMachine) Update() {
	sm.Current.update()
}

func (sm *StateMachine) Render(screen *ebiten.Image) {
	sm.Current.render(screen)
}
