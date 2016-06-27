package main

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{-3, 1, 2, -2, 5, 6},
			Expected: 60,
		},
		{
			A:        []int{-5, 5, -5, 4},
			Expected: 125,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
