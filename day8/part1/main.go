package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
		}
	}

	steps := navigate(instructions, network)
	fmt.Println(steps)
}

func navigate(instructions string, network map[string]*Branch) int {
	steps := 0
	curr := "AAA"
	goal := "ZZZ"

	i := 0
	for {
		dir := instructions[i]
		if curr == goal {
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
