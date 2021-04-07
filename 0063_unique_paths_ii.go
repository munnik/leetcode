package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// slice that holds the number of unique paths from each cell
	// the number of unique paths is 0 if the cell contains an obstacle
	// otherwise it is the sum of the number of unique paths of the cell
	// right and below to it
	uniquePaths := make([][]int, len(obstacleGrid))
	for i := range obstacleGrid {
		uniquePaths[i] = make([]int, len(obstacleGrid[0]))
	}

	width := len(obstacleGrid)
	height := len(obstacleGrid[0])

	// if the finish cell in the obtacle grid contains an obstacle, no
	// unique paths are possible and we can return immediatle
	if obstacleGrid[width-1][height-1] == 1 {
		return 0
	}
	// initialize the number of unique paths from the finish cell to the finish cell
	uniquePaths[width-1][height-1] = 1

	// scan all cells and sum the unique paths
	for i := width - 1; i >= 0; i-- {
		for j := height - 1; j >= 0; j-- {
			// if an obstacle is found, leave the number of unique paths to zero
			if obstacleGrid[i][j] == 1 {
				continue
			}

			right := i + 1
			below := j + 1

			if right < width {
				uniquePaths[i][j] += uniquePaths[right][j]
			}
			if below < height {
				uniquePaths[i][j] += uniquePaths[i][below]
			}
		}
	}

	// return the number of unique paths at the origin
	return uniquePaths[0][0]
}
