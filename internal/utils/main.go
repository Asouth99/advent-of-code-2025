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

func Comb(n, m int) [][]int {
	var result [][]int
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				combination := make([]int, len(s))
				copy(combination, s)
				result = append(result, combination)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
	return result
}

func CombWithReplacement(n, m int) [][]int {
	var result [][]int
	s := make([]int, m)
	last := m - 1

	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				combination := make([]int, len(s))
				copy(combination, s)
				result = append(result, combination)
			} else {
				rc(i+1, j)
			}
		}
		return
	}
	rc(0, 0)
	return result
}
