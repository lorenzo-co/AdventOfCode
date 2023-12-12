package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Point int
	Cards string
	Bid   int
}

var cardValues = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

const (
	HighCard int = iota
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	content, err := os.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	hands := make([]Hand, len(lines))

	for index, line := range lines {
		parts := strings.Split(line, " ")
		cards := parts[0]
		bid, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		cardCount := countOccurrences(cards)

		point := getHandPoint(cardCount)

		hands[index] = Hand{point, cards, bid}
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Point != hands[j].Point {
			return hands[i].Point < hands[j].Point
		}

		for k := 0; k < 5; k++ {
			if cardValues[rune(hands[i].Cards[k])] != cardValues[rune(hands[j].Cards[k])] {
				return cardValues[rune(hands[i].Cards[k])] < cardValues[rune(hands[j].Cards[k])]
			}
		}

		return false
	})

	sum := 0
	for index, hand := range hands {
		fmt.Println(hand.Cards, hand.Point, index+1, hand.Bid)
		sum += hand.Bid * (index + 1)
	}

	fmt.Println(sum)
}

func countOccurrences(s string) []int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}

	var countsSlice []int
	for _, count := range counts {
		countsSlice = append(countsSlice, count)
	}

	sort.Ints(countsSlice)
	return countsSlice
}

func compareHands(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func getHandPoint(cardCount []int) int {
	switch {
	case compareHands(cardCount, []int{5}):
		return FiveOfAKind
	case compareHands(cardCount, []int{1, 4}):
		return FourOfAKind
	case compareHands(cardCount, []int{2, 3}):
		return FullHouse
	case compareHands(cardCount, []int{1, 1, 3}):
		return ThreeOfAKind
	case compareHands(cardCount, []int{1, 2, 2}):
		return TwoPair
	case compareHands(cardCount, []int{1, 1, 1, 2}):
		return Pair
	case compareHands(cardCount, []int{1, 1, 1, 1, 1}):
		return HighCard
	default:
		panic(fmt.Sprintf("Unexpected card counts: %v", cardCount))
	}

}
