package main

import (
	"games50-go/breakout/assets"
	"games50-go/internal/particles"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type TestApp struct {
	pEmitter *particles.ParticleEmitter
}

func (a *TestApp) Update(screen *ebiten.Image) error {
	a.pEmitter.Update()
	return nil
}

func (a *TestApp) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 45, 52, 255})
	a.pEmitter.Render(screen)
}

func (a *TestApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 432, 243
}

func main() {
	pEmitter := &particles.ParticleEmitter{
		ParticleImage: assets.GetSprite("particles", "brick-explode"),
		Config: particles.ParticleEmitterConfig{
			MaxParticles: 64,
			EmitterLife:  5 * time.Second,
			Lifetime:     particles.Range{Min: 0.5, Max: 1},
			Acceleration: particles.Acceleration{MinX: -15, MinY: 0, MaxX: 15, MaxY: 80},
			Spawn: particles.Spawn{
				SpawnType: particles.SpawnTypeRect,
				SpawnRect: particles.SpawnRect{
					Height: 40,
					Width:  30,
					Offset: particles.Position{X: 0, Y: 0},
				},
				Frequency: 10 * time.Millisecond,
				Position:  particles.Position{X: 50, Y: 50},
			},
			Colours: []color.Color{color.RGBA{106, 190, 47, 110}, color.RGBA{106, 190, 47, 0}},
		},
	}
	pEmitter.Emit()

	runTestApp(pEmitter)
}

func runTestApp(pEmitter *particles.ParticleEmitter) {
	ebiten.SetWindowSize(1280, 720)

	if err := ebiten.RunGame(&TestApp{
		pEmitter: pEmitter,
	}); err != nil {
		log.Fatal(err)
	}
}
