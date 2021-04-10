package main

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	maxPosition := -1

	max := func(pos int) int {
		result := nums[pos]
		for i := pos; i < pos+k; i++ {
			if nums[i] >= result {
				result = nums[i]
				maxPosition = i
			}
		}
		return result
	}

	for i := range result {
		if maxPosition < i {
			result[i] = max(i)
		} else {
			if result[i-1] <= nums[i+k-1] {
				result[i] = nums[i+k-1]
				maxPosition = i + k - 1
			} else {
				result[i] = result[i-1]
			}
		}
	}
	return result
}
