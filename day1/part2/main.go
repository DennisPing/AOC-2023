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

// https://adventofcode.com/2023/day/1

var table = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var array = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

		line = replaceWords(line)

		asciiNum := "" // We need to concat two numbers together

		// Search forwards
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				asciiNum += string(line[i])
				break
			}
		}

		// Search backwards
		for j := len(line) - 1; j >= 0; j-- {
			if line[j] >= '0' && line[j] <= '9' {
				asciiNum += string(line[j])
				break
			}
		}

		num, _ := strconv.Atoi(asciiNum)
		total += num
	}

	fmt.Println(total)
}

// Replace the first and last words, if possible
func replaceWords(line string) string {

	// Replace the first word if possible
	matchIdx := math.MaxUint32
	word := ""

	for _, substr := range array {
		index := strings.Index(line, substr)
		if index >= 0 && index < matchIdx {
			matchIdx = index
			word = substr
		}
	}

	if word != "" {
		newWord := fmt.Sprintf("%s%s", table[word], string(word[len(word)-1])) // eight -> 8t
		line = strings.Replace(line, word, newWord, 1)
	}

	// Replace the last word if possible
	matchIdx = 0
	word = ""

	for _, substr := range array {
		index := strings.LastIndex(line, substr)
		if index >= 0 && index > matchIdx {
			matchIdx = index
			word = substr
		}
	}

	if word != "" {
		newWord := fmt.Sprintf("%s%s", string(word[0]), table[word]) // eight -> e8
		revLine := reverseString(line)
		revLine = strings.Replace(revLine, reverseString(word), reverseString(newWord), 1)
		line = reverseString(revLine)
	}

	return line
}

// https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/
func reverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
