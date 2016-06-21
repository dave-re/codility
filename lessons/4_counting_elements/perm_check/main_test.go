package main

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{4, 1, 3, 2},
			Expected: 1,
		},
		{
			A:        []int{4, 1, 3},
			Expected: 0,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
