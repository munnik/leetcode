package main

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		input [][]int
		want  float64
	}{
		{[][]int{{1, 3}, {2}}, 2.0},
		{[][]int{{1, 2}, {3, 4}}, 2.5},
		{[][]int{{0, 0}, {0, 0}}, 0.0},
		{[][]int{{}, {1}}, 1.0},
		{[][]int{{2}, {}}, 2.0},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := findMedianSortedArrays(testCase.input[0], testCase.input[1])
				if got != testCase.want {
					t.Fatalf("Got %v; want %v", got, testCase.want)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
