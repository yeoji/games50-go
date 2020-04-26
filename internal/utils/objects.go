package utils

import "image"

type Collidable interface {
	BoundingBox() image.Rectangle
}

func Collides(objectA Collidable, objectB Collidable) bool {
	// boxA := objectA.BoundingBox()
	// boxB := objectB.BoundingBox()

	// if boxA.Min.X > boxB.Max.X || boxB.Min.X > boxA.Max.X {
	// 	return false
	// }

	// if boxA.Min.Y > boxB.Max.Y || boxB.Min.Y > boxA.Max.Y {
	// 	return false
	// }

	return objectA.BoundingBox().Overlaps(objectB.BoundingBox())
}
