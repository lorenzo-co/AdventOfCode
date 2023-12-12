package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	countOfCards := make([]int, len(lines))

	for i := range countOfCards {
		countOfCards[i] = 1
	}

	for row, line := range lines {
		winningsCount := calcLinePoints(line)

		for i := row + 1; i <= row+winningsCount && i < len(countOfCards); i++ {
			countOfCards[i] += countOfCards[row]
		}
	}

	fmt.Println(countOfCards)

	total := 0
	for _, num := range countOfCards {
		total += num
	}

	fmt.Println(total)
}

func calcLinePoints(line string) int {
	numbersString := strings.Split(line, ":")[1]

	winningAndMyNumbersStrings := strings.Split(numbersString, "|")

	winningNumbers := strings.Fields(winningAndMyNumbersStrings[0])
	myNumbers := strings.Fields(winningAndMyNumbersStrings[1])

	count := 0
	for _, a := range winningNumbers {
		for _, b := range myNumbers {
			if a == b {
				count++
			}
		}
	}

	return count
}
