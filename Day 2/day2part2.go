package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Opponent 			(A, B, C) = (Rock, Paper, Scissors)
//Wanted Outcome 	(X, Y, Z) = (Lose, Draw, Win)
//Score hands 		(1, 2 , 3) = (Rock, Paper, Scissors)
//Score outcome 	(0, 3, 6) = (Lose, Draw, Win)

var opponentHand string
var myHand string
var winOutcomes = [3][2]string{{"A", "C"}, {"B", "A"}, {"C", "B"}} //Determines all winning outcomes

func main() {
	strategyGuide := importStrategyGuide()
	totalScore := playAllGames(strategyGuide)
	fmt.Println(totalScore)
}

func determineWinner() int {
	hands := [2]string{myHand, opponentHand}
	switch {
	case myHand == opponentHand:
		// It's a draw
		return 3
	case hands == winOutcomes[0] || hands == winOutcomes[1] || hands == winOutcomes[2]:
		// I win
		return 6
	default:
		// Opponent win
		return 0
	}
}

func importStrategyGuide() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	list := []string{}

	for scanner.Scan() { // internally, it advances token based on separator
		list = append(list, scanner.Text())
	}

	return list
}

func playAllGames(list []string) int {
	score := 0
	for i := 0; i < len(list); i++ {
		setHands(list[i])
		outcomeScore := determineWinner()
		score += calcScore(outcomeScore)
	}

	return score
}

func setHands(hands string) {
	opponentHand = string(hands[0])
	wantedOutcome := string(hands[2])

	switch {
	case wantedOutcome == "X":
		//Lose
		if opponentHand == "A" {
			myHand = "C"
		} else if opponentHand == "B" {
			myHand = "A"
		} else {
			myHand = "B"
		}
	case wantedOutcome == "Y":
		//Draw
		if opponentHand == "A" {
			myHand = "A"
		} else if opponentHand == "B" {
			myHand = "B"
		} else {
			myHand = "C"
		}
	case wantedOutcome == "Z":
		//Win
		if opponentHand == "A" {
			myHand = "B"
		} else if opponentHand == "B" {
			myHand = "C"
		} else {
			myHand = "A"
		}
	}
}

func calcScore(outcomeScore int) int {
	score := outcomeScore
	switch {
	case myHand == "A":
		score += 1
	case myHand == "B":
		score += 2
	case myHand == "C":
		score += 3
	}

	return score
}
