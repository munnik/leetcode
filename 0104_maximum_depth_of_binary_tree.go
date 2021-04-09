package main

func maxDepth(root *TreeNode) int {
	max := func(l int, r int) int {
		if l > r {
			return l
		}
		return r
	}

	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}
