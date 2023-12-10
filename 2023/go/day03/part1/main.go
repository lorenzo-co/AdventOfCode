package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	sum := 0

	for row, line := range lines {
		sum += linePartsSum(line, row, lines)
	}

	fmt.Println(sum)
}

func linePartsSum(line string, row int, lines []string) int {
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

func hasAdjacentSymbol(line string, row int, start int, length int, lines []string) bool {
	var safeStart int
	re := regexp.MustCompile("[#$+/*%@&=-]")

	if start > 0 {
		safeStart = start - 1
	} else {
		safeStart = start
	}

	if row > 0 {
		for col := safeStart; col <= start+length && col < len(line); col++ {

			if lines[row-1][col] != '.' && re.MatchString(string(lines[row-1][col])) {
				return true
			}
		}
	}

	if lines[row][safeStart] != '.' && re.MatchString(string(lines[row][safeStart])) {
		// fmt.Print(row, string(lines[row][safeStart]))
		return true
	}

	if lines[row][safeStart+length+1] != '.' && re.MatchString(string(lines[row][safeStart+length+1])) {
		//fmt.Print(row, string(lines[row][safeStart+length+1]))

		return true
	}

	if row < len(lines)-1 {
		for col := safeStart; col <= start+length && col < len(line); col++ {
			if lines[row+1][col] != '.' && re.MatchString(string(lines[row+1][col])) {
				return true
			}
		}
	}

	return false
}
