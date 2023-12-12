package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pointsSum := 0

	for scanner.Scan() {
		pointsSum += calcLinePoints(scanner.Text())
	}

	fmt.Println(pointsSum)
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

	return calcPoints(count)
}

func calcPoints(winCount int) int {
	if winCount < 2 {
		return winCount
	} else {
		result := 1
		for i := 2; i <= winCount; i++ {
			result *= 2
		}
		return result
	}
}
