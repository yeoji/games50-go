package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type Align int

const (
	LeftAlign Align = iota
	CenterAlign
	RightAlign
)

type TextOptions struct {
	Font            font.Face
	Color           color.Color
	HorizontalAlign Align
	VerticalAlign   Align // only CenterAlign will be taken into effect
}

func DrawText(screen *ebiten.Image, content string, x int, y int, options TextOptions) {
	var actualX = x
	var actualY = y

	screenWidth, screenHeight := screen.Size()
	textWidth, textHeight := getTextBounds(content, options.Font)

	// calculate the x depending on horizontal align
	switch options.HorizontalAlign {
	case CenterAlign:
		drawingBoxWidth := screenWidth - x
		actualX = drawingBoxWidth/2 - textWidth/2
		break
	case RightAlign:
		actualX = screenWidth - textWidth
		break
	}

	// calculate the y depending on vertical align
	switch options.VerticalAlign {
	case CenterAlign:
		drawingBoxHeight := screenHeight + y
		actualY = drawingBoxHeight/2 - textHeight/2
		break
	}

	text.Draw(screen, content, options.Font, actualX, actualY, options.Color)
}

// Get the width and height of the text
func getTextBounds(content string, face font.Face) (int, int) {
	bounds, _ := font.BoundString(face, content)
	width := (bounds.Max.X - bounds.Min.X).Ceil()
	height := (bounds.Max.Y - bounds.Min.Y).Ceil()
	return width, height
}
