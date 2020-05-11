package particles

import (
	"image/color"
	"time"
)

// ParticleEmitterConfig - These configs define how the particle emitter should behave
type ParticleEmitterConfig struct {
	// The max number of particles that will be active at any one time
	MaxParticles int
	// The lifetime range of a particle (how long it should stay on the screen)
	Lifetime Range
	// The X, Y acceleration of a particle - affects the velocity
	Acceleration Acceleration
	Spawn        Spawn
	// The colours to apply to the particle (interpolated depending on the particle lifetime)
	Colours []color.Color
	// How long the emitter should keep emitting for. If this is empty, the emitter will stop after max particles has been reached
	EmitterLife time.Duration
}

type Range struct {
	Min float64
	Max float64
}

type Acceleration struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

type Position struct {
	X float64
	Y float64
}

func (e *ParticleEmitter) SetColours(colours []color.Color) {
	e.Config.Colours = colours
}
