package gomes

import (
	"testing"
)

func TestBoardStartsOutWithXAsTheNextPiece(t *testing.T) {
	game := TicTacToe()

	if game.nextPiece != xPiece {
		t.Errorf("Expecting `%v` as the next piece but found `%v`", xPiece, game.nextPiece)
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
