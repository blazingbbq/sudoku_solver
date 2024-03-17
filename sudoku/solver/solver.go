package solver

import (
	"errors"

	"sudoku/sudoku"
	"sudoku/sudoku/solver/grid"
	"sudoku/sudoku/solver/strategies"
)

type Solver struct {
	sudoku *sudoku.Sudoku

	grid *grid.Grid
}

func NewSolver(sudoku *sudoku.Sudoku) *Solver {
	s := &Solver{sudoku: sudoku}
	s.newGrid()
	s.updateCandidates()
	return s
}

func (s *Solver) Solve() (*sudoku.Sudoku, error) {
	for {
		// Check if the sudoku is invalid
		if !s.sudoku.IsValid() {
			return nil, errors.New("invalid sudoku")
		}

		// Check if the sudoku is solved
		if s.sudoku.IsSolved() {
			return s.sudoku, nil
		}

		// Attempt to solve the next cell
		if err := s.SolveNextCell(); err != nil {
			return nil, err
		}
	}
}

func (s *Solver) CellMustBe(row, col int) (int, bool) {
	s.updateCandidates()
	return s.grid[row][col].SingleValue()
}

func (s *Solver) SolveNextCell() error {
	s.updateCandidates()

	// Solve the next cell that has only one candidate
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.sudoku.GetCell(i, j) != 0 {
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
	s.grid = &grid.Grid{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.grid[i][j] = s.newCell(i, j)
		}
	}
	s.grid.AssociateCells()
}

func (s *Solver) newCell(row, col int) *grid.Cell {
	return grid.NewCell(row, col)
}

func (s *Solver) sanitizedCandidatesForCell(row, col int) grid.CellGroupValues {
	if _, ok := s.grid[row][col].GetValue(); ok {
		return grid.CellGroupValues{}
	}
	return s.sudoku.CandidatesForCell(row, col)
}

func (s *Solver) updateCandidates() {
	// Reset all candidates and values
	s.grid.ForEachCell(func(c *grid.Cell) {
		c.SetValue(s.sudoku.GetCell(c.RowIndex(), c.ColIndex()))
		c.SetCandidates(s.sanitizedCandidatesForCell(c.RowIndex(), c.ColIndex()))
	})

	s.grid.ForEachCell(func(c *grid.Cell) {
		c.ForEachRegion(func(r grid.CellGroup) {
			strategies.UpdateNakedPairs(c, r)
			strategies.UpdateHiddenPairs(c, r)
			strategies.UpdateNakedTriples(c, r)
		})
	})
}
