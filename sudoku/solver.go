package sudoku

import (
	"errors"
)

type cell struct {
	row int
	col int
	sqr int

	possibleValues []int
}

type grid [9][9]cell

type Solver struct {
	sudoku *Sudoku

	grid [9][9]cell
}

func NewSolver(sudoku *Sudoku) *Solver {
	s := &Solver{sudoku: sudoku}
	s.resetPossibleValues()
	return s
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
	s.resetPossibleValues()

	// Solve the next cell that has only one possible value
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			if s.sudoku.board[i][j] != 0 {
				continue
			}

			mustBe, ok := s.CellMustBe(i, j)
			if !ok {
				continue
			}
			s.sudoku.SetCell(i, j, mustBe)
			return nil
		}
	}
	return errors.New("No cells with only one possible value found")
}

func (s *Solver) resetPossibleValues() {
	possibleValues := s.sudoku.PossibleValuesForAllCells()
	for i := 0; i < _gridSize; i++ {
		for j := 0; j < _gridSize; j++ {
			s.grid[i][j] = cell{
				row:            i,
				col:            j,
				sqr:            j/3 + (i/3)*3,
				possibleValues: possibleValues[i][j],
			}
		}
	}
}

func (s *Solver) cellHasOnePossibleValue(row, col int) bool {
	return len(s.grid[row][col].possibleValues) == 1
}

func (s *Solver) cellMustBeInRow(row, col int) (int, bool) {
	mustBe := []int{}
	otherCellsInRowCanBe := make(map[int]bool)
	for i := 0; i < _gridSize; i++ {
		if i == col {
			continue
		}
		if !s.sudoku.IsCellEmpty(row, i) {
			continue
		}
		for _, value := range s.grid[row][i].possibleValues {
			otherCellsInRowCanBe[value] = true
		}
	}
	for _, value := range s.grid[row][col].possibleValues {
		if !otherCellsInRowCanBe[value] {
			mustBe = append(mustBe, value)
		}
	}
	if len(mustBe) == 1 {
		return mustBe[0], true
	}
	return 0, false
}

func (s *Solver) cellMustBeInColumn(row, col int) (int, bool) {
	mustBe := []int{}
	otherCellsInColumnCanBe := make(map[int]bool)
	for i := 0; i < _gridSize; i++ {
		if i == row {
			continue
		}
		if !s.sudoku.IsCellEmpty(i, col) {
			continue
		}
		for _, value := range s.grid[i][col].possibleValues {
			otherCellsInColumnCanBe[value] = true
		}
	}
	for _, value := range s.grid[row][col].possibleValues {
		if !otherCellsInColumnCanBe[value] {
			mustBe = append(mustBe, value)
		}
	}
	if len(mustBe) == 1 {
		return mustBe[0], true
	}
	return 0, false
}

func (s *Solver) cellMustBeInSquare(row, col int) (int, bool) {
	mustBe := []int{}
	otherCellsInSquareCanBe := make(map[int]bool)
	startRow, startCol := row-row%3, col-col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if i == row && j == col {
				continue
			}
			if !s.sudoku.IsCellEmpty(i, j) {
				continue
			}
			for _, value := range s.grid[i][j].possibleValues {
				otherCellsInSquareCanBe[value] = true
			}
		}
	}
	for _, value := range s.grid[row][col].possibleValues {
		if !otherCellsInSquareCanBe[value] {
			mustBe = append(mustBe, value)
		}
	}
	if len(mustBe) == 1 {
		return mustBe[0], true
	}
	return 0, false
}

func (s *Solver) CellMustBe(row, col int) (int, bool) {
	s.resetPossibleValues()

	if s.cellHasOnePossibleValue(row, col) {
		return s.grid[row][col].possibleValues[0], true
	}

	if mustBe, ok := s.cellMustBeInRow(row, col); ok {
		return mustBe, true
	}
	if mustBe, ok := s.cellMustBeInColumn(row, col); ok {
		return mustBe, true
	}
	if mustBe, ok := s.cellMustBeInSquare(row, col); ok {
		return mustBe, true
	}

	return 0, false
}
