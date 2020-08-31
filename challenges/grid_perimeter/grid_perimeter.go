// You are given a 2D matrix of 1s and 0s where 1
// represents land and 0 represents water.

// Grid cells are connected horizontally or vertically
// (not diagonally). The grid is completely surrounded
// by water, and there is exactly one island.

// An island is a group is cells connected horizontally
// or vertically, but not diagonally. There is guaranteed
// to be exactly one island in this grid, and the island
// doesn't have water inside that isn't connected to the
// water around the island. Each cell has a side length of 1.

// Determine the perimeter of this island.

// For example, given the following matrix:

// [[0, 1, 1, 0],
// [1, 1, 1, 0],
// [0, 1, 1, 0],
// [0, 0, 1, 0]]

// Return 14.

package main

import (
	"fmt"
)

// Grid is a 2D matrix
type Grid [][]int

func (grid Grid) isLand(x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	return grid[x][y]
}

// GetPerimeter returns the perimeter of the island
// within the grid
func GetPerimeter(grid Grid) int {
	perimeter := 0
	for i := range grid {
		for j := range grid[i] {
			if grid.isLand(i, j) == 0 {
				continue
			}
			// is land
			perimeter += 4
			perimeter -= 2 * grid.isLand(i-1, j)
			perimeter -= 2 * grid.isLand(i, j-1)
		}
	}
	return perimeter
}

func main() {
	grid := Grid{
		{0, 1, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 1, 1, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0},
	}

	fmt.Println("Perimeter of the island:", GetPerimeter(grid))
}
