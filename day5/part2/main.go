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

type Section struct {
	Name string
	Data [][]int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rawSeeds []int
	var sections []*Section
	var rawLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // End of a section
			ParseSection(rawLines, &rawSeeds, &sections)
			rawLines = []string{} // Reset for next section
		} else {
			rawLines = append(rawLines, line)
		}
	}

	// Process the last section because it doesn't have a blank line
	ParseSection(rawLines, &rawSeeds, &sections)

	// fmt.Println(rawSeeds)
	// for _, v := range sections {
	// 	fmt.Printf("%+v\n", *v)
	// }

	// Calculate the total size required for seeds array
	totalSize := 0
	for i := 0; i < len(rawSeeds); i += 2 {
		totalSize += rawSeeds[i+1]
	}

	seeds := make([]int, totalSize)

	// Assign all the seed values
	curr := 0
	for i := 0; i < len(rawSeeds); i += 2 {
		start := rawSeeds[i]
		length := rawSeeds[i+1]
		for j := 0; j < length; j++ {
			seeds[curr] = start + j
			curr++
		}
	}

	fmt.Println(len(seeds))

	CalculateMinLocation(seeds, sections)
}

func CalculateMinLocation(seeds []int, sections []*Section) {
	minLoc := math.MaxInt
	locations := make([]int, 7) // Re-use this size 7 array
	for _, seed := range seeds {
		for i, section := range sections {
			for _, line := range section.Data {
				newLocation, found := MapLocation(seed, line)
				if found {
					seed = newLocation
					break
				}
			}
			locations[i] = seed
		}
		// fmt.Println(locations)
		last := locations[len(locations)-1]
		minLoc = Min(minLoc, last)
	}
	fmt.Println(minLoc)
}

// Checks if mapping needs to be done, and if yes, returns the new mapped location
func MapLocation(location int, line []int) (int, bool) {
	src := line[1]
	dst := line[0]
	len := line[2]
	offset := int(math.Abs(float64(src - dst)))

	if location >= src && location <= src+len-1 {
		if src > dst {
			return location - offset, true // Need to shift location down
		} else {
			return location + offset, true // Need to shift location up
		}
	}
	return 0, false
}

func ParseSection(rawLines []string, seeds *[]int, sections *[]*Section) {
	if len(rawLines) == 0 {
		return
	}

	switch {
	case strings.Contains(rawLines[0], "seeds"):
		*seeds = Parse1DArray(strings.Split(rawLines[0], ":")[1])
	case strings.Contains(rawLines[0], "map"):
		name := strings.Split(rawLines[0], " ")[0]
		data := Parse2DArray(rawLines[1:])
		section := &Section{
			Name: name,
			Data: data,
		}
		*sections = append(*sections, section)

	default:
		log.Fatalf("unable to parse lines: %v\n", rawLines)
	}
}

func Parse1DArray(line string) []int {
	var result []int
	for _, field := range strings.Fields(line) {
		num, _ := strconv.Atoi(field)
		result = append(result, num)
	}
	return result
}

func Parse2DArray(lines []string) [][]int {
	var array [][]int
	for _, line := range lines {
		ints := Parse1DArray(line)
		array = append(array, ints)
	}
	return array
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
