package game_test

import (
	"errors"
	"testing"

	game "github.com/limbuster/tic-tac-toe/internal/app"
)

func TestMove(t *testing.T) {
	playerOne := "player_one"
	playerTwo := "player_two"
	testcases := []struct {
		name      string
		player    string
		position  int
		err       error
		positions [9]string
	}{
		{
			name:     "Player two makes a valid move",
			player:   playerTwo,
			position: 3,
			positions: [9]string{
				playerOne, "", "",
				"", "", "",
				"", "", "",
			},
		},
		{
			name:     "Player one makes a valid move",
			player:   playerOne,
			position: 8,
			positions: [9]string{
				playerOne, playerTwo, playerTwo,
				playerTwo, playerTwo, playerOne,
				playerOne, playerOne, "",
			},
		},
		{
			name:     "Player one makes an invalid move - already_occupied",
			player:   playerOne,
			position: 2,
			positions: [9]string{
				playerOne, playerTwo, playerTwo,
				playerTwo, playerTwo, playerOne,
				playerOne, playerOne, "",
			},
			err: game.ErrAlreadyOccupied,
		},
		{
			name:     "Invalid move - invalid_player",
			player:   "random_guy",
			position: 2,
			positions: [9]string{
				playerOne, "", "",
				"", "", "",
				"", "", "",
			},
			err: game.ErrInvalidPlayer,
		},
		{
			name:     "Player one makes an invalid move - too high, invalid_position",
			player:   playerOne,
			position: 10,
			positions: [9]string{
				playerOne, "", "",
				"", "", "",
				"", "", "",
			},
			err: game.ErrInvalidPosition,
		},
		{
			name:     "Player one makes an invalid move - too low, invalid_position",
			player:   playerOne,
			position: -10,
			positions: [9]string{
				playerOne, "", "",
				"", "", "",
				"", "", "",
			},
			err: game.ErrInvalidPosition,
		},
	}
	for _, tt := range testcases {
		g := game.NewGame(playerOne, playerTwo)
		g.Positions = tt.positions
		e := g.Move(tt.player, tt.position)
		if !errors.Is(e, tt.err) {
			t.Errorf("[%s] Want: %v, Got: %v", tt.name, tt.err, e)
			t.FailNow()
		}
	}
}

func TestCheckComplete(t *testing.T) {
	playerOne := "player_one"
	playerTwo := "player_two"
	testcases := []struct {
		name      string
		winner    string
		complete  bool
		positions [9]string
	}{
		{
			name:   "Player one wins: Top row",
			winner: playerOne,
			positions: [9]string{
				playerOne, playerOne, playerOne,
				"", "", "",
				"", "", "",
			},
			complete: true,
		},
		{
			name:   "Player one wins: Middle row",
			winner: playerOne,
			positions: [9]string{
				"", "", "",
				playerOne, playerOne, playerOne,
				"", "", "",
			},
			complete: true,
		},
		{
			name:   "Player one wins: Bottom row",
			winner: playerOne,
			positions: [9]string{
				"", "", "",
				"", "", "",
				playerOne, playerOne, playerOne,
			},
			complete: true,
		},
		{
			name:   "Player one wins: Left column",
			winner: playerOne,
			positions: [9]string{
				playerOne, "", "",
				playerOne, "", "",
				playerOne, "", "",
			},
			complete: true,
		},
		{
			name:   "Player one wins: Middle column",
			winner: playerOne,
			positions: [9]string{
				"", playerOne, "",
				"", playerOne, "",
				"", playerOne, "",
			},
			complete: true,
		},
		{
			name:   "Player one wins: Right column",
			winner: playerOne,
			positions: [9]string{
				"", "", playerOne,
				"", "", playerOne,
				"", "", playerOne,
			},
			complete: true,
		},
		{
			name:   "Player one wins: Diagonal Left to Right",
			winner: playerOne,
			positions: [9]string{
				playerOne, "", "",
				"", playerOne, "",
				"", "", playerOne,
			},
			complete: true,
		},
		{
			name:   "Player one wins: Diagonal Right to Left",
			winner: playerOne,
			positions: [9]string{
				"", "", playerOne,
				"", playerOne, "",
				playerOne, "", "",
			},
			complete: true,
		},

		{
			name:   "Player two wins: Top row",
			winner: playerTwo,
			positions: [9]string{
				playerTwo, playerTwo, playerTwo,
				"", "", "",
				"", "", "",
			},
			complete: true,
		},
		{
			name:   "Player two wins: Middle row",
			winner: playerTwo,
			positions: [9]string{
				"", "", "",
				playerTwo, playerTwo, playerTwo,
				"", "", "",
			},
			complete: true,
		},
		{
			name:   "Player two wins: Bottom row",
			winner: playerTwo,
			positions: [9]string{
				"", "", "",
				"", "", "",
				playerTwo, playerTwo, playerTwo,
			},
			complete: true,
		},
		{
			name:   "Player two wins: Left column",
			winner: playerTwo,
			positions: [9]string{
				playerTwo, "", "",
				playerTwo, "", "",
				playerTwo, "", "",
			},
			complete: true,
		},
		{
			name:   "Player two wins: Middle column",
			winner: playerTwo,
			positions: [9]string{
				"", playerTwo, "",
				"", playerTwo, "",
				"", playerTwo, "",
			},
			complete: true,
		},
		{
			name:   "Player two wins: Right column",
			winner: playerTwo,
			positions: [9]string{
				"", "", playerTwo,
				"", "", playerTwo,
				"", "", playerTwo,
			},
			complete: true,
		},
		{
			name:   "Player two wins: Diagonal Left to Right",
			winner: playerTwo,
			positions: [9]string{
				playerTwo, "", "",
				"", playerTwo, "",
				"", "", playerTwo,
			},
			complete: true,
		},
		{
			name:   "Player two wins: Diagonal Right to Left",
			winner: playerTwo,
			positions: [9]string{
				"", "", playerTwo,
				"", playerTwo, "",
				playerTwo, "", "",
			},
			complete: true,
		},

		{
			name: "Game incomplete - I",
			positions: [9]string{
				"", playerTwo, "",
				"", playerOne, "",
				playerOne, "", "",
			},
			complete: false,
		},
		{
			name: "Game complete - No winner",
			positions: [9]string{
				playerOne, playerTwo, playerTwo,
				playerTwo, playerOne, playerOne,
				playerOne, playerOne, playerTwo,
			},
			complete: true,
		},
	}

	for _, tt := range testcases {
		g := game.NewGame(playerOne, playerTwo)
		want := game.State{
			Winner:   tt.winner,
			Complete: tt.complete,
		}
		g.Positions = tt.positions

		got := g.CheckGame()

		if want != got {
			t.Errorf("[%s] Want: %v, Got: %v", tt.name, want, got)
			t.FailNow()
		}
	}
}
