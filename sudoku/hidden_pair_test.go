package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHiddenPair(t *testing.T) {
	c := &cell{
		candidates: []int{1, 2, 3, 4, 7, 9}, // 2, 7
	}
	row := cellGroup{
		&cell{candidates: []int{2, 7, 8, 9}}, // 2, 7
		&cell{candidates: []int{3, 4, 8}},    // Unaffected
		&cell{candidates: []int{3, 4, 8, 9}}, // Unaffected
		&cell{candidates: []int{3, 4, 8, 9}}, // Unaffected
		&cell{candidates: []int{3, 9}},       // Unaffected
		&cell{candidates: []int{1, 3}},       // Unaffected
		&cell{value: 5},                      // Unaffected
		&cell{value: 6},                      // Unaffected
	}

	cExpected := &cell{
		candidates: []int{2, 7}, // 2, 7
	}
	expected := cellGroup{
		&cell{candidates: []int{2, 7}},       // 2, 7
		&cell{candidates: []int{3, 4, 8}},    // Unaffected
		&cell{candidates: []int{3, 4, 8, 9}}, // Unaffected
		&cell{candidates: []int{3, 4, 8, 9}}, // Unaffected
		&cell{candidates: []int{3, 9}},       // Unaffected
		&cell{candidates: []int{1, 3}},       // Unaffected
		&cell{value: 5},                      // Unaffected
		&cell{value: 6},                      // Unaffected
	}

	updated := updateHiddenPairs(c, row)

	assert.True(t, updated)
	assert.Equal(t, cExpected, c)
	for i, cell := range row {
		assert.Equal(t, expected[i], cell)
	}

	// Should not update anything now that the naked pair is gone
	updated = updateHiddenPairs(c, row)
	assert.False(t, updated)
}
