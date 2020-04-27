package utils

import "image"

type Collidable interface {
	BoundingBox() image.Rectangle
}

func Collides(objectA Collidable, objectB Collidable) bool {
	return objectA.BoundingBox().Overlaps(objectB.BoundingBox())
}
