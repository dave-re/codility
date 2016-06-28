package solution

func Solution(A []int) int {
	aLen := len(A)
	if aLen < 3 {
		return 0
	}
	// sort.IntSlice(A).Sort()
	A = mergeSort(A)

	for i := range A {
		if i+2 < aLen {
			if A[i]+A[i+1] > A[i+2] {
				return 1
			}
		}
	}
	return 0
}

func mergeSort(A []int) []int {
	aLen := len(A)
	if aLen == 1 {
		return A
	}

	mid := aLen / 2
	lArr := mergeSort(A[:mid])
	rArr := mergeSort(A[mid:])

	return merge(lArr, rArr)
}

func merge(left []int, right []int) []int {
	lLen, rLen := len(left), len(right)
	l, r := 0, 0
	mArr := make([]int, 0, lLen+rLen)

	for l < lLen && r < rLen {
		if left[l] < right[r] {
			mArr = append(mArr, left[l])
			l++
		} else {
			mArr = append(mArr, right[r])
			r++
		}
	}
	mArr = append(mArr, left[l:]...)
	mArr = append(mArr, right[r:]...)
	return mArr
}
