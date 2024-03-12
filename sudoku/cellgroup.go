package sudoku


type cellGroup []*cell

type cellGroupValues []int


// append returns a new cellGroupValues with value appended
func (cgv cellGroupValues) append(value ...int) cellGroupValues {
	return append(cgv, value...)
}

// contains returns true if value is in cgv
func (cgv cellGroupValues) contains(value int) bool {
	for _, v := range cgv {
		if v == value {
			return true
		}
	}
	return false
}

// subtract returns the values in cgv that are not in sub
func (cgv cellGroupValues) subtract(sub cellGroupValues) cellGroupValues {
	result := cellGroupValues{}
	for _, v := range cgv {
		if !sub.contains(v) {
			result = result.append(v)
		}
	}
	return result
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

func ptr[K any](val K) *K {
	return &val
}
