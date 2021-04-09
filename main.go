package main

import (
	"container/list"
	"fmt"
)

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Equals(other *TreeNode) bool {
	if t == nil && other == nil {
		return true
	}
	if other == nil || t.Val != other.Val {
		return false
	}
	result := true
	if t.Left != nil || other.Left != nil {
		if t.Left == nil && other.Left != nil {
			result = false
		} else if t.Left != nil && other.Left == nil {
			result = false
		} else {
			result = result && t.Left.Equals(other.Left)
		}
	}
	if t.Right != nil || other.Right != nil {
		if t.Right == nil && other.Right != nil {
			result = false
		} else if t.Right != nil && other.Right == nil {
			result = false
		} else {
			result = result && t.Right.Equals(other.Right)
		}
	}

	return result
}

// CreateBinaryTree returns the root node of a binary tree created from the input, a negative number in the input represents a nil node
func CreateBinaryTree(numbers []int) *TreeNode {
	if len(numbers) == 0 {
		return nil
	}
	queue := list.New()

	root := &TreeNode{Val: numbers[0]}
	queue.PushBack(root)

	for i := 1; i < len(numbers); i++ {
		e := queue.Front()
		if n, ok := e.Value.(*TreeNode); ok {
			if n.Left == nil {
				n.Left = &TreeNode{Val: numbers[i]}
				if n.Left.Val >= 0 {
					queue.PushBack(n.Left)
				}
			} else {
				n.Right = &TreeNode{Val: numbers[i]}
				if n.Right.Val >= 0 {
					queue.PushBack(n.Right)
				}
				queue.Remove(e)
			}
		}
	}

	queue = list.New()
	queue.PushBack(root)
	for e := queue.Front(); e != nil; e = queue.Front() {
		if n, ok := e.Value.(*TreeNode); ok {
			if n.Left != nil {
				if n.Left.Val == -1 {
					n.Left = nil
				} else {
					queue.PushBack(n.Left)
				}
			}
			if n.Right != nil {
				if n.Right.Val == -1 {
					n.Right = nil
				} else {
					queue.PushBack(n.Right)
				}
			}
		}
		queue.Remove(e)
	}

	return root
}
