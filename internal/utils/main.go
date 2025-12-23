package internal

import (
	"math"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Distance(v1, v2 []float64) float64 {
	var distance float64
	for i := range v1 {
		distance += math.Pow(v1[i]-v2[i], 2)
	}
	return math.Sqrt(distance) //, nil
}
