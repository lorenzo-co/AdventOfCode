package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	time := strings.Replace(strings.Split(lines[0], ":")[1], " ", "", -1)
	distance := strings.Replace(strings.Split(lines[1], ":")[1], " ", "", -1)

	product := calcWays(time, distance)

	fmt.Println(product)
}

func calcWays(timeString string, distanceString string) int {
	time, errTime := strconv.Atoi(timeString)
	distance, errDist := strconv.Atoi(distanceString)

	if errTime != nil {
		panic(errTime)
	}
	if errDist != nil {
		panic(errDist)
	}

	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			return time - 2*i + 1
		}
	}

	return 1
}
