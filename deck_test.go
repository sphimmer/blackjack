package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length to be 52 but got %d", len(cards))
	}

	if cards[0].Suit != "Dia" {
		t.Errorf("Expected first card to be 'Diamonds' but got %+v", cards[0])
	}

	if cards[len(cards)-1].Suit != "Spd" {
		t.Errorf("Expected last card to be 'Spades' but got %+v ", cards[len(cards)-1])
	}
}

func TestShuffle(t *testing.T) {
	cards := newDeck()
	firstCard := cards[0]
	tenthCard := cards[9]

	cards.shuffle()
	if firstCard == cards[0] {
		t.Errorf("Expected first card to have been shuffled")
	}

	if tenthCard == cards[9] {
		t.Errorf("Expected tenth card to have been shuffled")
	}
}

func TestDeal(t *testing.T) {
	cards := newDeck()
	originalLength := len(cards)

	hands := cards.deal(3, 10)
	if len(hands) != 3 {
		t.Errorf("Expected 3 hands but got %v", len(hands))
	}
	if len(cards) != originalLength-len(hands)*10 {
		t.Errorf("Expected Deck size to be %v but got %v", originalLength-len(hands)*10, len(cards))
	}

	for _, h := range hands {
		if len(h) != 10 {
			t.Errorf("hand did not get properly dealt, missing cards: %+v", h)
		}
	}

}

func TestHit(t *testing.T) {
	hand := deck{card{Value: "5"}, card{Value: "7", Suit: "Spades"}}
	cards := newDeck()
	cards.hit(&hand, 2)
	if len(hand) != 4 {
		t.Errorf("Expected to have 4 cards but got %v", len(hand))
	}
}
