package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	patterns := parseInput("input.txt")
	part1(patterns)
	part2(patterns)
}

func part1(patterns [][]string) {
	total := 0
	for _, pattern := range patterns {

		// Calculate row bitmask
		rowBitmasks := make([]int, len(pattern))
		for i, row := range pattern {
			rowBitmasks[i] = bitmask(row)
		}

		// Check rows
		mirror, idx := findCutIndex(rowBitmasks)
		if mirror {
			total += 100 * (idx + 1)
			continue
		}

		// Calculate col bitmask
		colBitmasks := make([]int, len(pattern[0]))
		for j := 0; j < len(pattern[0]); j++ {
			col := make([]byte, len(pattern))
			for i := 0; i < len(pattern); i++ {
				col[i] = pattern[i][j]
			}
			colBitmasks[j] = bitmask(string(col))
		}

		// Check cols
		mirror, idx = findCutIndex(colBitmasks)
		if mirror {
			total += idx + 1
		}
	}
	fmt.Println(total)
}

func part2(patterns [][]string) {
	total := 0
	for _, pattern := range patterns {

		// Calculate row bitmask
		rowBitmasks := make([]int, len(pattern))
		for i, row := range pattern {
			rowBitmasks[i] = bitmask(row)
		}

		// Check rows
		mirror, subtotal := findCutIndexHamming(rowBitmasks)
		if mirror {
			total += 100 * (subtotal + 1)
			continue
		}

		// Calculate col bitmask
		colBitmasks := make([]int, len(pattern[0]))
		for j := 0; j < len(pattern[0]); j++ {
			col := make([]byte, len(pattern))
			for i := 0; i < len(pattern); i++ {
				col[i] = pattern[i][j]
			}
			colBitmasks[j] = bitmask(string(col))
		}

		// Check cols
		mirror, subtotal = findCutIndexHamming(colBitmasks)
		if mirror {
			total += subtotal + 1
		}
	}
	fmt.Println(total)
}

func findCutIndex(bitmask []int) (bool, int) {
	for i := 0; i < len(bitmask)-1; i++ {
		if bitmask[i] == bitmask[i+1] {
			j, k := 0, 1
			for i-j >= 0 && i+k < len(bitmask) {
				if bitmask[i-j] != bitmask[i+k] {
					break // Asymmetry found, break the search loop
				}
				j++
				k++
			}
			if i-j < 0 || i+k >= len(bitmask) {
				return true, i
			}
		}
	}
	return false, 0
}

func findCutIndexHamming(bitmask []int) (bool, int) {
	editDist := 0
	for i := 0; i < len(bitmask)-1; i++ {
		editDist = 0 // Reset the edit dist
		j, k := 0, 1
		for i-j >= 0 && i+k < len(bitmask) {
			if bitmask[i-j] == bitmask[i+k] {
				j++
				k++
			} else if hammingDistance(bitmask[i-j], bitmask[i+k]) == 1 {
				editDist++
				j++
				k++
			} else {
				break // Asymmetry found, break the search loop
			}
		}

		// All patterns have exactly 1 smudge
		if (i-j < 0 || i+k >= len(bitmask)) && editDist == 1 {
			return true, i
		}
	}
	return false, 0
}

func bitmask(line string) int {
	var num int
	for _, char := range line {
		num <<= 1 // Left shift by 1
		if char == '#' {
			num |= 1
		}
	}
	return num
}

func hammingDistance(a, b int) int {
	xor := a ^ b
	distance := 0
	for xor > 0 {
		distance += xor & 1
		xor >>= 1
	}
	return distance
}

func parseInput(fname string) [][]string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := make([][]string, 0)
	pattern := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = nil
		} else {
			pattern = append(pattern, line)
		}
	}

	// Append the last pattern
	patterns = append(patterns, pattern)
	return patterns
}
