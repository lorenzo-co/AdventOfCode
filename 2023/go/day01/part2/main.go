package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		sum += firstAndLastDigitsToIntPart2(scanner.Text())
	}

	fmt.Println(sum)
}

func firstAndLastDigitsToIntPart2(lineWithWords string) int {
	firstAndLastDigits := ""

	line := ""

	for pos, val := range lineWithWords {
		if digit, err := strconv.Atoi(string(val)); err == nil {
			line += strconv.Itoa(digit)
		} else {
			for index, word := range digits {
				if strings.HasPrefix(lineWithWords[pos:], word) {
					line += strconv.Itoa(index + 1)
					break
				}
			}
		}
	}

	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			firstAndLastDigits += string(line[i])
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			firstAndLastDigits += string(line[i])
			break
		}
	}

	value, err := strconv.Atoi(firstAndLastDigits)
	if err != nil {
		panic(err)
	}

	return value
}
