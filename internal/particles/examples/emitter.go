package main

import (
	"games50-go/breakout/assets/graphics"
	"games50-go/internal/assets"
	"games50-go/internal/particles"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type TestApp struct {
	pEmitter *particles.ParticleEmitter
}

func (a *TestApp) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		a.pEmitter = newParticleEmitter()
		a.pEmitter.Emit()
	}

	if a.pEmitter != nil {
		a.pEmitter.Update()
	}
	return nil
}

func (a *TestApp) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 45, 52, 255})

	if a.pEmitter != nil {
		a.pEmitter.Render(screen)
	}
}

func (a *TestApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 432, 243
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(1280, 720)

	if err := ebiten.RunGame(&TestApp{}); err != nil {
		log.Fatal(err)
	}
}

func newParticleEmitter() *particles.ParticleEmitter {
	return &particles.ParticleEmitter{
		ParticleImage: assets.LoadImage(graphics.Particle_png),
		Config: particles.ParticleEmitterConfig{
			MaxParticles: 30,
			Lifetime:     particles.Range{Min: 0.5, Max: 1},
			Acceleration: particles.Acceleration{MinX: -60, MinY: 40, MaxX: 60, MaxY: 80},
			Spawn: particles.Spawn{
				SpawnType: particles.SpawnTypeRect,
				SpawnRect: particles.SpawnRect{
					Height: 25,
					Width:  20,
					Offset: particles.Position{X: 0, Y: 0},
				},
				Position: particles.Position{X: 50, Y: 50},
			},
			Colours: []color.Color{color.RGBA{106, 190, 47, 110}, color.RGBA{106, 190, 47, 0}},
		},
	}
}
