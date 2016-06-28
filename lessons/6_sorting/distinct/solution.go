package solution

import "sort"

func Solution(A []int) int {
	aLen := len(A)
	if aLen == 0 {
		return 0
	}

	sort.IntSlice(A).Sort()

	count := 1
	pv := A[0]
	for _, v := range A[1:] {
		if v != pv {
			pv = v
			count++
		}
	}
	return count
}
