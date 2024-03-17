package sudoku

func updateNakedTriples(c *cell, g cellGroup) (updated bool) {
	if c.candidates.length() != 3 {
		return false
	}
	otherCellsWithSameCandidates := g.filter(func(cell *cell) bool {
		// Only consider cells with candidates
		if cell.candidates.length() < 1 {
			return false
		}
		return cell.candidates.equalOrSubsetOf(c.candidates)
	})

	// Iff there are exactly three cells with the same candidates or subset of candidates
	if len(otherCellsWithSameCandidates) != 2 {
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
