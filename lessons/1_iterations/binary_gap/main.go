package main

// Solution is solution of the iterations lesson
func Solution(N int) int {
	maxGap := 0
	curGap := 0
	startCheck := false

	n := N
	for n > 0 {
		rest := (n % 2)
		n = (n / 2)
		if rest == 1 {
			startCheck = true
			if curGap > maxGap {
				maxGap = curGap
			}
			curGap = 0
		} else {
			if startCheck {
				curGap++
			}
		}
	}
	return maxGap
}
