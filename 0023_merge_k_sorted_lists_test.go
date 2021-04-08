package main

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		input []*ListNode
		want  *ListNode
	}{
		{[]*ListNode{CreateLinkedList(541), CreateLinkedList(431), CreateLinkedList(62)}, CreateLinkedList(65443211)},
		{make([]*ListNode, 0), nil},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := mergeKLists(testCase.input)
				if (got != nil || testCase.want != nil) && got.String() != testCase.want.String() {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
