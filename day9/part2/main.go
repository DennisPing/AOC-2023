package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Do the reverse of part 1

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	histories := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))
		for i := range fields {
			temp, _ := strconv.Atoi(fields[i])
			numbers[i] = temp
		}
		histories = append(histories, numbers)
	}

	total := 0
	for _, history := range histories {
		total += nextHistoryValue(history)
	}
	fmt.Println(total)
}

func nextHistoryValue(history []int) int {
	// fmt.Println(history)
	sum := 0
	for _, num := range history {
		sum += num
	}

	if sum == 0 {
		return 0 // Base case
	}

	diffs := make([]int, len(history)-1)
	for i := len(history) - 1; i >= 1; i-- {
		diff := history[i] - history[i-1]
		diffs[i-1] = diff
	}

	n := nextHistoryValue(diffs)
	newValue := history[0] - n
	return newValue
}
