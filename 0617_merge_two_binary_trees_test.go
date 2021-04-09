package main

import (
	"fmt"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	tests := []struct {
		input []*TreeNode
		want  *TreeNode
	}{
		{[]*TreeNode{CreateBinaryTree([]int{1}), CreateBinaryTree([]int{1, 2})}, CreateBinaryTree([]int{2, 2})},
		{[]*TreeNode{CreateBinaryTree([]int{1, 3, 2, 5}), CreateBinaryTree([]int{2, 1, 3, -1, 4, -1, 7})}, CreateBinaryTree([]int{3, 4, 5, 5, 4, -1, 7})},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := mergeTrees(testCase.input[0], testCase.input[1])
				if got == nil && testCase.want != nil || got != nil && testCase.want == nil || !got.Equals(testCase.want) {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
