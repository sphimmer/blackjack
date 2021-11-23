package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type card struct {
	Value      string
	Suit       string
	PointValue int
}

func getCardValueMap() map[string]int {
	return map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "Jack": 10, "Queen": 10, "King": 10, "Ace": 11}
}

func (c card) print() {
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
	fmt.Fprintln(w, "| "+c.Value+"\t|")
	fmt.Fprintln(w, "| "+c.Suit+"\t|")
	w.Flush()
}

func getCardSuits() []string {
	return []string{"Dia", "Clb", "Hrt", "Spd"}
}
