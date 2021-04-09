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

	var dfs func(m int, time chan int)
	dfs = func(m int, time chan int) {
		result := 0
		channels := make([]chan int, len(employeesOfManager[m]))
		for i, e := range employeesOfManager[m] {
			channels[i] = make(chan int)
			go dfs(e, channels[i])
		}
		for _, c := range channels {
			tmp := <-c
			if tmp > result {
				result = tmp
			}
		}

		time <- result + informTime[m]
	}

	result := make(chan int)
	go dfs(headID, result)

	return <-result
}
