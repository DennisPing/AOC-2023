package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://adventofcode.com/2023/day/1

// 1. Read each line
// 2. Init a total value
// 3. For each line, find the first numerical char and last numerical char.
// 4. Combine them, turn it into a number, and add it to the running total.

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		data := scanner.Bytes()
		asciiNum := "" // We need to concat two numbers together

		// Search forwards
		for i := 0; i < len(data); i++ {
			if data[i] >= '0' && data[i] <= '9' {
				asciiNum += string(data[i])
				break
			}
		}

		// Search backwards
		for j := len(data) - 1; j >= 0; j-- {
			if data[j] >= '0' && data[j] <= '9' {
				asciiNum += string(data[j])
				break
			}
		}

		num, _ := strconv.Atoi(asciiNum)
		total += num
	}

	fmt.Println(total)
}
