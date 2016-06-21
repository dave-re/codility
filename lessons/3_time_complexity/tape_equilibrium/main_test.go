package main

import "testing"

type testValue struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        []int{3, 1, 2, 4, 3},
			Expected: 1,
		},
		{
			A:        []int{1, 1000},
			Expected: 999,
		},
		{
			A:        []int{1, -1000},
			Expected: 1001,
		},
		{
			A:        []int{1000, -1000},
			Expected: 2000,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.A)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
