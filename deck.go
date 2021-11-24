package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// create a new type of deck
// which is a slice of strings
type deck []card

func newDeck() deck {
	cards := deck{}
	for _, suit := range getCardSuits() {
		for key, value := range getCardValueMap() {
			c := card{Value: key, Suit: suit, PointValue: value}
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		card.print()
		fmt.Println()
	}
}

func (d *deck) deal(numberOfHands int, numberOfCards int) []deck {
	dp := *d
	cardsToBeDealt := dp[:numberOfCards*numberOfHands]
	*d = dp[numberOfCards*numberOfHands:]
	hands := []deck{}
	for hi := 0; hi < numberOfHands; hi++ {
		hands = append(hands, deck{})
		for i := hi; i < len(cardsToBeDealt); i += numberOfHands {
			hands[hi] = append(hands[hi], cardsToBeDealt[i])
		}
	}
	return hands
}

func (d *deck) hit(hand *deck, numberOfCards int) {
	dp := *d
	hp := *hand
	cardsToBeDealt := dp[:numberOfCards]
	*d = dp[numberOfCards:]
	*hand = append(hp, cardsToBeDealt...)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}

func (d deck) toJson() ([]byte, error) {
	jsonDeck, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonDeck, nil
}

func (d deck) saveToFile(filename string) error {
	jsonString, err := d.toJson()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return ioutil.WriteFile(filename, jsonString, 0777)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	return deckFromJson(bs)
}

func deckFromJson(jsonString []byte) deck {
	nd := deck{}
	json.Unmarshal(jsonString, &nd)
	return nd
}
