package main

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		input *TreeNode
		want  bool
	}{
		{CreateBinaryTree([]int{1}), true},
		{CreateBinaryTree([]int{1, 2}), false},
		{CreateBinaryTree([]int{1, 2, 2}), true},
		{CreateBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}), true},
		{CreateBinaryTree([]int{1, 2, 2, 3, 4, 5, 3}), false},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := isSymmetric(testCase.input)
				if got != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
