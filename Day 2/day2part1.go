package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Opponent 	(A, B, C) = (Rock, Paper, Scissors)
//Me 		(X, Y, Z) = (Rock, Paper, Scissors)
//Score hands (1, 2 , 3) = (Rock, Paper, Scissors)
//Score outcome (0, 3, 6) = (Lose, Draw, Win)

var opponentHand string
var myHand string
var winOutcomes = [3][2]string{{"X", "C"}, {"Y", "A"}, {"Z", "B"}}  //Determines all winning outcomes
var drawOutcomes = [3][2]string{{"X", "A"}, {"Y", "B"}, {"Z", "C"}} //Determines all drawing outcomes, to skip conversion from a letter to a hand type

func main() {
	strategyGuide := importStrategyGuide()
	totalScore := playAllGames(strategyGuide)
	fmt.Println(totalScore)
}

func determineWinner() int {
	hands := [2]string{myHand, opponentHand}
	switch {
	case hands == drawOutcomes[0] || hands == drawOutcomes[1] || hands == drawOutcomes[2]:
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
	myHand = string(hands[2])
	opponentHand = string(hands[0])
}

func calcScore(outcomeScore int) int {
	score := outcomeScore
	switch {
	case myHand == "X":
		score += 1
	case myHand == "Y":
		score += 2
	case myHand == "Z":
		score += 3
	}

	return score
}
