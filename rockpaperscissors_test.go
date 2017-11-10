package gomes

import (
	"testing"
)

func TestGameStartsOutWithNoWinner(t *testing.T) {
	game := RockPaperScissors()
	if game.winner != rpsNone {
		t.Errorf("Expecting no winner but got `%v`", game.winner)
	}
}

func TestGameStartsOutWithNoPlays(t *testing.T) {
	game := RockPaperScissors()
	if game.one != rpsEmpty || game.two != rpsEmpty {
		t.Error("Expecting no plays")
	}
}

func TestTiesWhenExpected(t *testing.T) {
	game := RockPaperScissors()
	game.PlayerOneTurn(game.Rock())
	game.PlayerTwoTurn(game.Rock())

	if game.winner != rpsTie {
		t.Error("Expecting a tie")
	}
}

func TestWinsWhenExpected(t *testing.T) {
	game := RockPaperScissors()
	game.PlayerOneTurn(game.Rock())
	game.PlayerTwoTurn(game.Scissors())

	if game.winner != rpsOne {
		t.Error("Expecting player 1 to win")
	}
}
