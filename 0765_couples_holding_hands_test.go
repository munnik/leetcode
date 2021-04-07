package main

import (
	"fmt"
	"testing"
)

func TestMinSwapsCouples(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0, 2, 1, 3}, 1},
		{[]int{3, 2, 0, 1}, 0},
		{[]int{0, 2, 4, 6, 7, 1, 3, 5}, 3},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := minSwapsCouples(testCase.input)
				if got != testCase.want {
					t.Fatalf("Got %v; want %v", got, testCase.want)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
