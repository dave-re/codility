package solution

func Solution(A int, B int, K int) int {
	if A == 0 && B == 0 {
		return 1
	}

	a := ((A - 1) / K) + 1
	if A == 0 {
		a = 0
	}
	b := (B / K) + 1
	return b - a
}
