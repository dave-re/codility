package solution

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{4, 2, 2, 5, 1, 5, 8},
			Expected: 1,
		},
		{
			A:        []int{4, 3, 4, 2, 1, 5},
			Expected: 3,
		},
		{
			A:        []int{-3, -5, -8, -4, -10},
			Expected: 2,
		},
		{
			A:        []int{10, 10, -1, 2, 4, -1, 2, -1},
			Expected: 5,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
