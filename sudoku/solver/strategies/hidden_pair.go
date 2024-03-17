package strategies

import "sudoku/sudoku/solver/grid"

func UpdateHiddenPairs(c *grid.Cell, g grid.CellGroup) (updated bool) {
	allCandidates := g.Candidates()

	possibleHiddenPairVals := c.Candidates().Filter(func(cand int) bool {
		return allCandidates.Count(cand) == 1
	})

	possibleHiddenPairVals.ForEachUniquePair(func(v1, v2 int) {
		g.Filter(func(cell *grid.Cell) bool {
			return cell.Candidates().Contains(v1) && cell.Candidates().Contains(v2)
		}).ForEachCell(func(cell *grid.Cell) {
			newCandidates := grid.CellGroupValues{v1, v2}
			if !newCandidates.Equals(cell.Candidates()) {
				updated = true
			}
			cell.SetCandidates(newCandidates)
			c.SetCandidates(newCandidates)
		})
	})
	return updated
}
