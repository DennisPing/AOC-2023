package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gammazero/deque"
)

// This is BFS

type Coord [2]int

var directions = []struct {
	dr, dc     int
	curr, next string
}{
	{-1, 0, "S|JL", "|7F"}, // North
	{1, 0, "S|7F", "|JL"},  // South
	{0, -1, "S-J7", "-LF"}, // West
	{0, 1, "S-LF", "-J7"},  // East
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var curr Coord
	grid := make([]string, 0)

	r := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			if c == 'S' {
				curr = [2]int{r, i}
			}
		}

		grid = append(grid, line)
		r++
	}

	visited := make(map[Coord]bool)
	q := deque.New[Coord]()

	q.PushBack(curr)
	visited[curr] = true

	// BFS
	for q.Len() > 0 {
		coord := q.PopFront()
		r, c := coord[0], coord[1]
		currChar := string(grid[r][c])

		// Check all possible 4 directions
		for _, dir := range directions {
			newR, newC := r+dir.dr, c+dir.dc

			// If within bounds
			if newR >= 0 && newR < len(grid) && newC >= 0 && newC < len(grid[0]) {
				nextChar := string(grid[newR][newC])
				newCoord := Coord{newR, newC}

				// Check if current char is valid and if next char is valid
				if strings.ContainsAny(currChar, dir.curr) && strings.ContainsAny(nextChar, dir.next) && !visited[newCoord] {
					visited[newCoord] = true
					q.PushBack(newCoord)
				}
			}
		}
	}

	n := len(visited)
	fmt.Println(n / 2)
}
