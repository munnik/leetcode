package main

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	var current *ListNode
	var smallest *ListNode
	foundIndex := 0

	for foundIndex >= 0 {
		foundIndex = -1
		smallest = nil
		for index, l := range lists {
			if l != nil && (smallest == nil || smallest.Val > l.Val) {
				smallest = l
				foundIndex = index
			}
		}
		if foundIndex >= 0 {
			if result == nil {
				result = &ListNode{}
				current = result
			} else {
				current.Next = &ListNode{}
				current = current.Next
			}
			current.Val = smallest.Val
			lists[foundIndex] = lists[foundIndex].Next
		}
	}
	return result
}
