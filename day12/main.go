package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	configs, arrLengths := parseInput("input.txt")

	part1(configs, arrLengths)
	part2(configs, arrLengths)
}

func part1(configs []string, arrLengths [][]int) {
	total := 0
	for x, config := range configs {
		arrLength := arrLengths[x]

		// Buffer the front and back of the config
		bufferedConfig := make([]byte, len(config)+2)
		bufferedConfig[0] = '.'
		bufferedConfig[len(bufferedConfig)-1] = '.'
		copy(bufferedConfig[1:], config)

		// Expand the array lengths into a boolean array
		springs := make([]bool, 0)
		springs = append(springs, false) // Pad the front

		for _, length := range arrLength {
			for j := 0; j < length; j++ {
				springs = append(springs, true)
			}
			springs = append(springs, false) // Pad the back
		}

		total += solve(bufferedConfig, springs)
	}

	fmt.Println(total)
}

func part2(configs []string, arrLengths [][]int) {
	total := 0
	for x, config := range configs {
		arrLength := arrLengths[x]

		bigConfig := ""
		bigArrLengths := make([]int, 0)
		for i := 0; i < 5; i++ {
			chunk := config + "?"
			bigConfig += chunk
			bigArrLengths = append(bigArrLengths, arrLength...)
		}

		// Trim off the last ?
		bigConfig = bigConfig[:len(bigConfig)-1]

		// Buffer the front and back of the config
		bufferedConfig := make([]byte, len(bigConfig)+2)
		bufferedConfig[0] = '.'
		bufferedConfig[len(bufferedConfig)-1] = '.'
		copy(bufferedConfig[1:], bigConfig)

		// Expand the array lengths into a boolean array (eg. 1 1 3 -> T F T F T T T
		springs := make([]bool, 0)
		springs = append(springs, false) // Pad the front
		for _, length := range bigArrLengths {
			for j := 0; j < length; j++ {
				springs = append(springs, true)
			}
			springs = append(springs, false) // Pad the back
		}

		total += solve(bufferedConfig, springs)
	}

	fmt.Println(total)
}

// Iterative bottom up approach
func solve(config []byte, springs []bool) int {
	n := len(config)
	m := len(springs)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[n][m] = 1

	for i := n - 1; i >= 0; i-- { // Iterate over config characters
		for j := m - 1; j >= 0; j-- { // Iterate over spring bitmap
			damaged := false
			working := false

			switch config[i] {
			case '#':
				damaged = true
			case '.':
				working = true
			case '?':
				damaged = true
				working = true
			}

			var sum int
			if damaged && springs[j] {
				// If damaged and the spring slot is free
				sum = dp[i+1][j+1]

			} else if working && !springs[j] {
				// Else if working and spring slot is closed
				sum = dp[i+1][j+1] + dp[i+1][j]
			}
			dp[i][j] = sum
		}

	}
	return dp[0][0]
}

func parseInput(fname string) ([]string, [][]int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var configs []string
	var arrLengths [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		configs = append(configs, parts[0])

		fields := strings.Split(parts[1], ",")
		group := make([]int, len(fields))
		for i, val := range fields {
			num, _ := strconv.Atoi(val)
			group[i] = num
		}
		arrLengths = append(arrLengths, group)
	}

	return configs, arrLengths
}
