package main

func maxResult(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	bestResults := make(map[int]int)
	bestResults[0] = nums[0]

	const minInt int = -1 << 31

	var recursiveResult func(pos int) int
	recursiveResult = func(pos int) int {
		if result, ok := bestResults[pos]; ok {
			return result
		}
		result := minInt
		for i := 1; i <= k && pos-i >= 0; i++ {
			bestResultSoFar := recursiveResult(pos - i)
			if bestResultSoFar > result {
				result = bestResultSoFar
			}
		}
		for i := 0; i < pos-k; i++ {
			delete(bestResults, i)
		}
		bestResults[pos] = result + nums[pos]
		return bestResults[pos]
	}

	return recursiveResult(len(nums) - 1)
}
