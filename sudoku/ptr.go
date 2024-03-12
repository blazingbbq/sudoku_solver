package sudoku

func ptr[K any](val K) *K {
	return &val
}
