package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func NewHand(handString string) *Hand {
	if handString == "" {
		return nil
	}

	handFields := strings.Fields(handString)
	cards := handFields[0]
	bid, _ := strconv.Atoi(handFields[1])
	return &Hand{cards: cards, bid: bid}
}

func (h *Hand) String() string {
	return fmt.Sprintf("cards: %s bid: %d\n", h.cards, h.bid)
}

func (h *Hand) getHandStrength() int {
	// Check how many times each card occurs in hand
	cardCount := make(map[rune]int)
	for _, c := range h.cards {
		cardCount[c] += 1
	}

	// Get num pairs and highest cardCount
	numPairs := 0
	highCardCount := 0
	for _, v := range cardCount {
		if v >= highCardCount {
			highCardCount = v
		}
		if v >= 2 {
			numPairs += 1
		}
	}

	// Based on highCardCount, determine rank
	switch highCardCount {
	case 5: // 5 of a kind
		return 6
	case 4: // 4 of a kind
		return 5
	case 3:
		if numPairs == 2 { // full house
			return 4
		} else { // 3 of a kind
			return 3
		}
	case 2:
		if numPairs == 2 { // two pair
			return 2
		} else { // one pair
			return 1
		}
	default:
		return 0 // high card
	}
}

// returns true if h1 high card > h2 high card
func compareHandHighCard(h1 *Hand, h2 *Hand) bool {
	cardValue := map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	for i := 0; i < len(h1.cards); i++ {
		if cardValue[h1.cards[i]] > cardValue[h2.cards[i]] {
			return true
		} else if cardValue[h1.cards[i]] < cardValue[h2.cards[i]] {
			return false
		}
	}
	return false
}

func part1(handStrings []string) int {
	// parse cards and append to correct strength
	var handStrengths [7][]*Hand
	for _, handString := range handStrings {
		hand := NewHand(handString)
		strength := hand.getHandStrength()
		handStrengths[strength] = append(handStrengths[strength], hand)
	}

	//for i, strength := range handStrengths {
	//	fmt.Printf("handStrengths[%d] = %s\n", i, strength)
	//}

	// sort each strength category, and get total rank
	rank := 0
	for i := 0; i < len(handStrengths); i++ {
		rank += len(handStrengths[i])
		sort.Slice(handStrengths[i], func(j, k int) bool {
			return compareHandHighCard(handStrengths[i][j], handStrengths[i][k])
		})
	}

	for i, strength := range handStrengths {
		fmt.Printf("handStrengths[%d] = %s\n", i, strength)
	}

	// get winnings by iterating from bottom up through ranks
	winnings := 0
	for i := len(handStrengths) - 1; i >= 0; i-- {
		for _, hand := range handStrengths[i] {
			fmt.Printf("bid %d rank %d\n", hand.bid, rank)
			winnings += hand.bid * rank
			rank -= 1
		}
	}
	return winnings

}

func main() {
	f, _ := os.ReadFile("input.txt")
	fSplit := strings.Split(string(f), "\n")

	fmt.Printf("Total winnings = %d\n", part1(fSplit))
}
