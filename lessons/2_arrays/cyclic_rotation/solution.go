package solution

// Solution is solution
func Solution(A []int, K int) []int {
	aLen := len(A)
	retA := make([]int, aLen, aLen)

	for index, item := range A {
		retA[((index + K) % aLen)] = item
	}
	return retA
}
