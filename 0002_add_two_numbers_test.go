package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		input []*ListNode
		want  *ListNode
	}{
		{[]*ListNode{CreateLinkedList([]int{2, 4, 3}), CreateLinkedList([]int{5, 6, 4})}, CreateLinkedList([]int{7, 0, 8})},
		{[]*ListNode{CreateLinkedList([]int{0}), CreateLinkedList([]int{0})}, CreateLinkedList([]int{0})},
		{[]*ListNode{CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9}), CreateLinkedList([]int{9, 9, 9, 9})}, CreateLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := addTwoNumbers(testCase.input[0], testCase.input[1])
				if (got != nil || testCase.want != nil) && got.String() != testCase.want.String() {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
