package main

func Solution(A []int) int {
	aLen := len(A) + 1
	cArr := make([]bool, aLen, aLen)
	for _, v := range A {
		if 0 < v && v < aLen {
			cArr[v] = true
		}
	}

	for i, v := range cArr[1:] {
		if !v {
			return i + 1
		}
	}
	return aLen
}
