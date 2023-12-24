package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gammazero/deque"
)

type Vector struct {
	r  int
	c  int
	dr int
	dc int
}

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func part1(grid []string) {
	start := Vector{0, -1, 0, 1} // Start at [0, -1] going right
	visited := walk(grid, []Vector{start})

	// Only count [r, c]. Ignore directions and the starting cell.
	cells := make(map[[2]int]bool)
	for vec := range visited {
		cells[[2]int{vec.r, vec.c}] = true
	}
	fmt.Println(len(cells) - 1)
}

func part2(grid []string) {
	numRows := len(grid)
	numCols := len(grid[0])

	startingVecs := make([]Vector, 0)
	for c := range grid[0] {
		startingVecs = append(startingVecs, Vector{-1, c, 1, 0})       // Top row, down direction
		startingVecs = append(startingVecs, Vector{numRows, c, -1, 0}) // Bot row, up direction
	}

	for r := range grid {
		startingVecs = append(startingVecs, Vector{r, -1, 0, 1})
		startingVecs = append(startingVecs, Vector{r, numCols, 0, -1})
	}

	maxTiles := 0
	for _, startingVec := range startingVecs {
		visited := walk(grid, []Vector{startingVec})

		// Only count [r, c]. Ignore directions and the starting cell.
		cells := make(map[[2]int]bool)
		for vec := range visited {
			cells[[2]int{vec.r, vec.c}] = true
		}
		maxTiles = max(maxTiles, len(cells)-1)
	}
	fmt.Println(maxTiles)
}

func walk(grid []string, startVecs []Vector) map[Vector]bool {
	numRows := len(grid)
	numCols := len(grid[0])

	visited := make(map[Vector]bool)
	stack := deque.New[Vector]()

	for _, startVec := range startVecs {
		stack.PushBack(startVec)

		// DFS
		// 1. Move the vector to the next point.
		// 2. Then check bounds.
		// 3. Then check cell symbol.
		// 4. Then add the new vector to the stack with the correct displacement (dr/dc).
		for stack.Len() > 0 {
			vec := stack.PopBack()
			if !visited[vec] {
				visited[vec] = true

				// Look at the vector's next point
				r, c, dr, dc := vec.r, vec.c, vec.dr, vec.dc
				r += dr // Next r
				c += dc // Next c

				if !isValidPosition(r, c, numRows, numCols) {
					continue
				}

				if grid[r][c] == '.' || (grid[r][c] == '-' && dc != 0) || (grid[r][c] == '|' && dr != 0) { // Going same direction
					stack.PushBack(Vector{r, c, dr, dc})

				} else if grid[r][c] == '|' && dc != 0 { // Going horizontal and hit vertical bar
					stack.PushBack(Vector{r, c, -1, 0})
					stack.PushBack(Vector{r, c, 1, 0})

				} else if grid[r][c] == '-' && dr != 0 { // Going vertical and hit horizontal bar
					stack.PushBack(Vector{r, c, 0, -1})
					stack.PushBack(Vector{r, c, 0, 1})

				} else if grid[r][c] == '/' { // Col dir becomes the new inverted row dir. Row dir becomes the new interved col dir
					stack.PushBack(Vector{r, c, -dc, -dr})

				} else if grid[r][c] == '\\' { // Col dir becomes the new row dir. Row dir becomes the new col dir.
					stack.PushBack(Vector{r, c, dc, dr})

				} else {
					log.Printf("Hit an unhandled edge case: %v\n", grid[r][c])
				}
			}
		}
	}
	return visited
}

func isValidPosition(row, col, numRows, numCols int) bool {
	return row >= 0 && row < numRows && col >= 0 && col < numCols
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInput(fname string) []string {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(string(content), "\n")
}
