package solution

func Solution(N int, A []int) []int {
	lastMax := 0
	curMax := 0
	arr := make([]int, N, N)

	for _, v := range A {
		if 1 <= v && v <= N {
			if lastMax <= arr[v-1] {
				arr[v-1]++
			} else {
				arr[v-1] = lastMax + 1
			}
			if curMax < arr[v-1] {
				curMax = arr[v-1]
			}
		} else {
			lastMax = curMax
		}
	}

	for i, v := range arr {
		if v < lastMax {
			arr[i] = lastMax
		}
	}
	return arr
}
