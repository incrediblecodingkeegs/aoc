package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ScratchCard struct {
	id             int
	winningNumbers []int
	gameNumbers    map[int]bool
}

type ScratchCardNode struct {
	timesWon int
	card     *ScratchCard
	next     *ScratchCardNode
}

func main() {
	start := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)
	// initialise linked list
	scanner.Scan()
	root := &ScratchCardNode{
		timesWon: 0,
		card:     createCard(scanner.Text()),
		next:     nil,
	}

	node := root

	for scanner.Scan() {
		node.next = &ScratchCardNode{
			timesWon: 0,
			card:     createCard(scanner.Text()),
			next:     nil,
		}
		node = node.next
	}

	// iterate through list and update each cards times won
	result := processCardList(root, 0)
	//result := 0
	//node = root
	//for node != nil {
	//	result += node.timesWon
	//	node = node.next
	//}

	log.Printf("result = %d\n", result)
	fmt.Printf("Elapsed time : %s\n", time.Since(start))
}

func createCard(cardString string) *ScratchCard {
	card := new(ScratchCard)
	// split string into winning numbers + game numbers
	cardStringSplit := strings.Split(cardString, "|")
	winningNumbersString := cardStringSplit[0]
	gameNumbersString := cardStringSplit[1]

	// get card id from winning numbers string
	winningNumbersSplit := strings.Split(winningNumbersString, ":")
	winningNumbersString = winningNumbersSplit[1]
	card.id, _ = strconv.Atoi(strings.Fields(winningNumbersSplit[0])[1])

	// parse winningNumbersString for winningNumbers
	winningNumbersSplit = strings.Fields(winningNumbersString)
	for _, winningNumberString := range winningNumbersSplit {
		winningNumber, _ := strconv.Atoi(winningNumberString)
		card.winningNumbers = append(card.winningNumbers, winningNumber)
	}

	// parse gameNumbersString for gameNumbers
	gameNumbersSplit := strings.Fields(gameNumbersString)
	card.gameNumbers = make(map[int]bool)
	for _, gameNumbersString := range gameNumbersSplit {
		gameNumber, _ := strconv.Atoi(gameNumbersString)
		card.gameNumbers[gameNumber] = true
	}

	return card
}

// single-linked node list
// if card wins -> iterate through list, += node value by times that card has been won
// move to next card

func processCardList(root *ScratchCardNode, cardTotal int) int {
	tmp := root.next
	cardWinnings := getCardWinnings(root.card)
	for i := 0; i < cardWinnings; i++ {
		if tmp != nil {
			tmp.timesWon += root.timesWon + 1
		}
		if tmp.next != nil {
			tmp = tmp.next
		} else {
			break
		}
	}
	//fmt.Printf("id = %d cardWinnings = %d timesWon = %d cardTotal = %d\n", root.card.id, cardWinnings, root.timesWon, cardTotal)
	cardTotal += root.timesWon
	if root.next != nil {
		return processCardList(root.next, cardTotal)
	} else {
		return cardTotal + root.card.id
	}
}

func getCardWinnings(card *ScratchCard) int {
	winnings := 0
	// check if winning number exists in game numbers
	for _, number := range card.winningNumbers {
		if ok, _ := card.gameNumbers[number]; ok {
			winnings += 1
		}
	}
	return winnings
}
