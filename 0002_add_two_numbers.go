package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var recursiveAdd func(*ListNode, *ListNode, int) *ListNode
	recursiveAdd = func(l1 *ListNode, l2 *ListNode, remainder int) *ListNode {
		result := &ListNode{}

		// create an empty list node when it doesn't exist
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}

		// calculate the value and the remainder
		result.Val += l1.Val + l2.Val + remainder
		remainder = result.Val / 10
		result.Val %= 10

		// if one of the next nodes exist or we have a remainder do a recursive call
		if l1.Next != nil || l2.Next != nil || remainder > 0 {
			result.Next = recursiveAdd(l1.Next, l2.Next, remainder)
		}
		return result
	}

	return recursiveAdd(l1, l2, 0)
}
