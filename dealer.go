package main

type dealer player

func (d dealer) showHand(numCards int) {
	cardsToShow := d.hand[:numCards]
	for _, card := range cardsToShow {
		card.print()
	}
}

func (p *dealer) calculateScore() {
	var score int

	for _, card := range p.hand {
		if card.Value == "Ace" && score+card.PointValue > 21 {
			score += 1
		} else {
			score += card.PointValue
		}
	}
	p.score = score
}
