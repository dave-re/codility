package solution

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{9, 3, 9, 3, 9, 1000000000, 9},
			Expected: 1000000000,
		},
		{
			A:        []int{9, 3, 9, 3, 9, 7, 9},
			Expected: 7,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
