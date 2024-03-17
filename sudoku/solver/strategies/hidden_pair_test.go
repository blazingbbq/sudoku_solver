package strategies

import (
	"testing"

	"sudoku/sudoku/solver/grid"

	"github.com/stretchr/testify/assert"
)

func TestHiddenPair(t *testing.T) {
	c := grid.EmptyCell().WithCandidates(1, 2, 3, 4, 7, 9) // 2, 7
	row := grid.CellGroup{
		grid.EmptyCell().WithCandidates(2, 7, 8, 9), // 2, 7
		grid.EmptyCell().WithCandidates(3, 4, 8),    // Unaffected
		grid.EmptyCell().WithCandidates(3, 4, 8, 9), // Unaffected
		grid.EmptyCell().WithCandidates(3, 4, 8, 9), // Unaffected
		grid.EmptyCell().WithCandidates(3, 9),       // Unaffected
		grid.EmptyCell().WithCandidates(1, 3),       // Unaffected
		grid.EmptyCell().WithValue(5),               // Unaffected
		grid.EmptyCell().WithValue(6),               // Unaffected
	}

	cExpected := grid.EmptyCell().WithCandidates(2, 7) // 2, 7
	expected := grid.CellGroup{
		grid.EmptyCell().WithCandidates(2, 7),       // 2, 7
		grid.EmptyCell().WithCandidates(3, 4, 8),    // Unaffected
		grid.EmptyCell().WithCandidates(3, 4, 8, 9), // Unaffected
		grid.EmptyCell().WithCandidates(3, 4, 8, 9), // Unaffected
		grid.EmptyCell().WithCandidates(3, 9),       // Unaffected
		grid.EmptyCell().WithCandidates(1, 3),       // Unaffected
		grid.EmptyCell().WithValue(5),               // Unaffected
		grid.EmptyCell().WithValue(6),               // Unaffected
	}

	updated := UpdateHiddenPairs(c, row)

	assert.True(t, updated)
	assert.Equal(t, cExpected, c)
	for i, cell := range row {
		assert.Equal(t, expected[i], cell)
	}

	// Should not update anything now that the naked pair is gone
	updated = UpdateHiddenPairs(c, row)
	assert.False(t, updated)
}
