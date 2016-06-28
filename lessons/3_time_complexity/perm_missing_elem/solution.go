package solution

func Solution(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}

	n := len(A) + 1
	total := (n * (n + 1)) / 2

	return total - sum
}
