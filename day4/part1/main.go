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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

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

		n := 0
		for _, num := range rightNums {
			if leftNums[num] {
				n++
			}
		}
		total += int(math.Pow(2, float64(n-1)))
	}

	fmt.Println(total)
}

func StringToArray(str string) []int {
	fields := strings.Fields(str)
	var intArray []int
	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		intArray = append(intArray, num)
	}

	return intArray
}
