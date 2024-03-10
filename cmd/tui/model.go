package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"sudoku/sudoku"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type model struct {
	cursorX       int
	cursorY       int
	showSidePanel bool

	hintMode           bool
	puzzleCreationMode bool
	puzzleVals         [9][9]int

	sudoku *sudoku.Sudoku
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

		case "right", "d", "enter":
			m.cursorX++
		case "left", "a":
			m.cursorX--
		case "up", "w":
			m.cursorY--
		case "down", "s":
			m.cursorY++

		case "tab":
			m.showSidePanel = !m.showSidePanel
		case "h", " ":
			m.hintMode = !m.hintMode

		case "c", "i":
			m.puzzleCreationMode = !m.puzzleCreationMode

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			val := int(msg.String()[0] - '0')

			// Only allow modification of non-puzzle provided cells, unless in puzzle creation mode
			shouldSetCell := false
			if m.puzzleCreationMode {
				shouldSetCell = true
			} else if m.puzzleVals[m.cursorY][m.cursorX] == 0 {
				shouldSetCell = true
			}
			if shouldSetCell {
				m.sudoku.SetCell(m.cursorY, m.cursorX, val)
			}

			if m.puzzleCreationMode {
				m.puzzleVals[m.cursorY][m.cursorX] = val
			}
		case "backspace", "delete", "0":
			shouldSetCell := false
			if m.puzzleCreationMode {
				shouldSetCell = true
			} else if m.puzzleVals[m.cursorY][m.cursorX] == 0 {
				shouldSetCell = true
			}
			if shouldSetCell {
				m.sudoku.SetCell(m.cursorY, m.cursorX, 0)
			}

			if m.puzzleCreationMode {
				m.puzzleVals[m.cursorY][m.cursorX] = 0
			}
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

	// Render the gameboard
	gameboard := ""
	board := m.sudoku.GetBoard()
	for i := range board {
		for j := range board[i] {
			cellStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("F"))

			cell := fmt.Sprintf("%d", board[i][j])
			if board[i][j] == 0 {
				cell = " " // Empty cells are represented with 0. Render them as spaces

				if m.hintMode {
					// Display cells with a single possible value as a hint
					possibleValues := m.sudoku.PossibleValuesForCell(i, j)
					if len(possibleValues) == 1 {
						cell = "*"
						cellStyle = cellStyle.Foreground(lipgloss.Color("#FFFF00"))
					}
				}
			}

			isCellSelected := i == m.cursorY && j == m.cursorX
			isCellValid := m.sudoku.IsCellValid(i, j)
			isPuzzleProvided := m.puzzleVals[i][j] != 0

			if isCellSelected {
				if isCellValid || isPuzzleProvided {
					// Highlight the cell the cursor is on
					cellStyle.Background(lipgloss.Color("#FFA500"))
				} else {
					// Highlight in light red if the cell is invalid, selected and not a puzzle value
					cellStyle.Background(lipgloss.Color("#FF7547"))
				}
			} else if isPuzzleProvided {
				// Highlight in gray if the cell is a puzzle value
				cellStyle.Background(lipgloss.Color("#565656"))
			} else if !isCellValid {
				// Highlight in red if the cell is invalid
				cellStyle.Background(lipgloss.Color("#FF3030"))
			}
			gameboard += cellStyle.Render(cell)

			if j%3 == 2 && j != 8 {
				gameboard += " │ "
			} else if j != 8 {
				gameboard += " "
			}
		}
		if i < 8 {
			gameboard += "\n"
		}
		if i%3 == 2 && i != 8 {
			gameboard += "──────┼───────┼──────\n"
		}
	}

	// Render the side panel
	sidePanel := ""
	if m.showSidePanel {
		sidePanel += fmt.Sprintf("Cell (%d, %d)\n", m.cursorX, m.cursorY)

		// Possible values for the selected cell
		possibleValues := m.sudoku.PossibleValuesForCell(m.cursorY, m.cursorX)
		vals := "[ "
		for i, v := range possibleValues {
			if i > 0 {
				vals += ", "
			}
			vals += fmt.Sprintf("%d", v)
		}
		vals += " ]"
		sidePanel += vals

		// Render panel with side border and padding
		sidePanel = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder(), false, false, false, true). // Left border
			BorderForeground(lipgloss.Color("#AAAAAA")).
			MarginLeft(4).
			Padding(1, 3).
			Render("\n" + sidePanel + "\n")
	}

	puzzleCreationPrompt := ""
	if m.puzzleCreationMode {
		puzzleCreationPrompt += "\n"
		puzzleCreationPrompt += helpStyle("🔴 Enter puzzle values")
	}

	centerPanel := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 3).
		BorderForeground(lipgloss.Color(titleAndBorderColor)).
		// Width(64).
		// Height(16).
		Render(
			lipgloss.JoinVertical(lipgloss.Left,
				lipgloss.JoinHorizontal(lipgloss.Center, gameboard, sidePanel),
			),
			puzzleCreationPrompt,
		)
	helpText := helpStyle("  ←/↑/↓/→: Navigate • c: Puzzle entry • q: Quit")

	// Send the UI for rendering
	return title + "\n" + centerPanel + "\n" + helpText + "\n"
}
