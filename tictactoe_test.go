package gomes

import (
	"testing"
)

func TestBoardStartsOutWithXAsTheNextPiece(t *testing.T) {
	game := TicTacToe()

	if game.next != xPiece {
		t.Errorf("Expecting `%v` as the next piece but found `%v`", xPiece, game.next)
	}
}

func TestBoardStartsOutWithNoWinner(t *testing.T) {
	game := TicTacToe()

	if game.winner != emptyPiece {
		t.Errorf("Expecting `%v` as the next piece but found `%v`", xPiece, game.next)
	}
}

func TestBoardStartsOutWithJustEmptyPieces(t *testing.T) {
	game := TicTacToe()

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if game.board[y][x] != emptyPiece {
				t.Errorf("Found non-empty piece on [%v, %v]", x, y)
			}
		}
	}
}

func TestWinningGameStateIsReturned(t *testing.T) {
	var done bool

	game := TicTacToe()

	if done, _ = game.Turn(0, 0); done != false {
		t.Error("Should return false")
	}
	if done, _ = game.Turn(1, 0); done != false {
		t.Error("Should return false")
	}
	if done, _ = game.Turn(2, 0); done != false {
		t.Error("Should return false")
	}
	if done, _ = game.Turn(1, 1); done != false {
		t.Error("Should return false")
	}
	if done, _ = game.Turn(2, 1); done != false {
		t.Error("Should return false")
	}
	if done, _ = game.Turn(1, 2); done != true {
		t.Error("Should return false")
	}
}

func TestWinnerIsSet(t *testing.T) {
	game := TicTacToe()

	game.Turn(0, 0)
	game.Turn(1, 0)
	game.Turn(2, 0)
	game.Turn(1, 1)
	game.Turn(2, 1)
	game.Turn(1, 2)

	if game.winner != oPiece {
		t.Error("Expected `o` to be the winner")
	}
}

func TestCannotTakeTurnsAfterGameIsDone(t *testing.T) {
	game := TicTacToe()

	game.Turn(0, 0)
	game.Turn(1, 0)
	game.Turn(2, 0)
	game.Turn(1, 1)
	game.Turn(2, 1)
	game.Turn(1, 2)
	_, err := game.Turn(2, 2)

	if err == nil {
		t.Error("Expected an error when taking a turn on a completed game")
	}
}
