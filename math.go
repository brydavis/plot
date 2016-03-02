package main

import "math"

func round(x float64) float64 {
	return math.Floor(x + .5)
}

func base(x float64, base int) int {
	b := float64(base)
	return int(b * round(x/b))
}
