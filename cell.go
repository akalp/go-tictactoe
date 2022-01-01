package main

type CellState int

const (
	EMPTY CellState = iota
	X
	O
)

func (c CellState) String() string {
	switch c {
	case EMPTY:
		return " "
	case X:
		return "x"
	case O:
		return "o"
	}
	return "UNKNOWN"
}

type Cell struct {
	state CellState
}

func NewCell() Cell {
	return Cell{EMPTY}
}

func (c Cell) IsEmpty() bool {
	return c.state == EMPTY
}

func (c *Cell) SetState(ns CellState) {
	c.state = ns
}

func (c Cell) String() string {
	return c.state.String()
}
