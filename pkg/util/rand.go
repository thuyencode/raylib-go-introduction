package util

import "math/rand/v2"

func RandFloat32(max float32, min float32) float32 {
	return rand.Float32()*(max-min) + min
}

func RandInt(max int, min int) int {
	return (rand.IntN(min) + max) - min
}

func RandItem[S ~[]I, I any](s S) I {
	return s[rand.IntN(len(s))]
}
