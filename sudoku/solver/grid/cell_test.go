package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell(t *testing.T) {
	tests := []struct {
		name        string
		cell        *Cell
		expectOk    bool
		expectedVal int
	}{
		{
			name: "single value - bad",
			cell: &Cell{
				candidates: []int{1, 2, 3},
			},
			expectOk: false,
		},
		{
			name: "single value - ok",
			cell: &Cell{
				candidates: []int{1},
			},
			expectOk:    true,
			expectedVal: 1,
		},
		{
			name: "single value - empty",
			cell: &Cell{
				candidates: []int{},
			},
			expectOk: false,
		},
		{
			name: "single value - row",
			cell: &Cell{
				candidates: []int{1, 2, 3, 4},
				row: CellGroup{
					&Cell{value: 1},
					&Cell{value: 2},
					&Cell{candidates: []int{4}},
				},
			},
			expectOk:    true,
			expectedVal: 3,
		},
		{
			name: "single value - col",
			cell: &Cell{
				candidates: []int{1, 2, 3, 4},
				col: CellGroup{
					&Cell{value: 1},
					&Cell{value: 2},
					&Cell{candidates: []int{4}},
				},
			},
			expectOk:    true,
			expectedVal: 3,
		},
		{
			name: "single value - square",
			cell: &Cell{
				candidates: []int{1, 2, 3, 4},
				square: CellGroup{
					&Cell{value: 1},
					&Cell{value: 2},
					&Cell{candidates: []int{4}},
				},
			},
			expectOk:    true,
			expectedVal: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, ok := tt.cell.SingleValue()
			assert.Equal(t, tt.expectOk, ok)
			if tt.expectOk {
				assert.Equal(t, tt.expectedVal, value)
			}
		})
	}
}
