package sudoku

func updateHiddenPairs(c *cell, g cellGroup) (updated bool) {
	allCandidates := g.candidates()

	possibleHiddenPairVals := c.candidates.filter(func(cand int) bool {
		return allCandidates.count(cand) == 1
	})

	possibleHiddenPairVals.forEachUniquePair(func(v1, v2 int) {
		g.filter(func(cell *cell) bool {
			return cell.candidates.contains(v1) && cell.candidates.contains(v2)
		}).forEachCell(func(cell *cell) {
			newCandidates := cellGroupValues{v1, v2}
			if !newCandidates.equals(cell.candidates) {
				updated = true
			}
			cell.candidates = newCandidates
			c.candidates = newCandidates
		})
	})
	return updated
}
