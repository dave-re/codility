package main

// Solution is solution
func Solution(X int, Y int, D int) int {
	val := (Y - X) / D
	rest := (Y - X) % D
	if rest == 0 {
		return val
	}
	return val + 1
}
