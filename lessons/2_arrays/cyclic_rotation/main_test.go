package main

import (
	"reflect"
	"testing"
)

type testValue struct {
	A        []int
	K        int
	Expected []int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{3, 8, 9, 7, 6},
			K:        3,
			Expected: []int{9, 7, 6, 3, 8},
		},
	}

	for _, value := range testValues {
		arr := Solution(value.A, value.K)
		if !reflect.DeepEqual(arr, value.Expected) {
			t.Errorf("Expected %#v, but got %#v", value.Expected, arr)
		}
	}
}
