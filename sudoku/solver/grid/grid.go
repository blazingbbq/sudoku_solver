package grid

type Grid [9][9]*Cell

// ForEachCell calls the given function for each cell of the grid
func (g *Grid) ForEachCell(f func(c *Cell)) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			f(g[i][j])
		}
	}
}

// AssociateCells associates each cell with its row, column, square and region
func (g *Grid) AssociateCells() {
	g.ForEachCell(func(c *Cell) {
		g.associateCell(c.rowIndex, c.colIndex)
	})
}

func (g *Grid) associateCell(row, col int) {
	c := g[row][col]

	for i := 0; i < 9; i++ {
		if i != col {
			c.row = append(c.row, g[row][i])
		}
		if i != row {
			c.col = append(c.col, g[i][col])
		}
		if i != (row%3)*3+col%3 {
			c.square = append(c.square, g[c.sqrRowIndex*3+i/3][c.sqrColIndex*3+i%3])
		}
	}

	c.region = append(c.region, c.row...)
	c.region = append(c.region, c.col...)
	c.region = append(c.region, c.square...)
}
