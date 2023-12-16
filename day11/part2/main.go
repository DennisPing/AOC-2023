package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point [2]int

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	grid := strings.Split(string(content), "\n")

	// Basically a bitmask for all the extra rows and cols
	extraRows := make([]int, len(grid))
	extraCols := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		if !strings.Contains(grid[i], "#") {
			extraRows[i] = 1
		}
	}

	for j := 0; j < len(grid[0]); j++ { // iterate over cols
		found := false
		for i := 0; i < len(grid); i++ { // iterate over rows
			if grid[i][j] == '#' {
				found = true
				break
			}
		}
		if !found {
			extraCols[j] = 1
		}
	}

	// fmt.Println(extraRows)
	// fmt.Println(extraCols)

	galaxies := findGalaxies(grid) // Find hashtag points (#)

	total := 0
	// Calculate the sum of all unique distances between pairs
	for i := 0; i < len(galaxies); i++ {
		p1 := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			p2 := galaxies[j]

			r1 := min(p1[0], p2[0])
			r2 := max(p1[0], p2[0])

			c1 := min(p1[1], p2[1])
			c2 := max(p1[1], p2[1])

			dist := 0
			for r := r1; r < r2; r++ { // Iterate through the specific row range
				if extraRows[r] == 1 {
					dist += 1_000_000
				} else {
					dist += 1
				}
			}

			for c := c1; c < c2; c++ { // Iterate through the specific column range
				if extraCols[c] == 1 {
					dist += 1_000_000
				} else {
					dist += 1
				}
			}
			total += dist
		}
	}

	fmt.Println(total)
}

func findGalaxies(grid []string) []Point {
	points := make([]Point, 0)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '#' {
				points = append(points, Point{r, c})
			}
		}
	}
	return points
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
