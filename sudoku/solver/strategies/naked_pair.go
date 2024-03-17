package strategies

import "sudoku/sudoku/solver/grid"

func UpdateNakedPairs(c *grid.Cell, g grid.CellGroup) (updated bool) {
	if c.Candidates().Length() != 2 {
		return false
	}
	otherCellsWithSameCandidates := g.Filter(func(cell *grid.Cell) bool {
		return cell.Candidates().Equals(c.Candidates())
	})

	// Iff there are exactly two cells with the same candidates
	if len(otherCellsWithSameCandidates) != 1 {
		return false
	}
	g.Subtract(otherCellsWithSameCandidates).ForEachCell(func(cell *grid.Cell) {
		newCandidates := cell.Candidates().Subtract(c.Candidates())
		if !newCandidates.Equals(cell.Candidates()) {
			updated = true
		}
		cell.SetCandidates(newCandidates)
	})
	return updated
}
