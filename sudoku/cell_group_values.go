package sudoku


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
	if cgv == nil {
		return nil
	}

	result := cellGroupValues{}
	for _, v := range cgv {
		if !sub.contains(v) {
			result = result.append(v)
		}
	}
	return result
}

// equals returns true if cgv and other contain the same values
func (cgv cellGroupValues) equals(other cellGroupValues) bool {
	if len(cgv) != len(other) {
		return false
	}
	for _, v := range cgv {
		if !other.contains(v) {
			return false
		}
	}
	return true
}

// equalOrSubsetOf returns true if cgv is a subset of other
func (cgv cellGroupValues) equalOrSubsetOf(other cellGroupValues) bool {
	for _, v := range cgv {
		if !other.contains(v) {
			return false
		}
	}
	return true
}

// count returns the number of times value appears in cgv
func (cgv cellGroupValues) count(value int) int {
	count := 0
	for _, v := range cgv {
		if v == value {
			count++
		}
	}
	return count
}

// length returns the number of values in cgv
func (cgv cellGroupValues) length() int {
	return len(cgv)
}

// transform returns a new cellGroupValues with the result of f for each value in cgv
func (cgv cellGroupValues) transform(f func(int) (int, bool)) cellGroupValues {
	result := cellGroupValues{}
	for _, v := range cgv {
		if newVal, ok := f(v); ok {
			result = result.append(newVal)
		}
	}
	return result
}

// filter returns a new cellGroupValues with values for which f returns true
func (cgv cellGroupValues) filter(f func(int) bool) cellGroupValues {
	result := cellGroupValues{}
	for _, v := range cgv {
		if f(v) {
			result = result.append(v)
		}
	}
	return result
}

// forEach calls f for each value in cgv
func (cgv cellGroupValues) forEach(f func(int)) {
	for _, v := range cgv {
		f(v)
	}
}

// forEachUniquePair calls f for each unique pair of values in cgv
func (cgv cellGroupValues) forEachUniquePair(f func(int, int)) {
	for i, v1 := range cgv {
		for _, v2 := range cgv[i+1:] {
			f(v1, v2)
		}
	}
}
