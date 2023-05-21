package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// role here means if he is going to play as an X or an O
func NewPlayer(board *Board, role string) (Player, error) {
	if role != "X" && role != "O" {
		fmt.Println("veio aqui")
		return Player{}, errors.New("role can only be a 'X' or an 'O' string")
	}
	return Player{board: board, role: role}, nil
}

type Player struct {
	board *Board
	role  string
}

func (v *Player) PlayRandom() error {
	av := v.board.GetAvailable()
	if len(av) == 0 {
		return errors.New("board is full")
	}

	choiceId := rand.Intn(len(av))

	switch choice := av[choiceId]; v.role {
	case "X":
		v.board.X(choice[0], choice[1])
	case "O":
		v.board.O(choice[0], choice[1])
	default:
		return errors.New("please use the NewPlayer function to create a new Player")
	}
	return nil
}
