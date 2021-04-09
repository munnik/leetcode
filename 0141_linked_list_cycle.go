package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	visited := make(map[*ListNode]interface{})
	for {
		if _, ok := visited[head]; ok {
			return true
		}
		visited[head] = nil
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	return false
}
