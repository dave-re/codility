package solution

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{1, 5, 2, 1, 4, 0},
			Expected: 11,
		},
		{
			A:        []int{1, 1, 1},
			Expected: 3,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
