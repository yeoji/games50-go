package particles

import (
	"games50-go/internal/utils"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type ParticleEmitter struct {
	ParticleImage *ebiten.Image
	Config        ParticleEmitterConfig
	particles     []*Particle
	active        bool
	stopped       chan bool
}

func (e *ParticleEmitter) Emit() {
	if e.active {
		return
	}

	e.stopped = make(chan bool, 1)
	e.active = true
	go func() {
		for {
			select {
			case <-e.stopped:
				return
			default:
				if len(e.particles) < e.Config.MaxParticles {
					e.spawnParticle()

					time.Sleep(time.Duration(e.Config.Spawn.Frequency) * time.Second)
				}
			}
		}
	}()
}

func (e *ParticleEmitter) Update() {
	if !e.active {
		return
	}

	e.Config.EmitterLife -= 1 / ebiten.CurrentTPS()

	var particlesToRemove []int
	for index, particle := range e.particles {
		particle.Update()

		if particle.OutOfLife() {
			particlesToRemove = append(particlesToRemove, index)
		}
	}

	if len(particlesToRemove) > 0 {
		e.removeParticles(particlesToRemove)
	}

	if e.Config.EmitterLife <= 0 {
		e.active = false
		e.stopped <- true
	}
}

func (e *ParticleEmitter) removeParticles(particlesToRemove []int) {
	var updatedParticles []*Particle
	updatedParticles = append(updatedParticles, e.particles[:particlesToRemove[0]]...)

	for i := 1; i < len(particlesToRemove); i++ {
		updatedParticles = append(updatedParticles, e.particles[particlesToRemove[i-1]+1:particlesToRemove[i]]...)
	}

	e.particles = updatedParticles
}

func (e *ParticleEmitter) spawnParticle() {
	// spawn new particle
	lifetime := utils.RandomFloatInRange(e.Config.Lifetime.Min, e.Config.Lifetime.Max)

	position := e.Config.Spawn.SpawnPosition()
	e.particles = append(e.particles, NewParticle(e.ParticleImage, lifetime, position, e.Config.Acceleration, e.Config.Colours))
}

func (e *ParticleEmitter) Render(screen *ebiten.Image) {
	if !e.active {
		return
	}

	for _, particle := range e.particles {
		particle.Render(screen)
	}
}
