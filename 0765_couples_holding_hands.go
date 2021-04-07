package main

func minSwapsCouples(row []int) int {
	result := 0
	positionMap := make(map[int]int, len(row))

	for position, number := range row {
		positionMap[number] = position
	}

	for i := 0; i < len(row); i += 2 {
		if row[i]%2 == 0 && row[i]+1 != row[i+1] {
			row[positionMap[row[i]+1]] = row[i+1]
			positionMap[row[i+1]] = positionMap[row[i]+1]
			row[i+1] = row[i] + 1
			positionMap[row[i+1]] = i + 1
			result++
		} else if row[i]%2 == 1 && row[i]-1 != row[i+1] {
			row[positionMap[row[i]-1]] = row[i+1]
			positionMap[row[i+1]] = positionMap[row[i]-1]
			row[i+1] = row[i] - 1
			positionMap[row[i+1]] = i + 1
			result++
		}
	}

	return result
}
