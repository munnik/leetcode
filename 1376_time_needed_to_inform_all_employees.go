package main

import "container/list"

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	timeNeededToInform := make(map[int]int, n)
	timeNeededToInform[headID] = 0
	result := 0

	whoNeedToInform := list.New()
	whoNeedToInform.PushBack(headID)

	for len(timeNeededToInform) < n {
		e := whoNeedToInform.Front()
		if headID, ok := e.Value.(int); ok {
			for i, m := range manager {
				if headID == m {
					whoNeedToInform.PushBack(i)
					timeNeededToInform[i] = timeNeededToInform[headID] + informTime[headID]
					if timeNeededToInform[i] > result {
						result = timeNeededToInform[i]
					}
				}
			}
		}
		whoNeedToInform.Remove(e)
	}
	return result
}
