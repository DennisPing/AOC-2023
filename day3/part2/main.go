package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Point struct {
	R int
	C int
}

type PartNumber struct {
	Value int
	Point Point
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Get the line length
	lineLen := 0
	if scanner.Scan() {
		lineLen = len(scanner.Text())
	}

	// Reset the file pointer
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	scanner = bufio.NewScanner(file)
	grid := make([][]byte, 0)

	// Pad all sides so we don't have to check bounds
	grid = append(grid, DummyLine(lineLen+2))
	for scanner.Scan() {
		line := scanner.Bytes()

		paddedLine := make([]byte, 0)
		paddedLine = append(paddedLine, '.')
		paddedLine = append(paddedLine, line...)
		paddedLine = append(paddedLine, '.')

		grid = append(grid, paddedLine)
	}
	grid = append(grid, DummyLine(lineLen+2))

	gearList := make([]Point, 0)
	pnSet := make(map[PartNumber]bool)

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {

			// Find all numbers, check if it has a gear neighbor. If yes, save that gear's Point.
			if IsDigit(grid[i][j]) {
				num, size := GetNumber(grid[i][:], j)

				upper := grid[i-1][j-1 : j+size+1]
				mid := grid[i][j-1 : j+size+1]
				lower := grid[i+1][j-1 : j+size+1]

				point, found := FindGear(upper, mid, lower, i, j)
				if found {
					pn := PartNumber{
						Value: num,
						Point: point,
					}
					pnSet[pn] = true
				}
				j += size - 1 // Skip the rest of the found number
			} else if grid[i][j] == '*' {
				point := Point{
					R: i,
					C: j,
				}
				gearList = append(gearList, point)
			}
		}
	}

	// for each, _ := range pnSet {
	// 	fmt.Printf("%+v\n", each)
	// }
	// fmt.Printf("%v\n", gearList)

	// Loop through the gear list. Find all numbers which have the same Point.
	total := 0
	for _, gear := range gearList {
		nums := make([]int, 0)
		for pn := range pnSet {
			if pn.Point == gear {
				nums = append(nums, pn.Value)
				delete(pnSet, pn)
			}
		}
		if len(nums) == 2 {
			total += (nums[0] * nums[1])
		}
	}

	fmt.Println(total)
}

// Get the entire number and the number of digits. Eg: "567" -> (567, 3)
func GetNumber(line []byte, i int) (int, int) {
	j := i
	for IsDigit(line[j]) {
		j++
	}
	num, _ := strconv.Atoi(string(line[i:j]))
	return num, j - i
}

// Find the possible gear 'Point' in the upper, mid, and lower rows
func FindGear(upper, mid, lower []byte, rOff, cOff int) (Point, bool) {
	/*
		.....* (upper)
		.1234. (mid)
		...... (lower)
	*/
	sizeOfNum := len(upper) - 2

	for i, b := range upper {
		if IsGear(b) {
			return Point{
				R: rOff - 1,
				C: cOff - 1 + i,
			}, true
		}
	}
	for i, b := range lower {
		if IsGear(b) {
			return Point{
				R: rOff + 1,
				C: cOff - 1 + i,
			}, true
		}
	}
	if IsGear(mid[0]) {
		return Point{
			R: rOff,
			C: cOff - 1,
		}, true
	}
	if IsGear(mid[len(mid)-1]) {
		return Point{
			R: rOff,
			C: cOff + sizeOfNum,
		}, true
	}

	return Point{}, false
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsGear(char byte) bool {
	return char == '*'
}

func DummyLine(size int) []byte {
	line := make([]byte, size)
	for i := 0; i < size; i++ {
		line[i] = '.'
	}
	return line
}
