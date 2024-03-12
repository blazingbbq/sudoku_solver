package sudoku

import (
	"errors"
)

type cell struct {
	rowIndex    int
	colIndex    int
	sqrRowIndex int
	sqrColIndex int

	value  int
	row    cellGroup
	col    cellGroup
	square cellGroup
	region cellGroup

	candidates cellGroupValues
}

type grid [9][9]*cell

type Solver struct {
	sudoku *Sudoku

	grid *grid
}

func NewSolver(sudoku *Sudoku) *Solver {
	s := &Solver{sudoku: sudoku}
	s.newGrid()
	s.updateCandidates()
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

func (s *Solver) CellMustBe(row, col int) (int, bool) {
	s.updateCandidates()
	return s.grid[row][col].singleValue()
}

func (s *Solver) solveNextCell() error {
	s.updateCandidates()

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
	s.grid = &grid{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.grid[i][j] = s.newCell(i, j)
		}
	}

	s.grid.forEachCell(func(c *cell) {
		s.associateCell(c.rowIndex, c.colIndex)
	})
}

func (s *Solver) newCell(row, col int) *cell {
	return &cell{
		rowIndex:    row,
		colIndex:    col,
		sqrRowIndex: row / 3,
		sqrColIndex: col / 3,
	}
}

func (s *Solver) associateCell(row, col int) {
	c := s.grid[row][col]

	for i := 0; i < 9; i++ {
		if i != col {
			c.row = append(c.row, s.grid[row][i])
		}
		if i != row {
			c.col = append(c.col, s.grid[i][col])
		}
		if i != (row%3)*3+col%3 {
			c.square = append(c.square, s.grid[c.sqrRowIndex*3+i/3][c.sqrColIndex*3+i%3])
		}
	}

	c.region = append(c.region, c.row...)
	c.region = append(c.region, c.col...)
	c.region = append(c.region, c.square...)
}

func (s *Solver) sanitizedCandidatesForCell(row, col int) cellGroupValues {
	if _, ok := s.grid[row][col].getValue(); ok {
		return cellGroupValues{}
	}
	return s.sudoku.CandidatesForCell(row, col)
}

func (s *Solver) updateCandidates() {
	// Reset all candidates and values
	s.grid.forEachCell(func(c *cell) {
		c.value = s.sudoku.GetCell(c.rowIndex, c.colIndex)
		c.candidates = s.sanitizedCandidatesForCell(c.rowIndex, c.colIndex)
	})

	s.grid.forEachCell(func(c *cell) {
		c.forEachRegion(func(r cellGroup) {
			updateNakedPairs(c, r)
		})
	})
}

func (g *grid) forEachCell(f func(c *cell)) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			f(g[i][j])
		}
	}
}

func (c *cell) forEachRegion(f func(cellGroup)) {
	f(c.row)
	f(c.col)
	f(c.square)
}

func (c *cell) getValue() (int, bool) {
	return c.value, c.value != 0
}

func (c *cell) singleValue() (int, bool) {
	// If the cell has only one candidate, return it
	if len(c.candidates) == 1 {
		return c.candidates[0], true
	}

	// If a candidate is only present in one region, return it
	val, found := 0, false
	c.forEachRegion(func(region cellGroup) {
		if vals := c.candidates.subtract(region.knownAndCandidateValues()); len(vals) == 1 {
			val = vals[0]
			found = true
		}
	})
	return val, found
}
