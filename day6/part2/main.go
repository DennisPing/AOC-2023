package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(content), "\n")

	rawTimes := strings.Split(lines[0], ":")[1]
	fields1 := strings.Fields(rawTimes)
	timeStr := strings.Join(fields1, "")

	rawDistances := strings.Split(lines[1], ":")[1]
	fields2 := strings.Fields(rawDistances)
	distanceStr := strings.Join(fields2, "")

	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distanceStr)

	tt := float64(time)
	dd := float64(dist)
	totalWins := 1

	// y = x(c-x) or x^2-cx+y = 0
	x1, x2 := SolveQuadratic(1, -tt, dd)

	x1 = math.Floor(x1) // x1 is the higher one lol
	x2 = math.Ceil(x2)  // x2 is the lower one

	if x1*(tt-x1) <= dd {
		x1--
	}

	if x2*(tt-x2) <= dd {
		x2++
	}
	totalWins *= int(x1 - x2 + 1)
	fmt.Println(totalWins)
}

func Parse1DArray(line string) []int {
	var result []int
	for _, field := range strings.Fields(line) {
		num, _ := strconv.Atoi(field)
		result = append(result, num)
	}
	return result
}

// ax^2 + bx + c = 0
func SolveQuadratic(a, b, c float64) (float64, float64) {
	d := b*b - 4*a*c
	if d < 0 {
		log.Fatalln("no real roots") // Should not happen
	}

	dSqrt := math.Sqrt(d)
	root1 := (-b + dSqrt) / (2 * a)
	root2 := (-b - dSqrt) / (2 * a)

	if root1 < 0 || root2 < 0 {
		log.Fatalln("negative roots:", root1, root2) // Should not happen
	}

	return root1, root2
}
