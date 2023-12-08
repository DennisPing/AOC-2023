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
	times := Parse1DArray(rawTimes)

	rawDistances := strings.Split(lines[1], ":")[1]
	distances := Parse1DArray(rawDistances)

	// var numWaysToWin []int
	// for i := 0; i < len(times); i++ {
	// 	time := times[i]
	// 	dist := distances[i]

	// 	// Quadratic equation is: y = x(t-x) or dist = hold(time-hold)

	// 	var wins int
	// 	for x := 0; x < time; x++ {
	// 		y := x*time - int(math.Pow(float64(x), 2))
	// 		if y > dist {
	// 			wins++
	// 		}
	// 	}

	// 	if wins > 0 {
	// 		numWaysToWin = append(numWaysToWin, wins)
	// 	}
	// }

	// fmt.Println(numWaysToWin)
	// totalWins := 1
	// for _, win := range numWaysToWin {
	// 	totalWins *= win
	// }
	// fmt.Println(totalWins)

	// Math way
	totalWins := 1
	for i := 0; i < len(times); i++ {
		time := float64(times[i])
		dist := float64(distances[i])

		// y = x(c-x) or x^2-cx+y = 0
		x1, x2 := SolveQuadratic(1, -time, dist)

		x1 = math.Floor(x1) // x1 is the higher one lol
		x2 = math.Ceil(x2)  // x2 is the lower one

		if x1*(time-x1) <= dist {
			x1--
		}

		if x2*(time-x2) <= dist {
			x2++
		}
		totalWins *= int(x1 - x2 + 1)
	}
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
