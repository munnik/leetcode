package main

import "container/list"

func isSymmetric(root *TreeNode) bool {
	leftQueue := list.New()
	rightQueue := list.New()
	var tmpQueue *list.List

	tmpQueue = list.New()
	tmpQueue.PushBack(root.Left)
	for tmpQueue.Len() > 0 {
		e := tmpQueue.Front()
		leftQueue.PushBack(e.Value)
		if n, ok := e.Value.(*TreeNode); ok {
			if n != nil {
				tmpQueue.PushBack(n.Left)
				tmpQueue.PushBack(n.Right)
			}
		}
		tmpQueue.Remove(e)
	}

	tmpQueue = list.New()
	tmpQueue.PushBack(root.Right)
	for tmpQueue.Len() > 0 {
		e := tmpQueue.Front()
		rightQueue.PushBack(e.Value)
		if n, ok := e.Value.(*TreeNode); ok {
			if n != nil {
				tmpQueue.PushBack(n.Right)
				tmpQueue.PushBack(n.Left)
			}
		}
		tmpQueue.Remove(e)
	}

	if leftQueue.Len() != rightQueue.Len() {
		return false
	}

	for leftQueue.Len() > 0 {
		le := leftQueue.Front()
		re := rightQueue.Front()
		if ln, ok := le.Value.(*TreeNode); ok {
			if rn, ok := re.Value.(*TreeNode); ok {
				if ln == nil && rn == nil {
				} else if ln == nil || rn == nil || ln.Val != rn.Val {
					return false
				}
				leftQueue.Remove(le)
				rightQueue.Remove(re)
			}
		}
	}

	return true
}
