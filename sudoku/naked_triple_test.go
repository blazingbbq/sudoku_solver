package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNakedTriple(t *testing.T) {
	c := &cell{
		candidates: []int{3, 4, 5},
	}
	row := cellGroup{
		&cell{candidates: []int{3, 4, 5}},       // Unaffected, part of the naked triple
		&cell{candidates: []int{3, 5}},          // Unaffected, part of the naked triple
		&cell{candidates: []int{2, 4, 5, 6, 9}}, // 2, 6, 9
		&cell{candidates: []int{2, 5, 9}},       // 2, 9
		&cell{candidates: []int{4, 5, 6, 9}},    // 6, 9
		&cell{value: 1},                         // Unaffected
		&cell{value: 7},                         // Unaffected
		&cell{value: 8},                         // Unaffected
	}

	expected := cellGroup{
		&cell{candidates: []int{3, 4, 5}}, // Unaffected, part of the naked triple
		&cell{candidates: []int{3, 5}},    // Unaffected, part of the naked triple
		&cell{candidates: []int{2, 6, 9}}, // 2, 6, 9
		&cell{candidates: []int{2, 9}},    // 2, 9
		&cell{candidates: []int{6, 9}},    // 6, 9
		&cell{value: 1},                   // Unaffected
		&cell{value: 7},                   // Unaffected
		&cell{value: 8},                   // Unaffected
	}

	updated := updateNakedTriples(c, row)

	assert.True(t, updated)
	for i, cell := range row {
		assert.Equal(t, expected[i], cell)
	}

	// Should not update anything now that the naked triple is gone
	updated = updateNakedTriples(c, row)
	assert.False(t, updated)
}
