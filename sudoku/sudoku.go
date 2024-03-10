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

func (s *Sudoku) GetCell(row, col int) int {
	return s.board[row][col]
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

			if j%3 == 2 && j != _gridSize-1 {
				result += " | "
			} else {
				result += " "
			}
		}
		result += "\n"
		if i%3 == 2 && i != _gridSize-1 {
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

func (s *Sudoku) IsCellValid(x, y int) bool {
	if s.board[x][y] == 0 {
		return true
	}

	// Check row
	for i := 0; i < _gridSize; i++ {
		if i != x && s.board[i][y] == s.board[x][y] {
			return false
		}
	}

	// Check column
	for i := 0; i < _gridSize; i++ {
		if i != y && s.board[x][i] == s.board[x][y] {
			return false
		}
	}

	// Check 3x3 square
	startRow, startCol := x-x%3, y-y%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if i != x && j != y && s.board[i][j] == s.board[x][y] {
				return false
			}
		}
	}

	return true
}

func (s *Sudoku) isValid() bool {
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			if s.board[i][j] != 0 && !s.IsCellValid(i, j) {
				return false
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

func (s *Sudoku) PossibleValuesForCell(row, col int) []int {
	if s.board[row][col] != 0 {
		return nil
	}

	seen := make(map[int]bool)
	for i := 0; i < _gridSize; i++ {
		seen[s.board[row][i]] = true
		seen[s.board[i][col]] = true
	}

	startRow, startCol := row-row%3, col-col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
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
