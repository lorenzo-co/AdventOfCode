package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	idSum := 0

	for scanner.Scan() {
		idSum += getMinCubesNumber(scanner.Text())
	}

	fmt.Println(idSum)
}

func getMinCubesNumber(line string) int {
	trimmed := strings.TrimPrefix(line, "Game ")
	gameIdAndSets := strings.Split(trimmed, ":")

	setsText := gameIdAndSets[1]

	setsList := strings.Split(setsText, ";")

	fmt.Println(line)

	reds, greens, blues := 0, 0, 0

	for _, set := range setsList {
		cubeCounts := strings.Split(set, ",")

		for _, cubeCount := range cubeCounts {
			reNumber := regexp.MustCompile(`\d+`)

			number, err := strconv.Atoi(reNumber.FindString(cubeCount))

			if err != nil {
				panic(err)
			}

			if strings.HasSuffix(cubeCount, "red") {
				reds = max(reds, number)
			}

			if strings.HasSuffix(cubeCount, "green") {
				greens = max(greens, number)
			}

			if strings.HasSuffix(cubeCount, "blue") {
				blues = max(blues, number)
			}
		}
	}

	fmt.Println(reds, greens, blues)

	return reds * greens * blues
}
