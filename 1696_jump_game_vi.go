package main

import "container/list"

func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	queue := list.New()
	type best struct {
		Value    int
		Position int
	}
	queue.PushBack(best{Value: nums[0], Position: 0})

	for i := 1; i < len(nums); i++ {
		e := queue.Front()
		if b, ok := e.Value.(best); ok {
			dp[i] = nums[i] + b.Value
		}

		for {
			if queue.Len() == 0 {
				break
			}
			e = queue.Back()
			if b, ok := e.Value.(best); ok {
				if b.Value >= dp[i] {
					break
				}
				queue.Remove(e)
			}
		}
		queue.PushBack(best{Value: dp[i], Position: i})

		e = queue.Front()
		if b, ok := e.Value.(best); ok && b.Position == i-k {
			queue.Remove(e)
		}
	}

	return dp[len(dp)-1]
}
