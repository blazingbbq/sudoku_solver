package strategies

import (
	"testing"

	"sudoku/sudoku/solver/grid"

	"github.com/stretchr/testify/assert"
)

func TestNakedPair(t *testing.T) {
	c := grid.EmptyCell().WithCandidates(2, 5) // 2, 5
	row := grid.CellGroup{
		grid.EmptyCell().WithCandidates(2, 5),             // Unaffected, part of the naked pair
		grid.EmptyCell().WithCandidates(3, 7),             // Unaffected
		grid.EmptyCell().WithCandidates(4, 5),             // 4
		grid.EmptyCell().WithCandidates(2, 4, 5, 6),       // 4, 6
		grid.EmptyCell().WithCandidates(2, 3, 4, 5, 6, 7), // 3, 4, 6, 7
		grid.EmptyCell().WithValue(1),                     // Unaffected
		grid.EmptyCell().WithValue(8),                     // Unaffected
		grid.EmptyCell().WithValue(9),                     // Unaffected
	}

	expected := grid.CellGroup{
		grid.EmptyCell().WithCandidates(2, 5),       // Unaffected, part of the naked pair
		grid.EmptyCell().WithCandidates(3, 7),       // Unaffected
		grid.EmptyCell().WithCandidates(4),          // 4
		grid.EmptyCell().WithCandidates(4, 6),       // 4, 6
		grid.EmptyCell().WithCandidates(3, 4, 6, 7), // 3, 4, 6, 7
		grid.EmptyCell().WithValue(1),               // Unaffected
		grid.EmptyCell().WithValue(8),               // Unaffected
		grid.EmptyCell().WithValue(9),               // Unaffected
	}

	updated := UpdateNakedPairs(c, row)

	assert.True(t, updated)
	for i, cell := range row {
		assert.Equal(t, expected[i], cell)
	}

	// Should not update anything now that the naked pair is gone
	updated = UpdateNakedPairs(c, row)
	assert.False(t, updated)
}
