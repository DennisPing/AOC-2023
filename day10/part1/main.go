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

	// fmt.Println(grid)
	// fmt.Println(curr)

	found := make(map[Coord]bool)
	q := deque.New[Coord]()

	q.PushBack(curr)
	found[curr] = true

	for q.Len() > 0 {
		coord := q.PopFront()
		r := coord[0]
		c := coord[1]
		char := string(grid[r][c])

		// Check going up (north)
		if r > 0 {
			upchar := string(grid[r-1][c])
			upcoord := [2]int{r - 1, c}
			if strings.ContainsAny(char, "S|JL") && strings.ContainsAny(upchar, "|7F") && !found[upcoord] {
				found[upcoord] = true
				q.PushBack(upcoord)
			}
		}

		// Check going down (south)
		if r < len(grid)-1 {
			downchar := string(grid[r+1][c])
			downcoord := [2]int{r + 1, c}
			if strings.ContainsAny(char, "S|7F") && strings.ContainsAny(downchar, "|JL") && !found[downcoord] {
				found[downcoord] = true
				q.PushBack(downcoord)
			}
		}

		// Check going left (west)
		if c > 0 {
			leftchar := string(grid[r][c-1])
			leftcoord := [2]int{r, c - 1}
			if strings.ContainsAny(char, "S-J7") && strings.ContainsAny(leftchar, "-LF") && !found[leftcoord] {
				found[leftcoord] = true
				q.PushBack(leftcoord)
			}
		}

		// Check going right (east)
		if c < len(grid[0])-1 {
			rightchar := string(grid[r][c+1])
			rightcoord := [2]int{r, c + 1}
			if strings.ContainsAny(char, "S-LF") && strings.ContainsAny(rightchar, "-J7") && !found[rightcoord] {
				found[rightcoord] = true
				q.PushBack(rightcoord)
			}
		}
	}

	n := len(found)
	fmt.Println(n / 2)
}
