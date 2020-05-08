package states

import (
	"fmt"
	"games50-go/breakout/assets"
	"games50-go/breakout/constants"
	"games50-go/internal/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func renderHealth(health int, screen *ebiten.Image) {
	healthX := constants.ScreenWidth - 100

	for i := 0; i < health; i++ {
		heartOptions := &ebiten.DrawImageOptions{}
		heartOptions.GeoM.Translate(float64(healthX), 4)
		screen.DrawImage(assets.GetSprite("hearts", "full"), heartOptions)

		healthX += constants.HeartWidth
	}

	for i := 0; i < constants.MaxHearts-health; i++ {
		heartOptions := &ebiten.DrawImageOptions{}
		heartOptions.GeoM.Translate(float64(healthX), 4)
		screen.DrawImage(assets.GetSprite("hearts", "empty"), heartOptions)

		healthX += constants.HeartWidth
	}
}

func renderScore(score int, screen *ebiten.Image) {
	utils.DrawText(screen, "Score:", constants.ScreenWidth-60, 5, utils.TextOptions{
		Font:  assets.GetFont("small"),
		Color: color.White,
	})
	utils.DrawText(screen, fmt.Sprintf("%d", score), constants.ScreenWidth-10, 5, utils.TextOptions{
		Font:            assets.GetFont("small"),
		Color:           color.White,
		HorizontalAlign: utils.RightAlign,
	})
}
