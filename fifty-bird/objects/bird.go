package objects

import (
	"games50-go/internal/assets"
	"games50-go/internal/utils"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const Gravity = 20
const JumpHeight = -5

type Bird struct {
	image *ebiten.Image
	x     float64
	y     float64
	dy    float64 // vertical velocity
}

func NewBird(screen *ebiten.Image) Bird {
	screenWidth, screenHeight := screen.Size()
	birdImage := assets.LoadImage("assets/art/bird.png")
	birdWidth, birdHeight := birdImage.Size()

	return Bird{
		image: birdImage,
		x:     float64(screenWidth/2 - birdWidth/2),
		y:     float64(screenHeight/2 - birdHeight/2),
		dy:    0,
	}
}

func (b *Bird) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.dy = JumpHeight
	} else {
		b.dy += Gravity * 1 / ebiten.CurrentTPS()
	}
	b.y += b.dy
}

func (b *Bird) Render(screen *ebiten.Image) {
	birdOptions := &ebiten.DrawImageOptions{}
	birdOptions.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, birdOptions)
}

func (b *Bird) HasHitTheGround(groundY int) bool {
	return int(b.y) > groundY
}

func (b *Bird) HasHitAnyPipes(pipePairs []*PipePair) bool {
	for _, pipePair := range pipePairs {
		if utils.Collides(b, &pipePair.Bottom) || utils.Collides(b, &pipePair.Top) {
			return true
		}
	}
	return false
}

func (b *Bird) BoundingBox() image.Rectangle {
	birdWidth, birdHeight := b.image.Size()

	return image.Rect(int(b.x), int(b.y), int(b.x)+birdWidth, int(b.y)+birdHeight)
}
