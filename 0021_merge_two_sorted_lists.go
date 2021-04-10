package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	result := &ListNode{}
	current := result

	for l1 != nil || l2 != nil {
		if l1 != nil {
			if l2 == nil || l2.Val > l1.Val {
				current.Val = l1.Val
				l1 = l1.Next
			} else {
				current.Val = l2.Val
				l2 = l2.Next
			}
		} else {
			if l1 == nil || l1.Val > l2.Val {
				current.Val = l2.Val
				l2 = l2.Next
			} else {
				current.Val = l1.Val
				l1 = l1.Next
			}
		}
		if l1 != nil || l2 != nil {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return result
}
