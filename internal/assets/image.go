package assets

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
)

func LoadImage(imageData []byte) *ebiten.Image {
	decoded, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		log.Fatalf("Error loading image asset: %v", err)
	}

	loadedImage, _ := ebiten.NewImageFromImage(decoded, ebiten.FilterDefault)
	return loadedImage
}
