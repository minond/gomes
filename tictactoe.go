package gomes

/**
 * game := TicTacToe()
 *
 * game.Turn(0, 0)
 * game.Turn(1, 0)
 * game.Turn(2, 0)
 * game.Turn(1, 1)
 * game.Turn(2, 1)
 * game.Turn(1, 2)
 *
 * fmt.Println(game)
 */
import (
	"fmt"
	"strings"
)

type ticTacToePiece string

type ticTacToeGameState struct {
	board     [3][3]ticTacToePiece `json:"board"`
	nextPiece ticTacToePiece       `json:"next"`
}

const (
	emptyPiece ticTacToePiece = " "
	oPiece     ticTacToePiece = "o"
	xPiece     ticTacToePiece = "x"
)

const gameTmpl = `
-------------
| %v | %v | %v |
-------------
| %v | %v | %v |
-------------
| %v | %v | %v |
-------------`

var oppPiece = map[ticTacToePiece]ticTacToePiece{
	xPiece: oPiece,
	oPiece: xPiece,
}

func TicTacToe() (state ticTacToeGameState) {
	state.board[0][0] = emptyPiece
	state.board[0][1] = emptyPiece
	state.board[0][2] = emptyPiece
	state.board[1][0] = emptyPiece
	state.board[1][1] = emptyPiece
	state.board[1][2] = emptyPiece
	state.board[2][0] = emptyPiece
	state.board[2][1] = emptyPiece
	state.board[2][2] = emptyPiece

	state.nextPiece = xPiece

	return
}

func (state *ticTacToeGameState) Turn(x, y int) error {
	if state.board[y][x] != emptyPiece {
		return fmt.Errorf("Tic Tac Toe board already has a piece on [%v, %v] coordinates", x, y)
	}

	state.board[y][x] = state.nextPiece
	state.nextPiece = oppPiece[state.nextPiece]

	return nil
}

func (state ticTacToeGameState) String() string {
	return strings.TrimSpace(fmt.Sprintf(gameTmpl,
		state.board[0][0], state.board[0][1], state.board[0][2],
		state.board[1][0], state.board[1][1], state.board[1][2],
		state.board[2][0], state.board[2][1], state.board[2][2],
	))
}
