package main

import (
	"errors"
	"fmt"
)

type rpsPlayer int
type rpsOption int

type rpsGameState struct {
	winner rpsPlayer
	turns  []map[rpsPlayer]rpsOption
}

const (
	rpsNone rpsPlayer = -1
	rpsOne  rpsPlayer = 0
	rpsTwo  rpsPlayer = 1

	rpsRock     rpsOption = 0
	rpsPaper    rpsOption = 1
	rpsScissors rpsOption = 2
)

var (
	rpsOtherPlayer = map[rpsPlayer]rpsPlayer{
		rpsOne: rpsTwo,
		rpsTwo: rpsOne,
	}
)

func RockPaperScissors() (game rpsGameState) {
	game.winner = rpsNone
	game.turns = make([]map[rpsPlayer]rpsOption, 0)
	game.turns = append(game.turns, map[rpsPlayer]rpsOption{})
	return game
}

func (game *rpsGameState) Turn(player rpsPlayer, option rpsOption) (bool, error) {
	if player == rpsNone {
		return false, errors.New("Invalid player")
	}

	// if len(game.turns) == 0 {
	// 	game.turns = make([]map[rpsPlayer]rpsOption, 0)
	// 	game.turns = append(game.turns, map[rpsPlayer]rpsOption{player: option})
	// } else {
	// 	game.turns
	// }

	curr := game.turns[len(game.turns)-1]
	turn, set := curr[player]

	if set == true {
		return false, errors.New("Player has already taken a turn")
	}

	curr[player] = option

	// Has the other player taken this turn?
	otherOption, set := curr[rpsOtherPlayer[player]]

	fmt.Println(len(game.turns))

	return false, nil
}

func (game rpsGameState) Rock() rpsOption {
	return rpsRock
}

func (game rpsGameState) Paper() rpsOption {
	return rpsPaper
}

func (game rpsGameState) Scissors() rpsOption {
	return rpsScissors
}

func main() {
	game := RockPaperScissors()

	game.Turn(rpsOne, rpsRock)
	game.Turn(rpsOne, rpsRock)

	fmt.Println(game)
}
