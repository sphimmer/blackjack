package main

type player struct {
	hand  deck
	score int
}

func (p *player) calculateScore() {
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
