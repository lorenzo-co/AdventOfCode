package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Coordinates struct {
	Row    int
	Column int
}

func main() {
	content, err := os.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	totalAsterisks := 0
	possibleGears := 0

	for row, line := range lines {
		gearsColumn := findGear(line, row, lines)

		totalAsterisks += len(gearsColumn)
		for _, col := range gearsColumn {
			adjacentNumbers := hasAdjacentNumbers(row, col, 1, lines)

			fmt.Println(row, adjacentNumbers)
		}
	}

	fmt.Println(totalAsterisks)
	fmt.Println(possibleGears)
}

func findGear(line string, row int, lines []string) []int {
	gearsCol := []int{}

	for pos, val := range line {
		if string(val) == "*" {
			gearsCol = append(gearsCol, pos)
		}
	}

	return gearsCol
}

func hasAdjacentNumbers(row int, start int, length int, lines []string) []int {
	var safeStart int
	var safeThreeBefore int
	var safeEnd int
	var safeThreeAfter int

	numbers := []int{}

	if start > 2 {
		safeStart = start - 1
		safeThreeBefore = start - 3
	} else {
		if start > 0 {
			safeStart = start - 1
		} else {
			safeStart = start
		}
	}

	if start < len(lines[row])-3 {
		safeEnd = start + 1
		safeThreeAfter = start + 4
	} else {
		if start < len(lines[row])-1 {
			safeStart = start + 1
		} else {
			safeStart = start
		}
	}

	if row > 0 {
		textNumbers := strings.Split(strings.Trim(lines[row-1][safeThreeBefore:safeThreeAfter], "."), ".")

		for _, text := range textNumbers {
			i, err := strconv.Atoi(text)
			if err == nil {
				numbers = append(numbers, i)
			}
		}
	}

	if unicode.IsDigit(rune(lines[row][safeStart])) {
		text := strings.Trim(lines[row][safeThreeBefore:safeStart+1], ".")

		i, err := strconv.Atoi(text)
		if err == nil {
			numbers = append(numbers, i)
		}
	}

	if unicode.IsDigit(rune(lines[row][safeEnd])) {
		text := strings.Trim(lines[row][safeEnd:safeThreeAfter], ".")

		i, err := strconv.Atoi(text)
		if err == nil {
			numbers = append(numbers, i)
		}
	}

	if row > 0 {
		textNumbers := strings.Split(strings.Trim(lines[row+1][safeThreeBefore:safeThreeAfter], "."), ".")

		for _, text := range textNumbers {
			i, err := strconv.Atoi(text)
			if err == nil {
				numbers = append(numbers, i)
			}
		}
	}

	return numbers
}

/* func linePartsSum(line string, row int, lines []string) int {
	lineSum := 0
	for col := 0; col < len(line); col++ {
		if unicode.IsDigit(rune(line[col])) {
			numberLength := getNumberLength(line, col)

			if hasAdjacentSymbol(line, row, col, numberLength, lines) {
				number, err := strconv.Atoi(line[col : col+numberLength])

				if err != nil {
					panic(err)
				}

				lineSum += number
			}

			col += numberLength
		}
	}
	return lineSum
}

func getNumberLength(line string, index int) int {
	length := 1
	for index < len(line) && unicode.IsDigit(rune(line[index+1])) {
		length += 1
		index += 1
	}

	return length
}
*/
