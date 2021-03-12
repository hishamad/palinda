package main

import (
	"testing"
)

// test that ConcurrentSum sums an even-length array correctly
type testStruct struct {
	elements []int
	sum int
}

var tests = []testStruct {
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
	{[]int{-1, -2, -3, -4, 1, 2, 3, 4}, 0},
	{[]int{-10, 10, 0, 5, -5, 3, 2}, 5},
}
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	for _, test := range tests {
		expected := test.sum
		actual := ConcurrentSum(test.elements)
		if actual != expected {
			t.Errorf("expected %d, was %d", expected, actual)
		}
	}


}