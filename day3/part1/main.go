package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
	schematic := make([][]byte, 0)

	// Append all sides with a buffer so we don't have to check bounds
	schematic = append(schematic, dummyLine(lineLen+2))
	for scanner.Scan() {
		line := scanner.Bytes()

		bufferedLine := make([]byte, 0)
		bufferedLine = append(bufferedLine, '.')
		bufferedLine = append(bufferedLine, line...)
		bufferedLine = append(bufferedLine, '.')

		schematic = append(schematic, bufferedLine)
	}
	schematic = append(schematic, dummyLine(lineLen+2))

	total := 0
	for i := 1; i < len(schematic)-1; i++ {
		for j := 1; j < len(schematic)-1; j++ {
			if isDigit(schematic[i][j]) {
				num, size := getNumber(schematic[i][:], j)

				upper := schematic[i-1][j-1 : j+size+1]
				mid := schematic[i][j-1 : j+size+1]
				lower := schematic[i+1][j-1 : j+size+1]
				if checkAdj(upper, mid, lower) {
					total += num
				}
				j += size
			}
		}
	}

	fmt.Println(total)
}

// Get the entire number and the number of digits. Eg: "567" -> (567, 3)
func getNumber(line []byte, i int) (int, int) {
	j := i
	for isDigit(line[j]) {
		j++
	}
	num, _ := strconv.Atoi(string(line[i:j]))
	return num, j - i
}

func checkAdj(upper, mid, lower []byte) bool {
	/*
		.....+ (upper)
		.1234. (mid)
		...... (lower)
	*/

	// fmt.Println(string(upper))
	// fmt.Println(string(mid))
	// fmt.Println(string(lower))
	// fmt.Println()

	for _, b := range upper {
		if isSymbol(b) {
			return true
		}
	}
	for _, b := range lower {
		if isSymbol(b) {
			return true
		}
	}
	if isSymbol(mid[0]) {
		return true
	}
	if isSymbol(mid[len(mid)-1]) {
		return true
	}
	return false
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isSymbol(char byte) bool {
	return !isDigit(char) && char != '.'
}

func dummyLine(size int) []byte {
	line := make([]byte, size)
	for i := 0; i < size; i++ {
		line[i] = '.'
	}
	return line
}
