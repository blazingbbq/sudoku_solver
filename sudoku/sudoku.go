package sudoku

import (
	"errors"
	"fmt"
)

const _gridSize = 9

type Sudoku struct {
	board [_gridSize][_gridSize]int
}

func NewSudoku() *Sudoku {
	return &Sudoku{}
}

func (s *Sudoku) GetBoard() [9][9]int {
	return s.board
}

func (s *Sudoku) SetCell(row, col, value int) {
	s.board[row][col] = value
}

func (s *Sudoku) ReadFromStrings(lines []string) (*Sudoku, error) {
	if len(lines) < _gridSize {
		return nil, errors.New("invalid number of rows")
	}

	for i, line := range lines {
		if len(line) < _gridSize {
			return nil, fmt.Errorf("invalid number of columns (row %d)", i)
		}
		if i >= _gridSize {
			break
		}

		for j, char := range line {
			if j >= _gridSize {
				break
			}

			if char == ' ' {
				s.board[i][j] = 0
			} else {
				s.board[i][j] = int(char - '0')
			}

			if s.board[i][j] < 0 || s.board[i][j] > 9 {
				return nil, errors.New("invalid character")
			}
		}
	}
	return s, nil
}

func (s *Sudoku) String() string {
	var result string
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			if s.board[i][j] == 0 {
				result += " "
			} else {
				result += fmt.Sprintf("%d", s.board[i][j])
			}

			if j % 3 == 2 && j != _gridSize - 1 {
				result += " | "
			} else {
				result += " "
			}
		}
		result += "\n"
		if i % 3 == 2 && i != _gridSize - 1 {
			result += "---------------------\n"
		}
	}
	return result
}

func (s *Sudoku) Solve() (*Sudoku, error) {
	if !s.isValid() {
		return nil, errors.New("invalid sudoku")
	}

	solver := NewSolver(s)
	return solver.Solve()
}

func (s *Sudoku) isValid() bool {
	// Check rows
	for i := 0; i < _gridSize; i++ {
		seen := make(map[int]bool)
		for j := 0; j < _gridSize; j++ {
			if s.board[i][j] != 0 {
				if seen[s.board[i][j]] {
					return false
				}
				seen[s.board[i][j]] = true
			}
		}
	}

	// Check columns
	for j := 0; j < _gridSize; j++ {
		seen := make(map[int]bool)
		for i := 0; i < _gridSize; i++ {
			if s.board[i][j] != 0 {
				if seen[s.board[i][j]] {
					return false
				}
				seen[s.board[i][j]] = true
			}
		}
	}

	// Check 3x3 squares
	for i := 0; i < _gridSize; i += 3 {
		for j := 0; j < _gridSize; j += 3 {
			seen := make(map[int]bool)
			for k := i; k < i + 3; k++ {
				for l := j; l < j + 3; l++ {
					if s.board[k][l] != 0 {
						if seen[s.board[k][l]] {
							return false
						}
						seen[s.board[k][l]] = true
					}
				}
			}
		}
	}

	return true
}

func (s *Sudoku) isSolved() bool {
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			if s.board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) possibleValuesForCell(row, col int) []int {
	if s.board[row][col] != 0 {
		return nil
	}

	seen := make(map[int]bool)
	for i := 0; i < _gridSize; i++ {
		seen[s.board[row][i]] = true
		seen[s.board[i][col]] = true
	}

	startRow, startCol := row - row % 3, col - col % 3
	for i := startRow; i < startRow + 3; i++ {
		for j := startCol; j < startCol + 3; j++ {
			seen[s.board[i][j]] = true
		}
	}

	var result []int
	for i := 1; i <= _gridSize; i++ {
		if !seen[i] {
			result = append(result, i)
		}
	}
	return result
}
