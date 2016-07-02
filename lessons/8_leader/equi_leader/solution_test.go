package solution

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{4, 3, 4, 4, 4, 2},
			Expected: 2,
		},
		{
			A:        []int{4, 4, 2, 5, 3, 4, 4, 4},
			Expected: 3,
		},
		{
			A:        []int{0, 0},
			Expected: 1,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
