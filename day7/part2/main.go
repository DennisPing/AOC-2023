package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

var CardValues = map[byte]int{
	'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'T': 10, 'Q': 11, 'K': 12, 'A': 13,
}

type Hand struct {
	Cards string
	Bid   int
	Rank  int
	Type  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hands []*Hand
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])

		hand := &Hand{
			Cards: parts[0],
			Bid:   bid,
		}
		hands = append(hands, hand)
	}

	for _, hand := range hands {
		classifyHand(hand)
	}

	// Sort hands in descending order based on their strength
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[j], hands[i])
	})

	for _, h := range hands {
		fmt.Println(*h)
	}

	totalWinnings := calculateWinnings(hands)
	fmt.Println(totalWinnings)
}

func classifyHand(hand *Hand) {
	jokers := 0
	for _, card := range hand.Cards {
		if card == 'J' {
			jokers++
		}
	}

	if jokers == 5 { // Edge case
		hand.Type = FiveKind
	} else if jokers > 0 {
		hand.Type = substituteClassify(hand.Cards)
	} else {
		hand.Type = simpleClassify(hand.Cards)
	}
}

func substituteClassify(cards string) int {
	// Generate combinations by replacing J
	strongestType := HighCard
	for _, card := range cards { // Simple brute force
		if card != 'J' {
			newHand := strings.Replace(cards, "J", string(card), -1)
			handType := simpleClassify(newHand)
			if handType > strongestType {
				strongestType = handType
			}
		}
	}
	return strongestType
}
func simpleClassify(cards string) int {
	cardCount := make(map[rune]int)
	var maxCount int

	// Count letter/number frequency
	for _, card := range cards {
		cardCount[card]++
		if cardCount[card] > maxCount {
			maxCount = cardCount[card]
		}
	}

	switch {
	case len(cardCount) == 1:
		return FiveKind
	case len(cardCount) == 2:
		if maxCount == 4 {
			return FourKind
		} else {
			return FullHouse
		}
	case len(cardCount) == 3:
		if maxCount == 3 {
			return ThreeKind
		} else {
			return TwoPair
		}
	case len(cardCount) == 4:
		return OnePair
	default:
		return HighCard
	}
}

func compareHands(h1, h2 *Hand) bool {
	// Compare the card type first
	if h1.Type != h2.Type {
		return h1.Type > h2.Type
	}

	// If types are the same, compare based on the strength of individual cards
	for i := 0; i < len(h1.Cards); i++ {
		if CardValues[h1.Cards[i]] != CardValues[h2.Cards[i]] {
			return CardValues[h1.Cards[i]] > CardValues[h2.Cards[i]]
		}
	}

	return false // Hands are identical
}

func calculateWinnings(hands []*Hand) int {
	totalWinnings := 0
	for i, hand := range hands {
		rank := i + 1
		totalWinnings += rank * hand.Bid
	}
	return totalWinnings
}
