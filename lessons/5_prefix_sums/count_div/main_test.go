package main

import "testing"

type testValue struct {
	A        int
	B        int
	K        int
	Expected int
}

func TestSolution(t *testing.T) {
	testValues := []testValue{
		{
			A:        6,
			B:        11,
			K:        2,
			Expected: 3,
		},
		{
			A:        6,
			B:        6,
			K:        2,
			Expected: 1,
		},
		{
			A:        5,
			B:        5,
			K:        2,
			Expected: 0,
		},
		{
			A:        3,
			B:        10,
			K:        3,
			Expected: 3,
		},
		{
			A:        0,
			B:        5,
			K:        2,
			Expected: 3,
		},
		{
			A:        0,
			B:        0,
			K:        11,
			Expected: 1,
		},
	}

	for _, value := range testValues {
		ret := Solution(value.A, value.B, value.K)
		if ret != value.Expected {
			t.Errorf("Expected %d, but got %d", value.Expected, ret)
		}
	}
}
