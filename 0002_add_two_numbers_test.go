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
		{[]*ListNode{CreateLinkedList(342), CreateLinkedList(465)}, CreateLinkedList(807)},
		{[]*ListNode{CreateLinkedList(0), CreateLinkedList(0)}, CreateLinkedList(0)},
		{[]*ListNode{CreateLinkedList(9999999), CreateLinkedList(9999)}, CreateLinkedList(10009998)},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := addTwoNumbers(testCase.input[0], testCase.input[1])
				if (got != nil || testCase.want != nil) && got.String() != testCase.want.String() {
					t.Fatalf("Got %v; want %v", got, testCase.want)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
