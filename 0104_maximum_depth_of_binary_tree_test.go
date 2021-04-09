package main

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		input *TreeNode
		want  int
	}{
		{CreateBinaryTree([]int{3, 9, 20, -1, -1, 15, 7}), 3},
		{CreateBinaryTree([]int{1, -1, 2}), 2},
		{CreateBinaryTree([]int{}), 0},
		{CreateBinaryTree([]int{0}), 1},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := maxDepth(testCase.input)
				if got != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
