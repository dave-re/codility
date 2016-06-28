package solution

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{2, 1, 1, 2, 3, 1},
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
