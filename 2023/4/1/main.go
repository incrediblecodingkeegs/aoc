package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id             int
	winningNumbers []int
	gameNumbers    map[int]bool
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		card := createCard(scanner.Text())
		result += getCardWinnings(card)
	}

	log.Printf("result = %d\n", result)
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

func getCardWinnings(card *ScratchCard) int {
	winnings := 0
	// check if winning number exists in game numbers
	for _, number := range card.winningNumbers {
		if ok, _ := card.gameNumbers[number]; ok {
			if winnings != 0 {
				winnings *= 2
			} else {
				winnings = 1
			}
		}

	}
	return winnings
}
