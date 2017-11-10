package main

import (
	"errors"
	"fmt"
)

type rpsPlayer int
type rpsOption int

type rpsGameState struct {
	winner rpsPlayer
	one    rpsOption
	two    rpsOption
}

const (
	rpsNone rpsPlayer = -1
	rpsOne  rpsPlayer = 0
	rpsTwo  rpsPlayer = 1

	rpsEmpty    rpsOption = -1
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
	game.one = rpsEmpty
	game.two = rpsEmpty
	return game
}

func (game *rpsGameState) Turn(player rpsPlayer, option rpsOption) (bool, error) {
	if player == rpsNone {
		return false, errors.New("Invalid player")
	}

	if option == rpsEmpty {
		return false, errors.New("Invalid option")
	}

	switch player {
	case rpsOne:
		game.one = option

	case rpsTwo:
		game.two = option
	}

	if game.one != rpsEmpty && game.two != rpsEmpty {
		game.winner = calculateRpcWinner(game.one, game.two)
		return true, nil
	}

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

func calculateRpcWinner(optOne, optTwo rpsOption) rpsPlayer {
	if optOne == optTwo {
		return rpsNone
	} else if optOne == rpsRock && optTwo == rpsScissors {
		return rpsOne
	} else if optOne == rpsScissors && optTwo == rpsPaper {
		return rpsOne
	} else if optOne == rpsPaper && optTwo == rpsRock {
		return rpsOne
	} else {
		return rpsTwo
	}
}

func main() {
	game := RockPaperScissors()

	game.Turn(rpsOne, rpsRock)
	game.Turn(rpsTwo, rpsPaper)

	fmt.Println(game)
}
