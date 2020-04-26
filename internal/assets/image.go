package assets

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

func LoadImage(filePath string) *ebiten.Image {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error loading image asset: %v", err)
	}

	decoded, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Error loading image asset: %v", err)
	}

	loadedImage, _ := ebiten.NewImageFromImage(decoded, ebiten.FilterDefault)
	return loadedImage
}
