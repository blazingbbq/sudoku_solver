package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNakedPair(t *testing.T) {
	c := &cell{
		candidates: []int{2, 5},
	}
	row := cellGroup{
		&cell{candidates: []int{2, 5}}, // Unaffected, part of the naked pair
		&cell{candidates: []int{3, 7}}, // Unaffected
		&cell{candidates: []int{4, 5}}, // 4
		&cell{candidates: []int{2, 4, 5, 6}}, // 4, 6
		&cell{candidates: []int{2, 3, 4, 5, 6, 7}}, // 3, 4, 6, 7
		&cell{value: 1}, // Unaffected
		&cell{value: 8}, // Unaffected
		&cell{value: 9}, // Unaffected
	}
	
	expected := cellGroup{
		&cell{candidates: []int{2, 5}}, // Unaffected, part of the naked pair
		&cell{candidates: []int{3, 7}}, // Unaffected
		&cell{candidates: []int{4}},    // 4
		&cell{candidates: []int{4, 6}}, // 4, 6
		&cell{candidates: []int{3, 4, 6, 7}}, // 3, 4, 6, 7
		&cell{value: 1}, // Unaffected
		&cell{value: 8}, // Unaffected
		&cell{value: 9}, // Unaffected
	}

	updated := updateNakedPairs(c, row)

	assert.True(t, updated)
	for i, cell := range row {
		assert.Equal(t, expected[i], cell)
	}

	// Should not update anything now that the naked pair is gone
	updated = updateNakedPairs(c, row)
	assert.False(t, updated)
}
