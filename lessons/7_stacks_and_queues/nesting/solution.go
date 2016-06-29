package solution

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}

	bCount := 0
	for i := range S {
		c := S[i]
		if c == '(' {
			bCount++
		} else {
			bCount--
			if bCount < 0 {
				return 0
			}
		}
	}

	if bCount == 0 {
		return 1
	}
	return 0
}
