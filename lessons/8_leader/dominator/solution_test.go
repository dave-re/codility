package solution

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{3, 4, 3, 2, 3, -1, 3, 3},
			Expected: 0,
		},
		{
			A:        []int{3, 4, 3, 2, 3, -1},
			Expected: -1,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
