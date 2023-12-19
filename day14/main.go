package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func part1(grid [][]byte) {
	grid = shiftRocksDirection(grid, 'N')

	n := len(grid)
	total := 0
	for i, row := range grid {
		factor := n - i
		count := strings.Count(string(row), "O")
		total += factor * count
	}
	fmt.Println(total)
}

func part2(grid [][]byte) {
	// Don't run for 1 billion cycles... it just repeats
	for i := 0; i < 1_000; i++ {
		grid = shiftRocksDirection(grid, 'N') // North
		grid = shiftRocksDirection(grid, 'W') // West
		grid = shiftRocksDirection(grid, 'S') // South
		grid = shiftRocksDirection(grid, 'E') // East
	}

	n := len(grid)
	total := 0
	for i, row := range grid {
		factor := n - i
		count := strings.Count(string(row), "O")
		total += factor * count
	}
	fmt.Println(total)
}

// func shiftGrid(grid [][]byte) [][]byte {
// 	for col := 0; col < len(grid[0]); col++ {
// 		for row := 0; row < len(grid); row++ {
// 			if grid[row][col] == 'O' {
// 				shiftRockUp(grid, row, col)
// 			}
// 		}
// 	}
// 	return grid
// }

// func shiftRockUp(grid [][]byte, row, col int) {
// 	for row > 0 && grid[row-1][col] == '.' {
// 		grid[row-1][col] = 'O' // Move rock up
// 		grid[row][col] = '.'   // Clear the old position
// 		row--                  // Update row to new position
// 	}
// }

func shiftRocksDirection(grid [][]byte, direction byte) [][]byte {
	numRows := len(grid)
	numCols := len(grid[0])

	switch direction {
	case 'N':
		for col := 0; col < numCols; col++ {
			for row := 0; row < numRows; row++ {
				if grid[row][col] == 'O' {
					grid = shiftRock(grid, row, col, -1, 0)
				}
			}
		}
	case 'W':
		for row := 0; row < numRows; row++ {
			for col := 0; col < numCols; col++ {
				if grid[row][col] == 'O' {
					grid = shiftRock(grid, row, col, 0, -1)
				}
			}
		}
	case 'S':
		for col := 0; col < numCols; col++ {
			for row := numRows - 1; row >= 0; row-- {
				if grid[row][col] == 'O' {
					grid = shiftRock(grid, row, col, 1, 0)
				}
			}
		}
	case 'E':
		for row := 0; row < numRows; row++ {
			for col := numCols - 1; col >= 0; col-- {
				if grid[row][col] == 'O' {
					grid = shiftRock(grid, row, col, 0, 1)
				}
			}
		}
	}

	return grid
}

func shiftRock(grid [][]byte, row, col, dr, dc int) [][]byte {
	for isValidPosition(grid, row+dr, col+dc) && grid[row+dr][col+dc] == '.' {
		grid[row+dr][col+dc] = 'O' // Move rock in the direction of the offset
		grid[row][col] = '.'       // Clear the old position
		row += dr                  // Update row to new position
		col += dc                  // Update col to new position
	}
	return grid
}

func isValidPosition(grid [][]byte, row, col int) bool {
	numRows := len(grid)
	numCols := len(grid[0])
	return row >= 0 && row < numRows && col >= 0 && col < numCols
}

func parseInput(fname string) [][]byte {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rawGrid := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		bLine := make([]byte, len(line))
		for i, char := range line {
			bLine[i] = byte(char)
		}
		rawGrid = append(rawGrid, bLine)
	}
	return rawGrid
}
