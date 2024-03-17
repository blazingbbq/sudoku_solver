package strategies

import "sudoku/sudoku/solver/grid"

func UpdateNakedTriples(c *grid.Cell, g grid.CellGroup) (updated bool) {
	if c.Candidates().Length() != 3 {
		return false
	}
	otherCellsWithSameCandidates := g.Filter(func(cell *grid.Cell) bool {
		// Only consider cells with candidates
		if cell.Candidates().Length() < 1 {
			return false
		}
		return cell.Candidates().EqualOrSubsetOf(c.Candidates())
	})

	// Iff there are exactly three cells with the same candidates or subset of candidates
	if len(otherCellsWithSameCandidates) != 2 {
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
