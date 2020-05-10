package particles

import (
	"games50-go/internal/utils"
	"time"
)

const SpawnTypePoint = "point"
const SpawnTypeRect = "rect"
const SpawnTypeCircle = "circle"

type SpawnRect struct {
	Height float64
	Width  float64
	Offset Position
}

func (s *SpawnRect) spawnPosition(emitterPos Position) Position {
	return Position{
		X: emitterPos.X + utils.RandomFloatInRange(s.Offset.X, s.Offset.X+s.Width),
		Y: emitterPos.Y + utils.RandomFloatInRange(s.Offset.Y, s.Offset.Y+s.Height),
	}
}

type SpawnCircle struct {
	Radius float64
	Position
}

type Spawn struct {
	SpawnType   string
	Frequency   time.Duration
	Position    Position
	SpawnRect   SpawnRect
	SpawnCircle SpawnCircle
}

func (s *Spawn) SpawnPosition() Position {
	// generate position based on spawn type
	var position Position

	switch s.SpawnType {
	case SpawnTypePoint:
		position = Position{
			X: s.Position.X,
			Y: s.Position.Y,
		}
		break
	case SpawnTypeRect:
		position = s.SpawnRect.spawnPosition(s.Position)
		break
	}

	return position
}
