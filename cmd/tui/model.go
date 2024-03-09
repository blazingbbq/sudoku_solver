package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"sudoku/sudoku"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type model struct {
	cursorX int
	cursorY int
	sudoku  *sudoku.Sudoku
}

var _ tea.Model = &model{}

func newModel() *model {
	m := &model{
		cursorX: 0,
		cursorY: 0,
		sudoku:  sudoku.NewSudoku(),
	}

	return m
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) sanitizeCursorPos() {
	if m.cursorX < 0 {
		m.cursorX = 0
	}
	if m.cursorX > 8 {
		m.cursorX = 8
	}
	if m.cursorY < 0 {
		m.cursorY = 0
	}
	if m.cursorY > 8 {
		m.cursorY = 8
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "right", "d", "tab":
			m.cursorX++
		case "left", "a", "shift+tab":
			m.cursorX--
		case "up", "w":
			m.cursorY--
		case "down", "s", "enter":
			m.cursorY++
		
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.sudoku.SetCell(m.cursorY, m.cursorX, int(msg.String()[0] - '0'))
		case "backspace", "delete", "0":
			m.sudoku.SetCell(m.cursorY, m.cursorX, 0)
		}
	}
	m.sanitizeCursorPos()

	return m, cmd
}

func (m *model) View() string {
	titleAndBorderColor := "#D3D3D3"
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color(titleAndBorderColor)).
		Render("Sudoku")

	gameBoarder := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 3).
		BorderForeground(lipgloss.Color(titleAndBorderColor))
		// Width(64). // TODO: Set dynamically
		// Height(16) // TODO: Set dynamically

	s := ""

	// Render the game board
	board := m.sudoku.GetBoard()
	for i := range board {
		for j := range board[i] {
			cell := fmt.Sprintf("%d", board[i][j])
			if board[i][j] == 0 {
				cell = " " // Empty cells are represented with 0. Render them as spaces
			}

			cellBg := ""
			cellFg := "F"
			if i == m.cursorY && j == m.cursorX {
				cellBg = "#FFA500" // Highlight the cell the cursor is on
			}
			s += lipgloss.NewStyle().
				Background(lipgloss.Color(cellBg)).
				Foreground(lipgloss.Color(cellFg)).
				Render(cell)

			if j % 3 == 2 && j != 8 {
				s += " │ "
			} else {
				s += " "
			}
		}
		if i < 8 {
			s += "\n"
		}
		if i % 3 == 2 && i != 8 {
			s += "──────┼───────┼──────\n"
		}
	}

	helpText := helpStyle("  ←/↑/↓/→: Navigate • q: Quit")

	// Send the UI for rendering
	return title + "\n" + gameBoarder.Render(s) + "\n" + helpText + "\n"
}
