package solution

func Solution(A []int) int {
	aLen := len(A)
	if aLen == 0 {
		return -1
	}

	candidate := A[0]
	count := 1
	for _, v := range A[1:] {
		if count > 0 {
			if candidate == v {
				count++
			} else {
				count--
			}
		} else {
			count++
			candidate = v
		}
	}

	if count == 0 {
		return -1
	}

	cCount := 0
	cIdx := -1
	halfCount := (aLen / 2)
	for i, v := range A {
		if v == candidate {
			cCount++
			if cIdx == -1 {
				cIdx = i
			}
			if cCount > halfCount {
				return cIdx
			}
		}
	}

	return -1
}
