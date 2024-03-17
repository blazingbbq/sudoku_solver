package sudoku

type cellGroup []*cell

// filter returns a new cellGroup with cells for which f returns true
func (cg cellGroup) filter(f func(*cell) bool) cellGroup {
	result := cellGroup{}
	for _, c := range cg {
		if f(c) {
			result = append(result, c)
		}
	}
	return result
}

// subtract returns the cells in cg that are not in sub
func (cg cellGroup) subtract(sub cellGroup) cellGroup {
	result := cellGroup{}
	for _, c := range cg {
		if !sub.contains(c) {
			result = append(result, c)
		}
	}
	return result
}

// contains returns true if c is in cg
func (cg cellGroup) contains(c *cell) bool {
	for _, cell := range cg {
		if cell == c {
			return true
		}
	}
	return false
}

// forEachCell calls f for each cell in cg
func (cg cellGroup) forEachCell(f func(*cell)) {
	for _, c := range cg {
		f(c)
	}
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

// removeCandidates removes candidates from cells in the group
func (cg cellGroup) removeCandidates(candidates cellGroupValues) {
	cg.forEachCell(func(c *cell) {
		c.candidates = c.candidates.subtract(candidates)
	})
}
