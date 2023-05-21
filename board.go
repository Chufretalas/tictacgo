package main

import (
	"errors"
	"fmt"
	"strings"
)

func NewBoard() Board {
	return Board{[3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}}
}

type Board struct {
	board [3][3]string
}

func (v *Board) Show() {
	fmt.Println("|=========|")
	for i := 0; i < len(v.board); i++ {
		fmt.Printf("   %v\n", strings.Join(v.board[i][:], " "))
	}
	fmt.Println("|=========|")
}

func (v *Board) X(line int, column int) error {
	return v.Put("X", line, column)
}

func (v *Board) O(line int, column int) error {
	return v.Put("O", line, column)
}

func (v *Board) Put(char string, line int, column int) error {
	if line > 2 || column > 2 {
		return errors.New("out of bounds")
	}
	if c := v.board[line][column]; c == "-" {
		v.board[line][column] = char
		return nil
	}
	return fmt.Errorf("space already used by: %v", v.board[line][column])
}

func (v *Board) Check() string {
	for i := 0; i < 3; i++ {
		if v.board[i][0] == v.board[i][1] && v.board[i][1] == v.board[i][2] && v.board[i][0] != "-" {
			fmt.Printf("%v Line: %o", v.board[i][0], i)
			return v.board[1][0]
		}
		if v.board[0][i] == v.board[1][i] && v.board[1][i] == v.board[2][i] && v.board[0][1] != "-" {
			fmt.Printf("%v Column: %o\n", v.board[0][i], i)
			return v.board[1][0]
		}
	}
	if v.board[0][0] == v.board[1][1] && v.board[1][1] == v.board[2][2] && v.board[1][1] != "-" {
		fmt.Printf("%v Diagonal\n", v.board[1][1])
		return v.board[1][0]
	}
	if v.board[0][2] == v.board[1][1] && v.board[1][1] == v.board[2][0] && v.board[1][1] != "-" {
		fmt.Printf("%v Inv. Diagonal\n", v.board[1][1])
		return v.board[1][0]
	}
	fmt.Println("none")
	return ""
}
