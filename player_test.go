package main

import "testing"

func TestCalculateScore(t *testing.T) {

	player1 := player{hand: []card{{Value: "Ace", PointValue: 11}, {Value: "Ace", PointValue: 11}}}
	player1.calculateScore()
	if player1.score != 12 {
		t.Errorf("Expected score to be 12 but got %v", player1.score)
	}

	player2 := player{hand: []card{{Value: "King", PointValue: 10}, {Value: "Ace", PointValue: 11}}}
	player2.calculateScore()
	if player2.score != 21 {
		t.Errorf("Expected score to be 21 but got %v", player2.score)
	}

	player3 := player{hand: []card{{Value: "2", PointValue: 2}, {Value: "5", PointValue: 5}, {Value: "Jack", PointValue: 10}, {Value: "Ace", PointValue: 11}}}
	player3.calculateScore()
	if player3.score != 18 {
		t.Errorf("Expected score to be 18 but got %v", player3.score)
	}

	player4 := player{hand: []card{{Value: "Ace", PointValue: 11}, {Value: "Ace", PointValue: 11}, {Value: "Ace", PointValue: 11}, {Value: "Ace", PointValue: 11}}}
	player4.calculateScore()
	if player4.score != 14 {
		t.Errorf("Expected score to be 12 but got %v", player4.score)
	}
}
