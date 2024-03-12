package sudoku

func updateNakedPairs(c *cell, g cellGroup) (updated bool) {
	if c.candidates.length() != 2 {
		return false
	}
	otherCellsWithSameCandidates := g.filter(func(cell *cell) bool {
		return cell.candidates.equals(c.candidates)
	})

	// Iff there are exactly two cells with the same candidates
	if len(otherCellsWithSameCandidates) != 1 {
		return false
	}
	g.subtract(otherCellsWithSameCandidates).forEachCell(func(cell *cell) {
		newCandidates := cell.candidates.subtract(c.candidates)
		if !newCandidates.equals(cell.candidates) {
			updated = true
		}
		cell.candidates = newCandidates
	})
	return updated
}
