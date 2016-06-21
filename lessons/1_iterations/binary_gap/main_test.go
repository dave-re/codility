package main

import "testing"

func TestSolution(t *testing.T) {
	testValues := map[int]int{
		9:          2,
		529:        4,
		20:         1,
		15:         0,
		1041:       5,
		2147483647: 0,
	}

	for param, expectVal := range testValues {
		val := Solution(param)
		if expectVal != val {
			t.Errorf("Expected %d, but got %d", expectVal, val)
		}
	}
}
