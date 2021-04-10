package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		input struct {
			left  *ListNode
			right *ListNode
		}
		want *ListNode
	}{
		{struct {
			left  *ListNode
			right *ListNode
		}{left: CreateLinkedList(1, 2, 4), right: CreateLinkedList(1, 3, 4)}, CreateLinkedList(1, 1, 2, 3, 4, 4)},
		{struct {
			left  *ListNode
			right *ListNode
		}{left: CreateLinkedList(), right: CreateLinkedList()}, CreateLinkedList()},
		{struct {
			left  *ListNode
			right *ListNode
		}{left: CreateLinkedList(), right: CreateLinkedList(0)}, CreateLinkedList(0)},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := mergeTwoLists(testCase.input.left, testCase.input.right)
				if testCase.want == nil {
					assert.Nil(t, got)
				} else {
					assert.Equal(t, testCase.want.String(), got.String())
				}
			},
		)
	}
}
