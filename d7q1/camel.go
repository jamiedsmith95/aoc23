package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CardGame struct {
	line     string
	hand     string
	handType int
	bid      int
	rank     int
	score    int
	maximum  int
	second   int
}

func (cg *CardGame) getScore() {
	cg.score = cg.bid * cg.rank
}

var games []*CardGame

func (cg *CardGame) compareGames(og *CardGame) bool {
	if cg.handType > og.handType {
		return true
	} else if cg.handType < og.handType {
		return false
	}
	currentString := cg.hand
	otherString := og.hand
	for i := 0; i < 5; i++ {

		if currentString[i] > otherString[i] {
			return true

		} else if otherString[i] > currentString[i] {
			return false
		}
	}
	return true
}

func (cg *CardGame) getHandType() {
	hand := cg.hand
	var matches []int
	for i := 0; i < len(hand); i++ {
		matches = append(matches, 0)
	}
	var jokers int
	for i, char := range hand {
		if char == '0' {
			continue
		}
		if char == '1' {
			jokers++
			continue
		}
		for j, rest := range hand {
			if rest == char {
				matches[i]++
				hand = hand[:j] + string('0') + hand[j+1:]
				if len(hand) != 5 {
					fmt.Println(hand)
				}

			}
		}
	}
	var maximum int
	var second int
	for _, num := range matches {
		if num > maximum {
			second = maximum
			maximum = num
		} else if num >= second {
			second = num
		}
	}
	fmt.Println(matches, " ", maximum)
	cg.second = second
	cg.maximum = maximum
	if maximum+jokers == 5 {
		cg.handType = 6
	} else if maximum+jokers == 4 {
		cg.handType = 5
	} else if (maximum == 3 && second == 2) || (jokers+maximum == 3 && second == 2) || (maximum == 3 && second+jokers == 2) {
		cg.handType = 4
	} else if maximum+jokers == 3 {
		cg.handType = 3
	} else if (maximum == 2 && second == 2) || (maximum == 2 && jokers+second == 2) {
		cg.handType = 2
	} else if maximum+jokers == 2 {
		cg.handType = 1
	} else {
		cg.handType = 0
	}

}

func getGame(line string) *CardGame {
	strings := strings.Split(line, " ")
	oldHand := strings[0]
	var hand string

	for _, r := range oldHand {
		if r == 'A' {
			r = 'e'
		}
		if r == 'K' {
			r = 'd'
		}
		if r == 'Q' {
			r = 'c'
		}
		if r == 'J' {
			r = '1'
		}
		if r == 'T' {
			r = 'a'
		}
		hand = hand + string(r)
	}
	bid, err := strconv.Atoi(strings[1])
	if err != nil {
		fmt.Println(err)
	}
	cardGame := CardGame{line: line, hand: hand, bid: bid}
	return &cardGame
}

func (cg *CardGame) setRank(rank int) {
	cg.rank = rank
}

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()

	for _, line := range lines {
		games = append(games, getGame(line))
	}
	for _, game := range games {
		game.getHandType()
	}
	// var root node
	// root = node {Game: &games[0]}
	// for i,game := range games[1:] {
	//   if i == 0 {
	//     continue
	//   }
	// }

	for _, current := range games {
		count := 0
		for _, other := range games {
			if current.compareGames(other) {
				count++
			}
		}
		current.setRank(count)
		count = 0
	}
	var sumScore int
	var sumRank int
	for _, game := range games {
		game.getScore()
		sumRank += game.rank
		sumScore += game.score
	}
	for _, game := range games {
		// fmt.Println(game.hand, " ", game.handType, " ", game.rank)
		for _, other := range games {
			if game.rank == other.rank && game.bid != other.bid {
				fmt.Println("same rank")
			}
		}
	}
	var order []CardGame
	for _, game := range games {
		order = append(order, *game)
	}
	for _, game := range games {
		order[game.rank-1] = *game
	}
	for _, game := range order {
		fmt.Println(game.handType, game.hand, game.bid)
	}
	fmt.Println(sumScore)
	fmt.Println(sumRank)
}
