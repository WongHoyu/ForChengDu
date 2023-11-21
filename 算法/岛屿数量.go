package 算法

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var res int
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return
		}
		if grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				res++
				dfs(r, c)
			}
		}
	}

	return res
}
