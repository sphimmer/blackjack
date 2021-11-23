package main

import "testing"

func TestCalculateWinner(t *testing.T) {
	gs := gameState{dealer: player{score: 21}, player: player{score: 21}}
	winner := gs.calculateWinner()

	if winner == "Player" {
		t.Errorf("Ties should go to the dealer")
	}

	gs = gameState{dealer: player{score: 16}, player: player{score: 22}}
	winner = gs.calculateWinner()
	if winner == "Player" {
		t.Errorf("Scores over 21 should not win")
	}

	gs = gameState{dealer: player{score: 16}, player: player{score: 20}}
	winner = gs.calculateWinner()
	if winner == "Dealer" {
		t.Errorf("Closest Score should win")
	}
	gs = gameState{dealer: player{score: 23}, player: player{score: 15}}
	winner = gs.calculateWinner()
	if winner == "Dealer" {
		t.Errorf("Closest Score should win")
	}

	// natural win
	gs = gameState{dealer: player{score: 20, hand: deck{card{}, card{}}}, player: player{score: 21, hand: deck{card{}, card{}}}}
	winner = gs.calculateWinner()
	if winner == "Dealer" {
		t.Errorf("Natural Win goes to player")
	}
}
