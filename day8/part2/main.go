package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Note: Brute force would take several weeks to compute, so need to compute the LCM

type Branch struct {
	Left  string
	Right string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	network := make(map[string]*Branch)
	firstLine := true
	var instructions string
	var startingNodes []string

	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			instructions = line
			firstLine = false
			continue
		}
		if line != "" {
			key := line[0:3]
			left := line[7:10]
			right := line[12:15]
			network[key] = &Branch{
				Left:  left,
				Right: right,
			}

			if key[2] == 'A' {
				startingNodes = append(startingNodes, key)
			}
		}
	}

	allSteps := make([]int, len(startingNodes))
	for i, node := range startingNodes {
		allSteps[i] = navigate(instructions, node, network)
	}

	total := lcmArray(allSteps)
	fmt.Println(total)
}

func navigate(instructions string, startingNode string, network map[string]*Branch) int {
	steps := 0
	curr := startingNode

	i := 0
	for {
		dir := instructions[i]
		if curr[2] == byte('Z') {
			return steps
		}

		var next string
		if dir == 'L' {
			next = network[curr].Left
		} else {
			next = network[curr].Right
		}

		curr = next
		steps++
		i++

		if i > len(instructions)-1 {
			i = 0
		}
	}
}

// Find the least common multiple from an array of numbers
// https://www.tutorialsfreak.com/java-tutorial/examples/lcm-array
func lcmArray(array []int) int {
	lcm := array[0]
	for i := 0; i < len(array); i++ {
		lcm = (lcm * array[i]) / gcd(lcm, array[i])
	}
	return lcm
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
