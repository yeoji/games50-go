package objects

import (
	"games50-go/breakout/constants"
	"games50-go/internal/utils"
	"math"
)

const BrickPadding = 8

type Level struct {
	Number int
	Bricks []Brick
}

type RowSettings struct {
	skipPattern      bool
	alternatePattern bool
	colours          []colour
	tiers            []brickTier
}

func NewLevel(level int) Level {
	lvl := Level{
		Number: level,
	}
	lvl.generateLevelBricks()
	return lvl
}

func (l *Level) generateLevelBricks() {
	rows := utils.RandomNumInRange(1, 5)

	// ensure an odd number of columns for symmetry
	cols := utils.RandomNumInRange(7, 13)
	if cols%2 == 0 {
		cols++
	}

	highestTier := l.getHighestPossibleTier()
	highestColour := l.getHighestPossibleColour()

	for y := 0; y < rows; y++ {
		rowSettings := generateRowSettings(highestTier, highestColour)

		skipCol := utils.RandomNumInRange(1, 2) == 1
		brickTierIndex := 0
		brickColourIndex := 0
		for x := 0; x < cols; x++ {
			skipCol = !skipCol
			if rowSettings.skipPattern && skipCol {
				continue
			}

			brickX := x*constants.BrickWidth + BrickPadding + (13-cols)*16
			brick := NewBrick(float64(brickX), float64((y+1)*constants.BrickHeight), rowSettings.tiers[brickTierIndex], rowSettings.colours[brickColourIndex])
			l.Bricks = append(l.Bricks, brick)

			if rowSettings.alternatePattern {
				brickTierIndex = (brickTierIndex + 1) % 2
				brickColourIndex = (brickColourIndex + 1) % 2
			}
		}
	}
}

func (l *Level) getHighestPossibleTier() brickTier {
	highestTier := Ultra
	levelTier := brickTier(math.Floor(float64(l.Number) / 5))
	if levelTier < highestTier {
		highestTier = levelTier
	}
	return highestTier
}

func (l *Level) getHighestPossibleColour() colour {
	highestColour := Yellow
	levelColour := colour(l.Number%5 + 3)
	if levelColour < highestColour {
		highestColour = levelColour
	}
	return highestColour
}

func generateRowSettings(highestTier brickTier, highestColour colour) RowSettings {
	return RowSettings{
		skipPattern:      utils.RandomNumInRange(1, 2) == 1,
		alternatePattern: utils.RandomNumInRange(1, 2) == 1,
		colours: []colour{
			colour(utils.RandomNumInRange(0, int(highestColour))),
			colour(utils.RandomNumInRange(0, int(highestColour))),
		},
		tiers: []brickTier{
			brickTier(utils.RandomNumInRange(0, int(highestTier))),
			brickTier(utils.RandomNumInRange(0, int(highestTier))),
		},
	}
}
