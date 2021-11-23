package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("BlackJack!!! You VS the dealer")
	printGameRules()
	gs := gameSetup()
	fmt.Println("Your hand has been Dealt!")
	gameLoop(gs)
	fmt.Println("Thanks for playing!")
}

func printGameRules() {
	fmt.Println("****Game Rules****")
	fmt.Println("Closes to a score of 21 wins")
	fmt.Println("Numbered cards are worth their number, royals are worth 10 points, Aces are 11 points or 1 point if over 21 points in total score")
	fmt.Println("Tie goes to the Dealer")
	fmt.Println("*****Beginning game now*****")
}

func gameSetup() gameState {
	gs := gameState{mainDeck: newDeck(), player: player{}, dealer: dealer{}, turn: 0}
	gs.mainDeck.shuffle()
	gs.mainDeck.deal([]*deck{&gs.dealer.hand, &gs.player.hand}, 2)
	gs.dealer.calculateScore()
	gs.player.calculateScore()
	return gs
}

func gameLoop(gs gameState) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n________________")
		fmt.Println("Dealer shows 1 card")
		gs.dealer.showHand(1)
		fmt.Printf("\nYour hand is as stands:\n Points: %v \n", gs.player.score)
		gs.player.hand.print()
		fmt.Print("Do you wish to take a card? (y/n)")
		answer, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("error reading your response")
			fmt.Println(err)
			continue
		}
		if answer == 'y' {
			gs.mainDeck.hit(&gs.player.hand, 1)
			gs.player.calculateScore()
		}
		if gs.dealer.score < 15 {
			fmt.Println("Dealer takes a card")
			gs.mainDeck.hit(&gs.dealer.hand, 1)
			gs.dealer.calculateScore()
		}
		if gs.player.score > 21 {
			printEndGameStatus(gs)
			break
		}
		if answer == 'n' {
			for gs.dealer.score < 17 {
				gs.mainDeck.hit(&gs.dealer.hand, 1)
				gs.dealer.calculateScore()
			}
			printEndGameStatus(gs)
			break
		}
	}
}

func printEndGameStatus(gs gameState) {
	winner := gs.calculateWinner()
	fmt.Printf("%v wins!\n", winner)
	fmt.Printf("Dealer's Score: %v \nDealer's Cards:\n", gs.dealer.score)
	gs.dealer.hand.print()
	fmt.Printf("Your Score: %v \nYour Cards:\n", gs.player.score)
	gs.player.hand.print()
}
