package game

import (
	"errors"
	"fmt"
)

var ErrInvalidPosition = errors.New("invalid_position")
var ErrInvalidPlayer = errors.New("invalid_player")
var ErrAlreadyOccupied = errors.New("already_occupied")
var dimension = 3

func NewGame(playerOne, playerTwo string) Game {
	return Game{
		Positions:     initGame(),
		PlayerOne:     playerOne,
		PlayerTwo:     playerTwo,
		CurrentPlayer: playerOne,
	}
}

type State struct {
	Winner   string
	Complete bool
}

type Game struct {
	PlayerOne     string
	PlayerTwo     string
	Positions     [9]string
	CurrentPlayer string
}

func (g *Game) Move(player string, position int) error {
	if position > 8 || position < 0 {
		return ErrInvalidPosition
	}
	if player != g.PlayerOne && player != g.PlayerTwo {
		return ErrInvalidPlayer
	}
	if g.Positions[position] != "" {
		return ErrAlreadyOccupied
	}
	g.Positions[position] = player
	g.CurrentPlayer = g.nextTurn(player)
	return nil
}

func (g *Game) PrintGame() {
	for i := 0; i < len(g.Positions); i++ {
		p := g.Positions[i]
		if p != "" {
			fmt.Printf("%s", p)
		} else {
			fmt.Printf("_")
		}
		if (i+1)%3 == 0 {
			fmt.Printf("\n")
		}
	}
}

func (g *Game) CheckGame() State {
	// first check diagonally
	if checkDiagonal(g.PlayerOne, g.Positions) {
		return State{
			Winner:   g.PlayerOne,
			Complete: true,
		}
	}

	if checkDiagonal(g.PlayerTwo, g.Positions) {
		return State{
			Winner:   g.PlayerTwo,
			Complete: true,
		}
	}

	// Check by row and by column
	for i := 0; i < 3; i++ {
		if checkRow(g.PlayerOne, g.Positions, i) {
			return State{
				Winner:   g.PlayerOne,
				Complete: true,
			}
		}
		if checkRow(g.PlayerTwo, g.Positions, i) {
			return State{
				Winner:   g.PlayerTwo,
				Complete: true,
			}
		}
		if checkColumn(g.PlayerOne, g.Positions, i) {
			return State{
				Winner:   g.PlayerOne,
				Complete: true,
			}
		}
		if checkColumn(g.PlayerTwo, g.Positions, i) {
			return State{
				Winner:   g.PlayerTwo,
				Complete: true,
			}
		}
	}

	complete := true
	// although there was no winner check to ensure the game is complete or not
	for i := 0; i < len(g.Positions); i++ {
		if g.Positions[i] == "" {
			complete = false
		}
	}

	// if no one has won the game then return complete as false
	return State{
		Complete: complete,
	}
}

func (g *Game) nextTurn(player string) string {
	if player == g.PlayerOne {
		return g.PlayerTwo
	}
	return g.PlayerOne
}

func initGame() [9]string {
	return [9]string{"", "", "", "", "", "", "", "", ""}
}

func checkRow(player string, positions [9]string, row int) bool {
	if player == positions[row*dimension] &&
		player == positions[row*dimension+1] &&
		player == positions[row*dimension+2] {
		return true
	}
	return false
}

func checkColumn(player string, positions [9]string, col int) bool {
	if player == positions[col] &&
		player == positions[col+dimension] &&
		player == positions[col+dimension*2] {
		return true
	}
	return false
}

func checkDiagonal(player string, positions [9]string) bool {
	if player == positions[0] &&
		player == positions[4] &&
		player == positions[8] {
		return true
	}

	if player == positions[2] &&
		player == positions[4] &&
		player == positions[6] {
		return true
	}
	return false
}
