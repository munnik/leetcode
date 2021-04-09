package main

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		input *ListNode
		want  bool
	}{
		{CreateLinkedList(1111).AddCycle(1), true},
		{CreateLinkedList(12).AddCycle(0), true},
		{CreateLinkedList(1).AddCycle(-1), false},
		{nil, false},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := hasCycle(testCase.input)
				if got != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
