package main

import (
	"fmt"
	"testing"
)

func TestListNodeString(t *testing.T) {
	tests := []struct {
		input *ListNode
		want  string
	}{
		{&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, "2 > 4 > 3 > [nil]"},
		{&ListNode{}, "0 > [nil]"},
		{&ListNode{Val: 0}, "0 > [nil]"},
		{&ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}}, "8 > 9 > 9 > 9 > 0 > 0 > 0 > 1 > [nil]"},
		{(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}).AddCycle(-1), "2 > 4 > 3 > [nil]"},
		{(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}).AddCycle(0), "2 > 4 > 3 > [cycle to position 0]"},
		{(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}).AddCycle(1), "2 > 4 > 3 > [cycle to position 1]"},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := testCase.input
				if got.String() != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}

func TestCreateLinkedList(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{[]int{2, 4, 3}, "2 > 4 > 3 > [nil]"},
		{[]int{0}, "0 > [nil]"},
		{[]int{8, 9, 9, 9, 0, 0, 0, 1}, "8 > 9 > 9 > 9 > 0 > 0 > 0 > 1 > [nil]"},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := CreateLinkedList(testCase.input)
				if got.String() != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}

func TestTreeNodeEquals(t *testing.T) {
	tests := []struct {
		input []*TreeNode
		want  bool
	}{
		{
			[]*TreeNode{
				{Val: 0},
				{Val: 1},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1},
				{Val: 1},
			},
			true,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}},
				{Val: 1},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1},
				{Val: 1, Left: &TreeNode{Val: 2}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}},
				{Val: 1, Left: &TreeNode{Val: 1}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}},
				{Val: 1, Left: &TreeNode{Val: 2}},
			},
			true,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				{Val: 1, Left: &TreeNode{Val: 2}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 0}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			},
			true,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
			},
			true,
		},
		{
			[]*TreeNode{
				{Val: 0, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
			},
			false,
		},
		{
			[]*TreeNode{
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
				{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 4}}},
			},
			false,
		},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := testCase.input[0].Equals(testCase.input[1])
				if got != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}

}

func TestCreateBinaryTree(t *testing.T) {
	tests := []struct {
		input []int
		want  *TreeNode
	}{
		{
			[]int{},
			nil,
		},
		{
			[]int{0},
			&TreeNode{Val: 0},
		},
		{
			[]int{0, 1},
			&TreeNode{Val: 0, Left: &TreeNode{Val: 1}},
		},
		{
			[]int{0, 1, 2},
			&TreeNode{Val: 0, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		},
		{
			[]int{0, 1, 2, 3},
			&TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}},
		},
		{
			[]int{0, 1, 2, -1, 3},
			&TreeNode{Val: 0, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}},
		},
		{
			[]int{1, 3, 4, 7, -1, 8, 2, 4, 3, 9, -1, 1, 7},
			&TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 7}}}},
		},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := CreateBinaryTree(testCase.input)
				if got == nil && testCase.want != nil || got != nil && testCase.want == nil || !got.Equals(testCase.want) {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
