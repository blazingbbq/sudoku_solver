package grid

type Cell struct {
	value       int
	rowIndex    int
	colIndex    int
	sqrRowIndex int
	sqrColIndex int

	row    CellGroup
	col    CellGroup
	square CellGroup
	region CellGroup

	candidates CellGroupValues
}

// NewCell creates a new cell
func NewCell(row, col int) *Cell {
	return &Cell{
		rowIndex:    row,
		colIndex:    col,
		sqrRowIndex: row / 3,
		sqrColIndex: col / 3,
	}
}

// EmptyCell creates a new uninitialized cell.
//
// To be used for testing purposes.
func EmptyCell() *Cell {
	return &Cell{}
}

// WithCandidates sets the candidates of the cell and returns the cell
func (c *Cell) WithCandidates(candidates ...int) *Cell {
	c.candidates = candidates
	return c
}

// WithValue sets the value of the cell and returns the cell
func (c *Cell) WithValue(value int) *Cell {
	c.value = value
	return c
}

// SetValue sets the value of the cell
func (c *Cell) SetValue(value int) {
	c.value = value
}

// SetCandidates sets the candidates of the cell
func (c *Cell) SetCandidates(candidates CellGroupValues) {
	c.candidates = candidates
}

// Candidates returns the candidates of the cell
func (c *Cell) Candidates() CellGroupValues {
	return c.candidates
}

// RowIndex returns the row index of the cell
func (c *Cell) RowIndex() int {
	return c.rowIndex
}

// ColIndex returns the column index of the cell
func (c *Cell) ColIndex() int {
	return c.colIndex
}

// SqrRowIndex returns the square row index of the cell
func (c *Cell) SqrRowIndex() int {
	return c.sqrRowIndex
}

// SqrColIndex returns the square column index of the cell
func (c *Cell) SqrColIndex() int {
	return c.sqrColIndex
}

// ForEachRegion calls the given function for each region of the cell
func (c *Cell) ForEachRegion(f func(CellGroup)) {
	f(c.row)
	f(c.col)
	f(c.square)
}

// GetValue returns the value of the cell and a boolean indicating if it is known
func (c *Cell) GetValue() (int, bool) {
	return c.value, c.value != 0
}

// SingleValue returns the single candidate value of the cell and a boolean indicating if it is known
func (c *Cell) SingleValue() (int, bool) {
	// If the cell has only one candidate, return it
	if len(c.candidates) == 1 {
		return c.candidates[0], true
	}

	// If a candidate is only present in one region, return it
	val, found := 0, false
	c.ForEachRegion(func(region CellGroup) {
		if vals := c.candidates.Subtract(region.KnownAndCandidateValues()); len(vals) == 1 {
			val = vals[0]
			found = true
		}
	})
	return val, found
}
