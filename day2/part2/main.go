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
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		sets := strings.Split(parts[1], ";")

		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, set := range sets {
			// A set is one semicolon group. Eg: 3 blue, 4 red;
			maxRed, maxGreen, maxBlue = calculateMaxes(set, maxRed, maxGreen, maxBlue)
		}

		power := 1
		if maxRed > 0 {
			power *= maxRed
		}
		if maxGreen > 0 {
			power *= maxGreen
		}
		if maxBlue > 0 {
			power *= maxBlue
		}

		total += power
	}

	fmt.Println(total)
}

func calculateMaxes(set string, maxRed int, maxGreen int, maxBlue int) (int, int, int) {
	hands := strings.Split(set, ",")
	cleanHands := make([]string, len(hands))

	for i := 0; i < len(hands); i++ {
		cleanHands[i] = strings.TrimSpace(hands[i])
	}

	for _, hand := range cleanHands {
		parts := strings.Split(hand, " ")
		count, _ := strconv.Atoi(parts[0])
		color := parts[1]

		switch color {
		case "red":
			if count > maxRed {
				maxRed = count
			}
		case "green":
			if count > maxGreen {
				maxGreen = count
			}
		case "blue":
			if count > maxBlue {
				maxBlue = count
			}
		}
	}

	return maxRed, maxGreen, maxBlue
}
