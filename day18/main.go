package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	r, c int
}

type Plan struct {
	dr, dc int
	steps  int
}

func main() {
	plans, totalSteps := parseInput1("input.txt")
	ans := part1(plans, totalSteps)
	fmt.Println(ans)

	plans, totalSteps = parseInput2("input.txt")
	ans = part1(plans, totalSteps)
	fmt.Println(ans)
}

func part1(plans []Plan, totalSteps int) int {
	minRows, minCols := math.MaxInt, math.MaxInt
	r, c := 0, 0

	// Walk the plan
	points := make([]Point, totalSteps)
	n := 0
	for _, plan := range plans {
		for i := 0; i < plan.steps; i++ {
			points[n+i] = Point{r, c}
			r += plan.dr
			c += plan.dc
		}
		n += plan.steps // Increment new offset

		minRows = min(r, minRows)
		minCols = min(c, minCols)
	}

	// Normalize all points
	offsetR := abs(minRows)
	offsetC := abs(minCols)
	normalizedPoints := make([]Point, len(points))
	for i, point := range points {
		normalizedPoints[i] = Point{point.r + offsetR, point.c + offsetC}
	}

	// Pick's Theorem
	// A = i + b/2 - 1
	b := len(normalizedPoints)            // boundary
	area := polygonArea(normalizedPoints) // area

	// i = a + 1 - b/2
	interior := area + 1 - b/2

	return interior + b
}

// Shoelace formula
func polygonArea(points []Point) int {
	n := len(points)
	sum1, sum2 := 0, 0

	// Dot product
	for i := range points[:n-1] {
		sum1 += points[i].r * points[i+1].c
		sum2 += points[i].c * points[i+1].r
	}

	// Dot product of final rollover sum
	sum1 += points[n-1].r * points[0].c
	sum2 += points[0].r * points[n-1].c

	area := abs(sum1-sum2) / 2
	return area
}

func parseInput1(fname string) ([]Plan, int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plans := make([]Plan, 0)
	totalSteps := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		dr, dc := 0, 0
		switch parts[0] {
		case "U":
			dr, dc = -1, 0
		case "D":
			dr, dc = 1, 0
		case "L":
			dr, dc = 0, -1
		case "R":
			dr, dc = 0, 1
		}

		steps, _ := strconv.Atoi(parts[1])
		totalSteps += steps
		plan := Plan{
			dr:    dr,
			dc:    dc,
			steps: steps,
		}
		plans = append(plans, plan)
	}
	return plans, totalSteps
}

func parseInput2(fname string) ([]Plan, int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plans := make([]Plan, 0)
	totalSteps := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		dr, dc := 0, 0
		n := len(parts[2])
		hexSteps := parts[2][2 : n-2]
		steps, _ := strconv.ParseInt(hexSteps, 16, 32)
		totalSteps += int(steps)
		dir := parts[2][7]
		switch dir {
		case '0':
			dr, dc = 0, 1
		case '1':
			dr, dc = 1, 0
		case '2':
			dr, dc = 0, -1
		case '3':
			dr, dc = -1, 0
		}
		plan := Plan{
			dr:    dr,
			dc:    dc,
			steps: int(steps),
		}
		plans = append(plans, plan)
	}
	return plans, totalSteps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
