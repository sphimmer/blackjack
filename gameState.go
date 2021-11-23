package main

type gameState struct {
	mainDeck deck
	dealer   player
	player   player
}

func (gs gameState) calculateWinner() string {
	// natural win or Black Jack
	if len(gs.player.hand) == 2 && gs.player.score == 21 && len(gs.dealer.hand) != 2 && gs.dealer.score != 21 {
		return "Player"
	}

	if gs.dealer.score <= 21 && (gs.player.score > 21 || 21-gs.dealer.score <= 21-gs.player.score) {
		return "Dealer"
	} else {
		return "Player"
	}
}
