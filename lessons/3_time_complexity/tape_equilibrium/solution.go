package solution

// Solution is solution
func Solution(A []int) int {
	total := 0
	for _, val := range A {
		total += val
	}

	arr := A[:len(A)-1]
	sum := 0
	diff := 2000
	for _, val := range arr {
		sum += val
		curDiff := abs(sum - (total - sum))
		if curDiff < diff {
			diff = curDiff
		}
	}
	return diff
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
