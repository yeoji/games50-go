package particles

import (
	"games50-go/internal/utils"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/rs/zerolog/log"
)

type ParticleEmitter struct {
	ParticleImage *ebiten.Image
	Config        ParticleEmitterConfig
	particles     []*Particle
	lifetime      float64
	active        bool
	stopped       chan bool
}

func (e *ParticleEmitter) Emit() {
	if e.active {
		return
	}

	e.stopped = make(chan bool, 1)
	e.active = true

	log.Printf("Started emitting")

	go func() {
		for {
			select {
			case <-e.stopped:
				return
			default:
				if len(e.particles) < e.Config.MaxParticles {
					e.spawnParticle()

					time.Sleep(e.Config.Spawn.Frequency)
				}
			}
		}
	}()
}

func (e *ParticleEmitter) Update() {
	if !e.active {
		return
	}

	e.lifetime += math.Max(0, 1/ebiten.CurrentTPS())

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

	if e.lifetime >= e.Config.EmitterLife.Seconds() {
		e.active = false
		e.stopped <- true
	}
}

func (e *ParticleEmitter) removeParticles(particlesToRemove []int) {
	log.Printf("removing particles %v", particlesToRemove)
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
