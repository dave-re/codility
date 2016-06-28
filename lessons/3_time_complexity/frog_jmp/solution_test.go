package solution

import "testing"

type testValue struct {
	X        int
	Y        int
	D        int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			X:        10,
			Y:        85,
			D:        30,
			Expected: 3,
		},
		{
			X:        5,
			Y:        80,
			D:        25,
			Expected: 3,
		},
		{
			X:        1,
			Y:        1000000000,
			D:        1,
			Expected: 999999999,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.X, value.Y, value.D)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
