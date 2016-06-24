package main

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{0, 1, 0, 1, 1},
			Expected: 5,
		},
		{
			A:        []int{0, 1},
			Expected: 1,
		},
		{
			A:        []int{0, 1, 0, 1, 1, 0, 1},
			Expected: 8,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
