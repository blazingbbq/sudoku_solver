package grid

type CellGroupValues []int

// Append returns a new CellGroupValues with value appended
func (cgv CellGroupValues) Append(value ...int) CellGroupValues {
	return append(cgv, value...)
}

// Contains returns true if value is in cgv
func (cgv CellGroupValues) Contains(value int) bool {
	for _, v := range cgv {
		if v == value {
			return true
		}
	}
	return false
}

// Subtract returns the values in cgv that are not in sub
func (cgv CellGroupValues) Subtract(sub CellGroupValues) CellGroupValues {
	if cgv == nil {
		return nil
	}

	result := CellGroupValues{}
	for _, v := range cgv {
		if !sub.Contains(v) {
			result = result.Append(v)
		}
	}
	return result
}

// Equals returns true if cgv and other contain the same values
func (cgv CellGroupValues) Equals(other CellGroupValues) bool {
	if len(cgv) != len(other) {
		return false
	}
	for _, v := range cgv {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// EqualOrSubsetOf returns true if cgv is a subset of other
func (cgv CellGroupValues) EqualOrSubsetOf(other CellGroupValues) bool {
	for _, v := range cgv {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// Count returns the number of times value appears in cgv
func (cgv CellGroupValues) Count(value int) int {
	count := 0
	for _, v := range cgv {
		if v == value {
			count++
		}
	}
	return count
}

// Length returns the number of values in cgv
func (cgv CellGroupValues) Length() int {
	return len(cgv)
}

// Transform returns a new CellGroupValues with the result of f for each value in cgv
func (cgv CellGroupValues) Transform(f func(int) (int, bool)) CellGroupValues {
	result := CellGroupValues{}
	for _, v := range cgv {
		if newVal, ok := f(v); ok {
			result = result.Append(newVal)
		}
	}
	return result
}

// Filter returns a new CellGroupValues with values for which f returns true
func (cgv CellGroupValues) Filter(f func(int) bool) CellGroupValues {
	result := CellGroupValues{}
	for _, v := range cgv {
		if f(v) {
			result = result.Append(v)
		}
	}
	return result
}

// ForEach calls f for each value in cgv
func (cgv CellGroupValues) ForEach(f func(int)) {
	for _, v := range cgv {
		f(v)
	}
}

// ForEachUniquePair calls f for each unique pair of values in cgv
func (cgv CellGroupValues) ForEachUniquePair(f func(int, int)) {
	for i, v1 := range cgv {
		for _, v2 := range cgv[i+1:] {
			f(v1, v2)
		}
	}
}
