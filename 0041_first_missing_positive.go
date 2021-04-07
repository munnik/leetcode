package main

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)

	result := 1
	previous := 0

	for _, num := range nums {
		if num < 1 {
			continue
		}
		if num == previous {
			continue
		}
		if num == result {
			previous = num
			result++
			continue
		}
		break
	}

	return result
}
