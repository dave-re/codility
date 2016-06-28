package solution

func Solution(X int, A []int) int {
	count := 0
	cMap := make(map[int]bool, X)
	for i, v := range A {
		if !cMap[v] {
			cMap[v] = true
			count++
			if count == X {
				return i
			}
		}
	}
	return -1
}
