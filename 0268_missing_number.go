package main

import (
	"errors"
	"sort"
)

func missingNumber(nums []int) int {
	var missingNumberRecursive func([]int, int, int) (int, error)
	missingNumberRecursive = func(nums []int, start int, end int) (int, error) {
		expectedLength := end - start + 1
		left := nums[:expectedLength/2]
		right := nums[expectedLength/2:]

		if nums[0] != start {
			return start, nil
		}
		if nums[len(nums)-1] != end {
			return end, nil
		}
		if left[len(left)-1]+1 != right[0] {
			return left[len(left)-1] + 1, nil
		}
		if left[0]+len(left)-1 != left[len(left)-1] {
			return missingNumberRecursive(left, left[0], left[len(left)-1])
		}
		if right[0]+len(right)-1 != right[len(right)-1] {
			return missingNumberRecursive(right, right[0], right[len(right)-1])
		}

		return 0, errors.New("could not find missing number")
	}

	sort.Ints(nums)
	result, _ := missingNumberRecursive(nums, 0, len(nums))
	return result
}
