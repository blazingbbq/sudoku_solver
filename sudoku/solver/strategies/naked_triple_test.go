package strategies

import (
	"testing"

	"sudoku/sudoku/solver/grid"

	"github.com/stretchr/testify/assert"
)

func TestNakedTriple(t *testing.T) {
	c := grid.EmptyCell().WithCandidates(3, 4, 5) // 3, 4, 5
	row := grid.CellGroup{
		grid.EmptyCell().WithCandidates(3, 4, 5),       // Unaffected, part of the naked triple
		grid.EmptyCell().WithCandidates(3, 5),          // Unaffected, part of the naked triple
		grid.EmptyCell().WithCandidates(2, 4, 5, 6, 9), // 2, 6, 9
		grid.EmptyCell().WithCandidates(2, 5, 9),       // 2, 9
		grid.EmptyCell().WithCandidates(4, 5, 6, 9),    // 6, 9
		grid.EmptyCell().WithValue(1),                  // Unaffected
		grid.EmptyCell().WithValue(7),                  // Unaffected
		grid.EmptyCell().WithValue(8),                  // Unaffected
	}

	expected := grid.CellGroup{
		grid.EmptyCell().WithCandidates(3, 4, 5), // Unaffected, part of the naked triple
		grid.EmptyCell().WithCandidates(3, 5),    // Unaffected, part of the naked triple
		grid.EmptyCell().WithCandidates(2, 6, 9), // 2, 6, 9
		grid.EmptyCell().WithCandidates(2, 9),    // 2, 9
		grid.EmptyCell().WithCandidates(6, 9),    // 6, 9
		grid.EmptyCell().WithValue(1),            // Unaffected
		grid.EmptyCell().WithValue(7),            // Unaffected
		grid.EmptyCell().WithValue(8),            // Unaffected
	}

	updated := UpdateNakedTriples(c, row)

	assert.True(t, updated)
	for i, cell := range row {
		assert.Equal(t, expected[i], cell)
	}

	// Should not update anything now that the naked triple is gone
	updated = UpdateNakedTriples(c, row)
	assert.False(t, updated)
}
