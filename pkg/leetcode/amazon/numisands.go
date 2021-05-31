package amazon

func numIslands(grid [][]byte) int {
	islands := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 {
				islands++
				search(grid, &GridNode{x, y})
			}
		}
	}
	return islands
}

func search(grid [][]byte, root *GridNode) {
	queue := make([]*GridNode, 0)
	grid[root.x][root.y] = 0 // mark as visited

	queue = append(queue, root)
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]

		if r.x -1 >= 0 && grid[r.x-1][r.y] != 0 {
			grid[r.x-1][r.y] = 0
			queue = append(queue, &GridNode{r.x-1, r.y})
		}
		if r.x + 1 < len(grid) && grid[r.x+1][r.y] != 0 {
			grid[r.x+1][r.y] = 0
			queue = append(queue, &GridNode{r.x + 1, r.y})
		}
		if r.y - 1 >= 0 && grid[r.x][r.y-1] != 0 {
			grid[r.x][r.y-1] = 0
			queue = append(queue, &GridNode{r.x, r.y-1})
		}
		if r.y + 1 < len(grid[0]) && grid[r.x][r.y+1] != 0 {
			grid[r.x][r.y+1] = 0
			queue = append(queue, &GridNode{r.x, r.y+1})
		}

	}
}

type GridNode struct {
	x int
	y int
}
