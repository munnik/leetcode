package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	if l.Next != nil {
		return fmt.Sprintf("%d > %s", l.Val, l.Next.String())
	}
	return fmt.Sprintf("%d", l.Val)
}

func CreateLinkedList(number int) *ListNode {
	result := &ListNode{}
	result.Val = number % 10
	number /= 10
	if number > 0 {
		result.Next = CreateLinkedList(number)
	}

	return result
}
