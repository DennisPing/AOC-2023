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
		idStr := strings.Split(parts[0], " ")[1]
		id, _ := strconv.Atoi(idStr)

		sets := strings.Split(parts[1], ";")

		valid := true
		for _, set := range sets {
			// A set is one semicolon group. Eg: 3 blue, 4 red;
			if !checkSet(set, 12, 13, 14) {
				valid = false
				break
			}
		}
		if valid {
			total += id
		}
	}

	fmt.Println(total)
}

func checkSet(set string, maxRed int, maxGreen int, maxBlue int) bool {
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
				return false
			}
		case "green":
			if count > maxGreen {
				return false
			}
		case "blue":
			if count > maxBlue {
				return false
			}
		}
	}
	return true
}
