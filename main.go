package main

import "fmt"

/*
Example 1:

Input: grid = [

	["1","1","1","1","0"],
	["1","1","0","1","0"],
	["1","1","0","0","0"],
	["0","0","0","0","0"]

]

Output: 1
*/
type pair struct {
	x, y int
}

var visited = make(map[pair]bool)
var grid = [][]byte{
	{'1', '1', '0', '0', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '1', '0', '0'},
	{'0', '0', '0', '1', '1'},
}

func main() {

	fmt.Println(searchNumberIsland(grid))

}

func dfs(r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}

	_, ok := visited[pair{r, c}]
	if ok {
		return
	}

	visited[pair{r, c}] = true

	dfs(r+1, c)
	dfs(r-1, c)
	dfs(r, c+1)
	dfs(r, c-1)

}

func searchNumberIsland(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	island := 0
outer:
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				_, ok := visited[pair{r, c}]
				if !ok {
					dfs(r, c)
					island += 1
				}
				continue outer
			}
		}
	}

	return island
}
