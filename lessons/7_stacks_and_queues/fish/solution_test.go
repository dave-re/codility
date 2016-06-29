package solution

import "testing"

type testValue struct {
	A        []int
	B        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{4, 3, 2, 1, 5},
			B:        []int{0, 1, 0, 0, 0},
			Expected: 2,
		},
		{
			A:        []int{4, 3, 2, 1, 5},
			B:        []int{0, 1, 0, 1, 0},
			Expected: 2,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.A, value.B)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
