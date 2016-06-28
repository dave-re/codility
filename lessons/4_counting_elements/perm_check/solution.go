package solution

func Solution(A []int) int {
	aLen := len(A)
	m := make(map[int]bool, aLen)
	sum := 0
	for _, v := range A {
		if m[v] {
			return 0
		}
		m[v] = true
		sum += v
	}

	total := (aLen * (aLen + 1)) / 2
	if sum == total {
		return 1
	}
	return 0
}
