package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Lense struct {
	label string
	focal int
}

func main() {
	inputs := parseInput("input.txt")

	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
	total := 0
	for _, input := range inputs {
		total += hash(input)
	}
	fmt.Println(total)
}

func part2(inputs []string) {
	table := make([][]Lense, 256)
	for _, input := range inputs {
		opIdx := strings.IndexAny(input, "=-")
		label := input[:opIdx]
		op := input[opIdx]
		hashcode := hash(label)

		if op == '=' {
			focal, _ := strconv.Atoi(string(input[opIdx+1:]))
			lense := Lense{
				label: label,
				focal: focal,
			}
			contains, idx := containsLabel(label, table[hashcode])
			if contains {
				// Label already exists, so replace the lense
				table[hashcode][idx] = lense
			} else {
				// Label does not exist, so append the lense
				table[hashcode] = append(table[hashcode], lense)
			}
		} else {
			// op is -
			contains, idx := containsLabel(label, table[hashcode])
			if contains {
				newSlice := slices.Delete(table[hashcode], idx, idx+1)
				table[hashcode] = newSlice
			}
		}
	}

	total := 0
	for i, list := range table {
		for j, lense := range list {
			total += (i + 1) * (j + 1) * (lense.focal)
		}
	}
	fmt.Println(total)
}

// Checks if the label is in the list, as well as the index of the label
func containsLabel(label string, list []Lense) (bool, int) {
	for i, lense := range list {
		if lense.label == label {
			return true, i
		}
	}
	return false, -1
}

func hash(input string) int {
	subtotal := 0
	for _, char := range input {
		subtotal = ((subtotal + int(char)) * 17) % 256
	}
	return subtotal
}

func parseInput(fname string) []string {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(content), ",")
}
