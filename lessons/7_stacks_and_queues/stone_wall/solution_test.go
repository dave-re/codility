package solution

import "testing"

type testValue struct {
	H        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			H:        []int{8, 8, 5, 7, 9, 8, 7, 4, 8},
			Expected: 7,
		},
		{
			H:        []int{9, 7, 4, 3, 4, 7},
			Expected: 6,
		},
		{
			H:        []int{3, 4, 5, 4, 4},
			Expected: 3,
		},
		{
			H:        []int{8, 5, 5, 8},
			Expected: 3,
		},
		{
			H:        []int{8, 5, 8},
			Expected: 3,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.H)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
