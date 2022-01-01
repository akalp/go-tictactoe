package main

import "fmt"

type GameState int

const (
	PLAYING GameState = iota
	X_WIN
	O_WIN
	TIE
)

func (g GameState) String() string {
	switch g {
	case TIE:
		return "Tie!"
	case X_WIN:
		return "X Win!"
	case O_WIN:
		return "O Win!"
	default:
		return "Playing"
	}
}

type Board struct {
	cells [][]Cell
}

func NewBoard() Board {
	cells := make([][]Cell, 3)
	for i := 0; i < 3; i++ {
		cells[i] = make([]Cell, 3)
	}
	return Board{
		cells: cells,
	}
}

func (b *Board) SetCell(x, y int, state CellState) (r bool) {
	if b.cells[y][x].IsEmpty() {
		b.cells[y][x].SetState(state)
		r = true
	}
	return
}

func (b Board) String() string {
	s := fmt.Sprintf("[%s] [%s] [%s]\n", b.cells[0][0], b.cells[0][1], b.cells[0][2])
	s += fmt.Sprintf("[%s] [%s] [%s]\n", b.cells[1][0], b.cells[1][1], b.cells[1][2])
	s += fmt.Sprintf("[%s] [%s] [%s]\n", b.cells[2][0], b.cells[2][1], b.cells[2][2])

	return s
}

func (b Board) State() GameState {
	if b.isWin(X) {
		return X_WIN
	}
	if b.isWin(O) {
		return O_WIN
	}
	if b.hasEmptyCell() {
		return PLAYING
	}
	return TIE
}

func (b Board) hasEmptyCell() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.cells[i][j].IsEmpty() {
				return true
			}
		}
	}
	return false
}

func (b Board) isWin(s CellState) bool {
	for i := 0; i < 3; i++ {
		if b.cells[i][0].state == s && b.cells[i][1].state == s && b.cells[i][2].state == s {
			return true
		}
		if b.cells[0][i].state == s && b.cells[1][i].state == s && b.cells[2][i].state == s {
			return true
		}
	}
	if b.cells[0][0].state == s && b.cells[1][1].state == s && b.cells[2][2].state == s {
		return true
	}
	if b.cells[2][0].state == s && b.cells[1][1].state == s && b.cells[0][2].state == s {
		return true
	}
	return false
}
