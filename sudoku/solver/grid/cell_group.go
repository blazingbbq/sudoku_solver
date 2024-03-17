package grid

type CellGroup []*Cell

// Filter returns a new CellGroup with cells for which f returns true
func (cg CellGroup) Filter(f func(*Cell) bool) CellGroup {
	result := CellGroup{}
	for _, c := range cg {
		if f(c) {
			result = append(result, c)
		}
	}
	return result
}

// Subtract returns the cells in cg that are not in sub
func (cg CellGroup) Subtract(sub CellGroup) CellGroup {
	result := CellGroup{}
	for _, c := range cg {
		if !sub.Contains(c) {
			result = append(result, c)
		}
	}
	return result
}

// Contains returns true if c is in cg
func (cg CellGroup) Contains(c *Cell) bool {
	for _, cell := range cg {
		if cell == c {
			return true
		}
	}
	return false
}

// ForEachCell calls f for each cell in cg
func (cg CellGroup) ForEachCell(f func(*Cell)) {
	for _, c := range cg {
		f(c)
	}
}

// KnownValues returns the known values of cells in the group
func (cg CellGroup) KnownValues() CellGroupValues {
	values := CellGroupValues{}
	for _, c := range cg {
		if val, ok := c.GetValue(); ok {
			values = values.Append(val)
		}
	}
	return values
}

// Candidates returns the candidates of cells in the group
func (cg CellGroup) Candidates() CellGroupValues {
	values := CellGroupValues{}
	for _, c := range cg {
		if _, ok := c.GetValue(); !ok {
			values = values.Append(c.candidates...)
		}
	}
	return values
}

// KnownAndCandidateValues returns the known values and candidates of cells in the group
func (cg CellGroup) KnownAndCandidateValues() CellGroupValues {
	return cg.KnownValues().Append(cg.Candidates()...)
}

// RemoveCandidates removes candidates from cells in the group
func (cg CellGroup) RemoveCandidates(candidates CellGroupValues) {
	cg.ForEachCell(func(c *Cell) {
		c.candidates = c.candidates.Subtract(candidates)
	})
}
