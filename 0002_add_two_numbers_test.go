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
		{[]*ListNode{CreateLinkedList(2, 4, 3), CreateLinkedList(5, 6, 4)}, CreateLinkedList(7, 0, 8)},
		{[]*ListNode{CreateLinkedList(0), CreateLinkedList(0)}, CreateLinkedList(0)},
		{[]*ListNode{CreateLinkedList(9, 9, 9, 9, 9, 9, 9), CreateLinkedList(9, 9, 9, 9)}, CreateLinkedList(8, 9, 9, 9, 0, 0, 0, 1)},
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
