package utils

import "math/rand"

func RandomNumInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomFloatInRange(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}
