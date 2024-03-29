package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Game struct {
	id       int
	maxRed   int
	maxGreen int
	maxBlue  int
}

func main() {

	f, _ := os.ReadFile("./input.txt")
	gameStrings := strings.Split(string(f), "\n")

	//red := 12
	//green := 13
	//blue := 14

	total := 0

	for _, gameString := range gameStrings {
		currentGame := parseGameString(gameString)
		total += currentGame.maxBlue * currentGame.maxGreen * currentGame.maxRed
		//if currentGame.maxRed <= red && currentGame.maxBlue <= blue && currentGame.maxGreen <= green {
		//	total += currentGame.id
		//	fmt.Printf("Adding game id %d new total = %d\n", currentGame.id, total)
		//}

	}

	fmt.Printf("total = %d\n", total)

}

func parseGameString(gameString string) *Game {
	fmt.Println(gameString)
	game := new(Game)

	// get game id from string
	gameSplit := strings.Split(gameString, ":")
	game.id, _ = strconv.Atoi(strings.Split(gameSplit[0], " ")[1])

	// simplify string for parsing
	gameString = gameSplit[1]
	gameString = strings.ReplaceAll(gameString, "maxRed", "r")
	gameString = strings.ReplaceAll(gameString, "maxGreen", "g")
	gameString = strings.ReplaceAll(gameString, "maxBlue", "b")

	// parse string for number of marbles
	game.maxRed, game.maxGreen, game.maxBlue = 0, 0, 0
	current := 0
	for _, c := range gameString {
		if unicode.IsDigit(c) {
			v, _ := strconv.Atoi(string(c))
			if current > 0 {
				current = current * 10
			}
			current += v
		}
		if c == 'r' {
			if game.maxRed < current {
				game.maxRed = current
			}
			current = 0
		} else if c == 'g' {
			if game.maxGreen < current {
				game.maxGreen = current
			}
			current = 0
		} else if c == 'b' {
			if game.maxBlue < current {
				game.maxBlue = current
			}
			current = 0
		}
	}
	fmt.Printf("game %d : r = %d, g = %d, b = %d\n", game.id, game.maxRed, game.maxGreen, game.maxBlue)

	return game
}
