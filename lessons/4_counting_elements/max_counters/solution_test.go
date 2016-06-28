package solution

import (
	"reflect"
	"testing"
)

type testValue struct {
	N        int
	A        []int
	Expected []int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			N:        5,
			A:        []int{3, 4, 4, 6, 1, 4, 4},
			Expected: []int{3, 2, 2, 4, 2},
		},
	}

	for _, value := range testValues {
		arr := Solution(value.N, value.A)
		if !reflect.DeepEqual(arr, value.Expected) {
			t.Errorf("Expected %#v, but got %#v", value.Expected, arr)
		}
	}
}
