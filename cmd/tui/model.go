package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"sudoku/sudoku"
	"sudoku/sudoku/solver"
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
	solver *solver.Solver
}

var _ tea.Model = &model{}

func newModel() *model {
	m := &model{
		cursorX: 0,
		cursorY: 0,
		sudoku:  sudoku.NewSudoku([9][9]int{}),
	}
	m.solver = solver.NewSolver(m.sudoku)

	// TODO: Remove this hardcoded puzzle
	m.sudoku.ReadFromStrings([]string{
		"      24 ",
		"9  6 2   ",
		"    8   5",
		"2      6 ",
		"8  2  35 ",
		"  13     ",
		"   9614 3",
		"       8 ",
		"13   49  ",
	})
	m.puzzleVals = m.sudoku.GetBoard()

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
		case "n":
			if m.hintMode {
				_ = m.solver.SolveNextCell()
			}

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
			if m.sudoku.IsCellEmpty(i, j) {
				cell = " " // Empty cells are represented with 0. Render them as spaces

				if m.hintMode {
					// Display cells with a single possible value as a hint
					_, ok := m.solver.CellMustBe(i, j)
					if ok {
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
		cursorPosIndicator := fmt.Sprintf("Cell (%d, %d)", m.cursorX, m.cursorY)

		// Possible values for the selected cell
		possibleValues := m.sudoku.CandidatesForCell(m.cursorY, m.cursorX)
		possibleValuesIndicator := "[ "
		for i, v := range possibleValues {
			if i > 0 {
				possibleValuesIndicator += ", "
			}
			possibleValuesIndicator += fmt.Sprintf("%d", v)
		}
		possibleValuesIndicator += " ]"

		// Hint mode indicator
		hintModeIndicator := "Hint mode: "
		if m.hintMode {
			hintModeIndicator += "ON"
		} else {
			hintModeIndicator += "OFF"
		}
		hintModeIndicator = helpStyle(hintModeIndicator)

		mustBeIndicator := ""
		if m.hintMode {
			mustBe, ok := m.solver.CellMustBe(m.cursorY, m.cursorX)
			if ok && m.sudoku.IsCellEmpty(m.cursorY, m.cursorX) {
				mustBeIndicator = helpStyle(fmt.Sprintf("-> %d", mustBe))
			}
		}

		// Render panel with side border and padding
		sidePanel = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder(), false, false, false, true). // Left border
			BorderForeground(lipgloss.Color("#AAAAAA")).
			MarginLeft(4).
			Padding(1, 3).
			Render(lipgloss.JoinVertical(lipgloss.Left,
				cursorPosIndicator,
				possibleValuesIndicator,
				mustBeIndicator,
				"",
				"",
				hintModeIndicator,
			))
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

	regularHelpText := "  ←/↑/↓/→: Navigate • c: Puzzle entry • q: Quit"
	hintModeHelp := ""
	if m.showSidePanel {
		hintModeHelp = "  h: Toggle hint mode • n: Solve next cell"
	}
	helpText := helpStyle(lipgloss.JoinVertical(lipgloss.Left,
		regularHelpText,
		hintModeHelp,
	))

	// Send the UI for rendering
	return title + "\n" + centerPanel + "\n" + helpText + "\n"
}
