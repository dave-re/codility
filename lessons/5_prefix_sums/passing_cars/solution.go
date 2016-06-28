package solution

func Solution(A []int) int {
	sum := 0
	x := 0
	for _, v := range A {
		if v == 0 {
			x++
		} else {
			sum += x
			if sum > 1000000000 {
				return -1
			}
		}
	}
	return sum
}
