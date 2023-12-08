package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		sum += firstAndLastDigitsToInt(scanner.Text())
	}

	fmt.Println(sum)
}

func firstAndLastDigitsToInt(line string) int {
	firstAndLastDigits := ""

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
