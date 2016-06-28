package solution

import "sort"

func Solution(A []int) int {
	sort.IntSlice(A).Sort()
	aLen := len(A)

	exPos := false
	for i := range A {
		if A[i] > 0 {
			exPos = true
		}
	}

	max := A[aLen-3] * A[aLen-2] * A[aLen-1]
	if A[0] < 0 && A[1] < 0 && exPos {
		val := A[0] * A[1] * A[aLen-1]
		if val > max {
			max = val
		}
	}
	return max
}
