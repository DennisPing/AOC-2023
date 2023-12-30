package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type State struct {
	r     int
	c     int
	dr    int
	dc    int
	steps int
}

type Vector struct {
	state State
	cost  int // Minimum running cost
	index int // The index in the heap, self managed by the heap
}

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func part1(grid [][]int) {
	dist := dijkstra(grid, 1, 3)
	fmt.Println(dist)
}

func part2(grid [][]int) {
	dist := dijkstra(grid, 4, 10)
	fmt.Println(dist)
}

func dijkstra(grid [][]int, minStep int, maxStep int) int {
	numRows := len(grid)
	numCols := len(grid[0])

	pqueue := &PriorityQueue{}
	heap.Init(pqueue)

	v0 := &Vector{
		state: State{0, 0, 1, 0, 0},
		cost:  0,
	}

	v1 := &Vector{
		state: State{0, 0, 0, 1, 0},
		cost:  0,
	}

	heap.Push(pqueue, v0)
	heap.Push(pqueue, v1)

	// 4D cost dp array
	dp := make([][][][]int, numRows) // N rows
	for i := range dp {
		dp[i] = make([][][]int, numCols) // M cols
		for j := range dp[i] {
			dp[i][j] = make([][]int, 4) // 4 dirs
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, maxStep+1) // idx 0 is unused
				for step := range dp[i][j][k] {
					dp[i][j][k][step] = math.MaxInt
				}
			}
		}
	}

	for pqueue.Len() > 0 {
		top := heap.Pop(pqueue).(*Vector)
		state, cost := top.state, top.cost

		if state.r == numRows-1 && state.c == numCols-1 {
			// fmt.Println(state, cost)
			return cost // The first time we hit the destination is the answer because of priority queue
		}

		// 0 = up, 1 = down, 2 = left, 3 = right
		for dir, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			dr := offset[0]
			dc := offset[1]
			nr := state.r + dr
			nc := state.c + dc
			if !isValidPosition(nr, nc, numRows, numCols) {
				continue
			}

			// Don't go backwards
			if state.dr == -dr && state.dc == -dc {
				continue
			}

			// Dot product trick to check if direction change is perpendicular
			if state.steps < minStep && ((state.dr*dr)+(state.dc*dc) == 0) {
				continue
			}

			// Check if we hit the maxStep limit
			if state.steps == maxStep && state.dr == dr && state.dc == dc {
				continue
			}

			// Update the step
			nextSteps := 1
			if state.dr == dr && state.dc == dc {
				nextSteps = state.steps + 1
			}

			newCost := cost + grid[nr][nc]
			if newCost < dp[nr][nc][dir][nextSteps] {
				dp[nr][nc][dir][nextSteps] = newCost
				heap.Push(pqueue, &Vector{
					state: State{nr, nc, dr, dc, nextSteps},
					cost:  newCost,
				})
			}
		}
	}
	return -1
}

func isValidPosition(row, col, numRows, numCols int) bool {
	return row >= 0 && row < numRows && col >= 0 && col < numCols
}

func parseInput(fname string) [][]int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := make([][]int, 0)
	n := 0
	first := true

	for scanner.Scan() {
		line := scanner.Text()
		if first {
			n = len(line)
			first = false
		}
		row := make([]int, n)
		for i, char := range line {
			num, _ := strconv.Atoi(string(char))
			row[i] = num
		}
		grid = append(grid, row)
	}
	return grid
}
