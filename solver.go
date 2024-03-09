package main

import (
	"errors"
	"fmt"
)

type Solver struct {
	sudoku *Sudoku

	// List of possible values for each cell, where an empty means the cell is
	// already solved
	possibleValues [_gridSize][_gridSize][]int
}

func NewSolver(sudoku *Sudoku) *Solver {
	s := &Solver{sudoku: sudoku}
	s.initPossibleValues()
	return s
}

func (s *Solver) initPossibleValues() {
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			s.possibleValues[i][j] = s.sudoku.possibleValuesForCell(i, j)
		}
	}
}

func (s *Solver) Solve() (*Sudoku, error) {
	for {
		// Check if the sudoku is invalid
		if !s.sudoku.isValid() {
			return nil, errors.New("invalid sudoku")
		}

		// Check if the sudoku is solved
		if s.sudoku.isSolved() {
			return s.sudoku, nil
		}

		// Attempt to solve the next cell
		if err := s.solveNextCell(); err != nil {
			return nil, err
		}
	}
}

func (s *Solver) solveNextCell() error {
	s.initPossibleValues() // TODO: optimize this (only update affected cells)

	// Find the cell with the fewest possible values
	minI, minJ := -1, -1
	minLen := _gridSize + 1
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			if s.sudoku.board[i][j] != 0 {
				continue
			}

			numPossibleValues := len(s.possibleValues[i][j])
			if numPossibleValues != 0 && numPossibleValues < minLen {
				minI, minJ = i, j
				minLen = len(s.possibleValues[i][j])
			}
		}
	}

	// Check that there are empty cells left
	if minI == -1 {
		return errors.New("no empty cell with possible values found")
	}

	// Ensure that the cell with the least possible values only has one possible
	// value
	if minLen > 1 {
		return errors.New("no unique solution found")
	}

	// Solve the cell
	s.sudoku.board[minI][minJ] = s.possibleValues[minI][minJ][0]
	return nil
}

func (s *Solver) FormatCellsWithSinglePossibleValue() string {
	var result string
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			if s.sudoku.board[i][j] != 0 {
				result += fmt.Sprintf("%d", s.sudoku.board[i][j])
			} else {
				if len(s.possibleValues[i][j]) == 1 {
					result += "*"
				} else {
					result += " "
				}
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

