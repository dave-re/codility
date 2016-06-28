package solution

import "math"

func Solution(A []int) int {
	aLen := len(A)
	if aLen == 2 {
		return 0
	}

	pref := makePrefixSum(A)

	minSum := float64(pref[2]) / float64(2)
	minIdx := 0
	for i := 0; i < aLen-1; i++ {
		sum := float64(pref[i+2]-pref[i]) / float64(2)
		if i+3 <= aLen {
			sum2 := float64(pref[i+3]-pref[i]) / float64(3)
			sum = math.Min(sum, sum2)
		}

		if sum < minSum {
			minSum = sum
			minIdx = i
		}
	}
	return minIdx
}

func makePrefixSum(A []int) []int {
	pref := make([]int, len(A)+1)
	for i := range A {
		pref[i+1] = pref[i] + A[i]
	}
	return pref
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
