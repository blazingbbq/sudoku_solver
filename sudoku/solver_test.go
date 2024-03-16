package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	tests := []struct {
		name      string
		sudoku    *Sudoku
		expected  *Sudoku
		expectErr bool
	}{
		{
			name: "solved",
			sudoku: &Sudoku{
				board: [9][9]int{
					{5, 3, 4, 6, 7, 8, 9, 1, 2},
					{6, 7, 2, 1, 9, 5, 3, 4, 8},
					{1, 9, 8, 3, 4, 2, 5, 6, 7},
					{8, 5, 9, 7, 6, 1, 4, 2, 3},
					{4, 2, 6, 8, 5, 3, 7, 9, 1},
					{7, 1, 3, 9, 2, 4, 8, 5, 6},
					{9, 6, 1, 5, 3, 7, 2, 8, 4},
					{2, 8, 7, 4, 1, 9, 6, 3, 5},
					{3, 4, 5, 2, 8, 6, 1, 7, 9},
				},
			},
			expected: &Sudoku{
				board: [9][9]int{
					{5, 3, 4, 6, 7, 8, 9, 1, 2},
					{6, 7, 2, 1, 9, 5, 3, 4, 8},
					{1, 9, 8, 3, 4, 2, 5, 6, 7},
					{8, 5, 9, 7, 6, 1, 4, 2, 3},
					{4, 2, 6, 8, 5, 3, 7, 9, 1},
					{7, 1, 3, 9, 2, 4, 8, 5, 6},
					{9, 6, 1, 5, 3, 7, 2, 8, 4},
					{2, 8, 7, 4, 1, 9, 6, 3, 5},
					{3, 4, 5, 2, 8, 6, 1, 7, 9},
				},
			},
			expectErr: false,
		},
		{
			name: "unsolvable",
			sudoku: &Sudoku{
				board: [9][9]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 2},
					{0, 0, 0, 0, 0, 0, 0, 0, 8},
					{0, 0, 0, 0, 0, 0, 0, 0, 7},
					{0, 0, 0, 0, 0, 0, 0, 0, 3},
					{0, 0, 0, 0, 0, 0, 0, 0, 1},
					{0, 0, 0, 0, 0, 0, 0, 0, 6},
					{0, 0, 0, 0, 0, 0, 0, 0, 4},
					{0, 0, 0, 0, 0, 0, 0, 0, 5},
					{0, 0, 0, 0, 0, 0, 0, 0, 9},
				},
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name: "invalid",
			sudoku: &Sudoku{
				board: [9][9]int{
					{5, 5, 4, 6, 7, 8, 9, 1, 2},
					{5, 7, 2, 1, 9, 5, 3, 4, 8},
					{1, 9, 8, 3, 4, 2, 5, 6, 7},
					{8, 5, 9, 7, 6, 1, 4, 2, 3},
					{4, 2, 6, 8, 5, 3, 7, 9, 1},
					{7, 1, 3, 9, 2, 4, 8, 5, 6},
					{9, 6, 1, 5, 3, 7, 2, 8, 4},
					{2, 8, 7, 4, 1, 9, 6, 3, 5},
					{3, 4, 5, 2, 8, 6, 1, 7, 9},
				},
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name: "solveable - Easy",
			sudoku: &Sudoku{
				board: [9][9]int{
					{0, 8, 9, 1, 0, 0, 0, 0, 7},
					{6, 0, 0, 2, 0, 0, 1, 0, 8},
					{0, 7, 0, 0, 0, 0, 5, 0, 0},
					{0, 0, 0, 8, 6, 0, 0, 0, 0},
					{0, 0, 0, 7, 4, 3, 0, 0, 0},
					{0, 0, 0, 0, 0, 1, 0, 0, 0},
					{1, 0, 6, 9, 0, 0, 0, 7, 0},
					{3, 0, 7, 0, 0, 2, 0, 0, 4},
					{4, 9, 0, 0, 0, 0, 0, 6, 0},
				},
			},
			expected: &Sudoku{
				board: [9][9]int{
					{5, 8, 9, 1, 3, 6, 4, 2, 7},
					{6, 3, 4, 2, 7, 5, 1, 9, 8},
					{2, 7, 1, 4, 9, 8, 5, 3, 6},
					{7, 4, 2, 8, 6, 9, 3, 5, 1},
					{9, 1, 5, 7, 4, 3, 6, 8, 2},
					{8, 6, 3, 5, 2, 1, 7, 4, 9},
					{1, 2, 6, 9, 5, 4, 8, 7, 3},
					{3, 5, 7, 6, 8, 2, 9, 1, 4},
					{4, 9, 8, 3, 1, 7, 2, 6, 5},
				},
			},
			expectErr: false,
		},
		{
			name: "solveable - Medium",
			sudoku: &Sudoku{
				board: [9][9]int{
					{4, 0, 0, 0, 1, 0, 8, 0, 6},
					{0, 1, 0, 0, 7, 0, 0, 0, 0},
					{7, 8, 0, 0, 0, 9, 0, 0, 0},
					{5, 0, 0, 0, 0, 1, 0, 0, 0},
					{0, 0, 1, 0, 4, 0, 3, 0, 0},
					{0, 0, 0, 8, 6, 0, 0, 0, 5},
					{0, 0, 0, 1, 0, 0, 0, 0, 2},
					{0, 0, 0, 0, 9, 0, 0, 3, 0},
					{9, 0, 2, 0, 3, 0, 0, 0, 8},
				},
			},
			expected: &Sudoku{
				board: [9][9]int{
					{4, 5, 9, 2, 1, 3, 8, 7, 6},
					{2, 1, 6, 4, 7, 8, 9, 5, 3},
					{7, 8, 3, 6, 5, 9, 2, 4, 1},
					{5, 6, 7, 3, 2, 1, 4, 8, 9},
					{8, 2, 1, 9, 4, 5, 3, 6, 7},
					{3, 9, 4, 8, 6, 7, 1, 2, 5},
					{6, 3, 5, 1, 8, 4, 7, 9, 2},
					{1, 7, 8, 5, 9, 2, 6, 3, 4},
					{9, 4, 2, 7, 3, 6, 5, 1, 8},
				},
			},
			expectErr: false,
		},
		{
			name: "solveable - Medium 2",
			sudoku: &Sudoku{
				board: [9][9]int{
					{0, 1, 0, 0, 7, 5, 0, 3, 0},
					{0, 0, 2, 0, 0, 0, 8, 0, 0},
					{0, 0, 3, 0, 0, 0, 0, 0, 0},
					{0, 0, 5, 6, 8, 0, 2, 0, 0},
					{0, 7, 0, 0, 1, 0, 0, 0, 0},
					{0, 0, 9, 0, 0, 2, 3, 0, 0},
					{0, 0, 0, 0, 0, 8, 5, 0, 0},
					{3, 0, 1, 0, 0, 0, 7, 0, 0},
					{0, 8, 0, 7, 9, 0, 0, 6, 0},
				},
			},
			expected: &Sudoku{
				board: [9][9]int{
					{8, 1, 6, 2, 7, 5, 9, 3, 4},
					{7, 4, 2, 9, 3, 6, 8, 1, 5},
					{9, 5, 3, 8, 4, 1, 6, 2, 7},
					{4, 3, 5, 6, 8, 7, 2, 9, 1},
					{2, 7, 8, 3, 1, 9, 4, 5, 6},
					{1, 6, 9, 4, 5, 2, 3, 7, 8},
					{6, 9, 7, 1, 2, 8, 5, 4, 3},
					{3, 2, 1, 5, 6, 4, 7, 8, 9},
					{5, 8, 4, 7, 9, 3, 1, 6, 2},
				},
			},
			expectErr: false,
		},
		{
			name: "solveable - Medium 3",
			sudoku: &Sudoku{
				board: [9][9]int{
					{8, 3, 2, 1, 0, 0, 4, 0, 0},
					{7, 0, 0, 0, 0, 9, 2, 1, 0},
					{1, 0, 0, 0, 4, 8, 0, 5, 0},
					{0, 8, 0, 0, 0, 7, 0, 3, 1},
					{6, 1, 0, 0, 0, 4, 0, 0, 0},
					{2, 0, 0, 8, 0, 0, 0, 4, 0},
					{0, 0, 9, 0, 0, 0, 0, 8, 5},
					{0, 0, 0, 0, 0, 0, 0, 6, 0},
					{0, 0, 0, 0, 0, 2, 3, 7, 0},
				},
			},
			expected: &Sudoku{
				board: [9][9]int{
				},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := NewSolver(tt.sudoku)
			solution, err := solver.Solve()

			if tt.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, solution)
			}
		})
	}
}

func TestCell(t *testing.T) {
	tests := []struct {
		name        string
		cell        *cell
		expectOk    bool
		expectedVal int
	}{
		{
			name: "single value - bad",
			cell: &cell{
				candidates: []int{1, 2, 3},
			},
			expectOk: false,
		},
		{
			name: "single value - ok",
			cell: &cell{
				candidates: []int{1},
			},
			expectOk:    true,
			expectedVal: 1,
		},
		{
			name: "single value - empty",
			cell: &cell{
				candidates: []int{},
			},
			expectOk: false,
		},
		{
			name: "single value - row",
			cell: &cell{
				candidates: []int{1, 2, 3, 4},
				row: cellGroup{
					&cell{value: 1},
					&cell{value: 2},
					&cell{candidates: []int{4}},
				},
			},
			expectOk:    true,
			expectedVal: 3,
		},
		{
			name: "single value - col",
			cell: &cell{
				candidates: []int{1, 2, 3, 4},
				col: cellGroup{
					&cell{value: 1},
					&cell{value: 2},
					&cell{candidates: []int{4}},
				},
			},
			expectOk:    true,
			expectedVal: 3,
		},
		{
			name: "single value - square",
			cell: &cell{
				candidates: []int{1, 2, 3, 4},
				square: cellGroup{
					&cell{value: 1},
					&cell{value: 2},
					&cell{candidates: []int{4}},
				},
			},
			expectOk:    true,
			expectedVal: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, ok := tt.cell.singleValue()
			assert.Equal(t, tt.expectOk, ok)
			if tt.expectOk {
				assert.Equal(t, tt.expectedVal, value)
			}
		})
	}
}

func TestSingleValueOnFullBoard(t *testing.T) {
	sudoku := &Sudoku{
		board: [9][9]int{
			{4, 0, 0, 0, 0, 0, 8, 0, 6},
			{0, 1, 0, 0, 7, 0, 0, 0, 0},
			{7, 8, 0, 0, 0, 9, 0, 0, 0},
			{5, 0, 0, 0, 0, 1, 0, 0, 0},
			{0, 0, 1, 0, 4, 0, 3, 0, 0},
			{0, 0, 0, 8, 6, 0, 0, 0, 5},
			{0, 0, 0, 1, 0, 0, 0, 0, 2},
			{0, 0, 0, 0, 9, 0, 0, 3, 0},
			{9, 0, 2, 0, 3, 0, 0, 0, 8},
		},
	}
	solver := NewSolver(sudoku)

	v, ok := solver.grid[0][7].singleValue()
	assert.True(t, ok)
	assert.Equal(t, 7, v)

	v, ok = solver.grid[3][2].singleValue()
	assert.False(t, ok)
}
