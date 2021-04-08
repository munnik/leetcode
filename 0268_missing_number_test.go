package main

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := missingNumber(testCase.input)
				if got != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
