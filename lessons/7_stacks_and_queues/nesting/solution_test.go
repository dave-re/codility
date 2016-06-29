package solution

import "testing"

type testValue struct {
	S        string
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			S:        "(()(())())",
			Expected: 1,
		},
		{
			S:        "())",
			Expected: 0,
		},
		{
			S:        ")(",
			Expected: 0,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.S)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
