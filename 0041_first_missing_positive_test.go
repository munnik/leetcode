package main

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := firstMissingPositive(testCase.input)
				if got != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
