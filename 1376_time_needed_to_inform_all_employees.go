package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	employeesOfManager := make(map[int][]int, n)
	for e, m := range manager {
		if _, ok := employeesOfManager[m]; !ok {
			employeesOfManager[m] = []int{e}
		} else {
			employeesOfManager[m] = append(employeesOfManager[m], e)
		}
	}

	var dfs func(m int) int
	dfs = func(m int) int {
		result := 0
		for _, e := range employeesOfManager[m] {
			tmp := dfs(e)
			if tmp > result {
				result = tmp
			}
		}
		return result + informTime[m]
	}

	return dfs(headID)
}
