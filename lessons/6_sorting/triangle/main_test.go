package main

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{10, 2, 5, 1, 8, 20},
			Expected: 1,
		},
		{
			A:        []int{10, 50, 5, 1},
			Expected: 0,
		},
		{
			A:        []int{1, 2, 3},
			Expected: 0,
		},
		{
			A:        []int{-1, -5, -3},
			Expected: 0,
		},
	}

	for idx, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("%d. Expected %d, but got %d", idx+1, value.Expected, ret)
		}
	}
}
