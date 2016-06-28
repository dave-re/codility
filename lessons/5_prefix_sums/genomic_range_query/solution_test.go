package solution

import (
	"reflect"
	"testing"
)

type testValue struct {
	S        string
	P        []int
	Q        []int
	Expected []int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			S:        "CAGCCTA",
			P:        []int{2, 5, 0},
			Q:        []int{4, 5, 6},
			Expected: []int{2, 4, 1},
		},
		{
			S:        "CAGCCTA",
			P:        []int{0, 1, 2, 3},
			Q:        []int{0, 5, 4, 3},
			Expected: []int{2, 1, 2, 2},
		},
		{
			S:        "AC",
			P:        []int{0, 0, 1},
			Q:        []int{0, 1, 1},
			Expected: []int{1, 1, 2},
		},
	}

	for _, value := range testValues {
		ret := Solution(value.S, value.P, value.Q)
		if !reflect.DeepEqual(ret, value.Expected) {
			t.Errorf("Expected %#v, but got %#v", value.Expected, ret)
		}
	}
}
