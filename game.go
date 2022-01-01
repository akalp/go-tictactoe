package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model{
		board: NewBoard(),
		turn:  X,
	}, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	board Board
	turn  CellState
	state GameState
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state = m.board.State()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
			return m, tea.Quit
		} else if m.state != PLAYING && s == "enter" {
			m.board = NewBoard()
			m.state = PLAYING
		}

	case tea.MouseMsg:
		x := (msg.X - 1) / 4
		y := (msg.Y - 3)
		if msg.Type == tea.MouseLeft && x >= 0 && x < 3 && y >= 0 && y < 3 && m.state == PLAYING {
			if m.board.SetCell(x, y, m.turn) {
				m.ToggleTurn()
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Famous Tic-Tac-Toe game in Go using bubbletea. (q for quit.)\n"

	if m.state == PLAYING {
		s += fmt.Sprintf("Turn: %s\n\n", m.turn)
		s += m.board.String()
	} else {
		s += fmt.Sprintf("Result: %s. (enter for restart, q for quit)", m.state)
	}

	return s
}

func (m *model) ToggleTurn() {
	if m.turn == X {
		m.turn = O
	} else {
		m.turn = X
	}
}
