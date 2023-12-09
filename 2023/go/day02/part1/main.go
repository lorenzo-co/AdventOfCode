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
		idSum += idPossibleGame(scanner.Text())
	}

	fmt.Println(idSum)
}

func idPossibleGame(line string) int {
	trimmed := strings.TrimPrefix(line, "Game ")
	gameIdAndSets := strings.Split(trimmed, ":")

	idText, setsText := gameIdAndSets[0], gameIdAndSets[1]

	setsList := strings.Split(setsText, ";")

	for _, set := range setsList {
		cubeCounts := strings.Split(set, ",")

		for _, cubeCount := range cubeCounts {
			reNumber := regexp.MustCompile(`\d+`)

			number, err := strconv.Atoi(reNumber.FindString(cubeCount))

			if err != nil {
				panic(err)
			}

			if strings.HasSuffix(cubeCount, "red") && number > 12 {
				return 0
			}

			if strings.HasSuffix(cubeCount, "green") && number > 13 {
				return 0
			}

			if strings.HasSuffix(cubeCount, "blue") && number > 14 {
				return 0
			}

		}

	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		panic(err)
	}

	return id
}
