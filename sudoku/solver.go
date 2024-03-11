package sudoku

import (
	"errors"
)

type cell struct {
	rowIndex    int
	colIndex    int
	sqrRowIndex int
	sqrColIndex int

	value  *int
	row    cellGroup
	col    cellGroup
	square cellGroup
	region cellGroup

	candidates cellGroupValues
}

type cellGroup []*cell

type cellGroupValues []int

type grid [9][9]cell

type Solver struct {
	sudoku *Sudoku

	grid grid
}

func NewSolver(sudoku *Sudoku) *Solver {
	s := &Solver{sudoku: sudoku}
	s.newGrid()
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
	s.newGrid()

	// Solve the next cell that has only one candidate
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
	return errors.New("No cells with only one candidate found")
}

func (s *Solver) newGrid() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.grid[i][j] = s.newCell(i, j)
		}
	}
}

func (s *Solver) newCell(row, col int) cell {
	c := cell{
		value:       ptr(s.sudoku.GetCell(row, col)),
		rowIndex:    row,
		colIndex:    col,
		sqrRowIndex: row / 3,
		sqrColIndex: col / 3,
		candidates:  s.sudoku.CandidatesForCell(row, col),
	}

	for i := 0; i < 9; i++ {
		if i != col {
			c.row = append(c.row, &s.grid[row][i])
		}
		if i != row {
			c.col = append(c.col, &s.grid[i][col])
		}
		if i != (row%3)*3+col%3 {
			c.square = append(c.square, &s.grid[c.sqrRowIndex*3+i/3][c.sqrColIndex*3+i%3])
		}
	}

	c.region = append(c.region, c.row...)
	c.region = append(c.region, c.col...)
	c.region = append(c.region, c.square...)

	return c
}

func (c *cell) getValue() (int, bool) {
	if c.value == nil {
		return 0, false
	}
	return *c.value, *c.value != 0
}

func (c *cell) singleValue() (int, bool) {
	// If the cell has only one candidate, return it
	if len(c.candidates) == 1 {
		return c.candidates[0], true
	}

	// If a candidate is only present in one axis of the region, return it
	if vals := c.candidates.subtract(c.row.knownAndCandidateValues()); len(vals) == 1 {
		return vals[0], true
	}
	if vals := c.candidates.subtract(c.col.knownAndCandidateValues()); len(vals) == 1 {
		return vals[0], true
	}
	if vals := c.candidates.subtract(c.square.knownAndCandidateValues()); len(vals) == 1 {
		return vals[0], true
	}

	return 0, false
}

func (s *Solver) CellMustBe(row, col int) (int, bool) {
	s.newGrid()

	return s.grid[row][col].singleValue()
}

// append returns a new cellGroupValues with value appended
func (cgv cellGroupValues) append(value ...int) cellGroupValues {
	return append(cgv, value...)
}

// contains returns true if value is in cgv
func (cgv cellGroupValues) contains(value int) bool {
	for _, v := range cgv {
		if v == value {
			return true
		}
	}
	return false
}

// subtract returns the values in cgv that are not in sub
func (cgv cellGroupValues) subtract(sub cellGroupValues) cellGroupValues {
	result := cellGroupValues{}
	for _, v := range cgv {
		if !sub.contains(v) {
			result = result.append(v)
		}
	}
	return result
}

// knownValues returns the known values of cells in the group
func (cg cellGroup) knownValues() cellGroupValues {
	values := cellGroupValues{}
	for _, c := range cg {
		if val, ok := c.getValue(); ok {
			values = values.append(val)
		}
	}
	return values
}

// candidates returns the candidates of cells in the group
func (cg cellGroup) candidates() cellGroupValues {
	values := cellGroupValues{}
	for _, c := range cg {
		if _, ok := c.getValue(); !ok {
			values = values.append(c.candidates...)
		}
	}
	return values
}

// knownAndCandidateValues returns the known values and candidates of cells in the group
func (cg cellGroup) knownAndCandidateValues() cellGroupValues {
	return cg.knownValues().append(cg.candidates()...)
}

func ptr[K any](val K) *K {
	return &val
}
