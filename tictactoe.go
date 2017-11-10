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
	"errors"
	"fmt"
	"strings"
)

type tttPiece string
type tttBoard [3][3]tttPiece

type tttGameState struct {
	board  tttBoard `json:"board"`
	next   tttPiece `json:"next"`
	winner tttPiece `json:"winner"`
}

const (
	tttEmptyPiece tttPiece = " "
	tttOPiece     tttPiece = "o"
	tttXPiece     tttPiece = "x"

	tttTmpl = `
		-------------
		| %v | %v | %v |
		-------------
		| %v | %v | %v |
		-------------
		| %v | %v | %v |
		-------------`
)

var (
	tttNextPiece = map[tttPiece]tttPiece{
		tttXPiece: tttOPiece,
		tttOPiece: tttXPiece,
	}

	tttWinningCoordinates = [][][2]int{
		[][2]int{[2]int{0, 0}, [2]int{1, 0}, [2]int{2, 0}}, // Row 1
		[][2]int{[2]int{0, 1}, [2]int{1, 1}, [2]int{2, 1}}, // Row 2
		[][2]int{[2]int{0, 2}, [2]int{1, 2}, [2]int{2, 2}}, // Row 3
		[][2]int{[2]int{0, 0}, [2]int{0, 1}, [2]int{0, 2}}, // Column 1
		[][2]int{[2]int{1, 0}, [2]int{1, 1}, [2]int{1, 2}}, // Column 2
		[][2]int{[2]int{2, 0}, [2]int{2, 1}, [2]int{2, 2}}, // Column 3
		[][2]int{[2]int{0, 0}, [2]int{1, 1}, [2]int{2, 2}}, // Top left to bottom right
		[][2]int{[2]int{2, 0}, [2]int{1, 1}, [2]int{0, 2}}, // Top right to bottom left
	}
)

func TicTacToe() (game tttGameState) {
	game.board[0][0] = tttEmptyPiece
	game.board[0][1] = tttEmptyPiece
	game.board[0][2] = tttEmptyPiece
	game.board[1][0] = tttEmptyPiece
	game.board[1][1] = tttEmptyPiece
	game.board[1][2] = tttEmptyPiece
	game.board[2][0] = tttEmptyPiece
	game.board[2][1] = tttEmptyPiece
	game.board[2][2] = tttEmptyPiece

	game.next = tttXPiece
	game.winner = tttEmptyPiece

	return
}

func (game *tttGameState) Turn(x, y int) (bool, error) {
	if game.winner != tttEmptyPiece {
		return false, errors.New("There's already a winner!")
	}

	if game.board[y][x] != tttEmptyPiece {
		return false, fmt.Errorf("Tic Tac Toe board already has a piece on [%v, %v] coordinates", x, y)
	}

	game.board[y][x] = game.next

	// Winner??
	for _, coors := range tttWinningCoordinates {
		if game.board.samePieceInCoors(game.next, coors...) {
			game.winner = game.next
			return true, nil
		}
	}

	game.next = tttNextPiece[game.next]

	return false, nil
}

func (game tttGameState) String() string {
	return strings.TrimSpace(fmt.Sprintf(tttTmpl,
		game.board[0][0], game.board[0][1], game.board[0][2],
		game.board[1][0], game.board[1][1], game.board[1][2],
		game.board[2][0], game.board[2][1], game.board[2][2],
	))
}

func (board tttBoard) samePieceInCoors(p tttPiece, coors ...[2]int) bool {
	for _, coor := range coors {
		if board[coor[0]][coor[1]] != p {
			return false
		}
	}

	return true
}
