package main

func NoOfIsland() int {
	// grid := [][]int{
	// 	{1, 1, 0, 0, 0},
	// 	{0, 1, 0, 0, 1},
	// 	{1, 0, 0, 1, 1},
	// 	{0, 0, 0, 1, 0},
	// 	{1, 0, 1, 0, 1},
	// }
	//  [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	// r, c := len(grid[0]), len(grid)
	// noOfIsland := 0
	// for i := 0; i < c; i++ {
	// 	for j := 0; j < r; j++ {
	// 		if grid[i][j] == 1 {
	// 			noOfIsland++
	// 			NoOfIslandRec(i, j, c, r, grid)
	// 		}
	// 	}
	// }
	// return noOfIsland
	return numIslands(grid)
}

func numIslands(grid [][]byte) int {
    r, c := len(grid[0]), len(grid)
	noOfIsland := 0
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if grid[i][j] == 1 {
				noOfIsland++
				NoOfIslandRec(i, j, c, r, grid)
			}
		}
	}
	return noOfIsland

}


func NoOfIslandRec(i, j, c, r int, grid [][]byte) {
	grid[i][j] = 0
	if IsValidPosition(i+1, j, c, r, grid) {
		NoOfIslandRec(i+1, j, c, r, grid)
	}
	if IsValidPosition(i-1, j, c, r, grid) {
		NoOfIslandRec(i-1, j, c, r, grid)
	}
	if IsValidPosition(i, j-1, c, r, grid) {
		NoOfIslandRec(i, j-1, c, r, grid)
	}
	if IsValidPosition(i, j+1, c, r, grid) {
		NoOfIslandRec(i, j+1, c, r, grid)
	}
}

func IsValidPosition(i, j, c, r int, grid [][]byte) bool {
	return (i >= 0 && j >= 0 && i <= c-1 && j <= r-1 && grid[i][j] == 1)
}
