package main

import (
	"errors"
	"fmt"
	"strconv"

	game "github.com/limbuster/tic-tac-toe/internal/app"
	"github.com/manifoldco/promptui"
)

func main() {
	playerOne := "X"
	playerTwo := "O"
	player := playerOne
	state := game.State{Complete: false}
	g := game.NewGame(playerOne, playerTwo)

	for !state.Complete {
		position := getInput(player)
		err := g.Move(g.CurrentPlayer, position)
		for err != nil {
			fmt.Println("Invalid position, try again")
			position = getInput(player)
			err = g.Move(g.CurrentPlayer, position)
		}
		if player == playerOne {
			player = playerTwo
		} else if player == playerTwo {
			player = playerOne
		}
		fmt.Printf("player: %q\n", player)
		state = g.CheckGame()
		g.PrintGame()

		if state.Complete {
			fmt.Printf("Game is complete\n")
			if state.Winner == "" {
				fmt.Printf("The game was drawn")
			} else {
				fmt.Printf("Winner is %s\n", state.Winner)
			}
		}
	}
}

func getInput(player string) int {
	validate := func(input string) error {
		position, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("invalid number")
		}
		if position < 0 || position > 8 {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Player %q - select a position from 0-8", player),
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}

	position, _ := strconv.Atoi(result)

	return position
}
