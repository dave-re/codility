package solution

import "testing"

type testValue struct {
	X        int
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			X:        5,
			A:        []int{1, 3, 1, 4, 2, 3, 5, 4},
			Expected: 6,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.X, value.A)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
