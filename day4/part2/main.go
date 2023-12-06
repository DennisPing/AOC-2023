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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	var matches []int

	for scanner.Scan() {
		leftNums := make(map[int]bool)

		line := scanner.Text()
		split := strings.Split(line, " | ")
		leftStr := strings.Split(split[0], ": ")[1]
		arr := StringToArray(leftStr)
		for _, each := range arr {
			leftNums[each] = true
		}

		rightStr := strings.TrimSpace(split[1])
		rightNums := StringToArray(rightStr)

		var count int
		for _, num := range rightNums {
			if leftNums[num] {
				count++
			}
		}
		matches = append(matches, count)
	}

	// fmt.Println(matches)
	// fmt.Println()

	// Build the parallel copies list
	copies := make([]int, len(matches))
	for i := 0; i < len(copies); i++ {
		copies[i] = 1
	}

	// Grow the copies list
	for i := 0; i < len(matches); i++ {
		for j := 0; j < matches[i]; j++ {
			// The subsequent lower copies get additional N nopies from the parent copy
			copies[i+j+1] += copies[i]
		}
		// fmt.Println(copies)
	}

	for _, count := range copies {
		total += count
	}

	fmt.Println(total)
}

// Convert a string of numbers to an array of integers
func StringToArray(str string) []int {
	fields := strings.Fields(str)
	var intArray []int
	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		intArray = append(intArray, num)
	}

	return intArray
}
