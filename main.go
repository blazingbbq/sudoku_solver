package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"sudoku/sudoku"
	"sudoku/sudoku/solver"
)

func main() {
	var lines []string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter sudoku:\n")
	for i := 0; i < 9; i++ {
		text, _ := reader.ReadString('\n')
		lines = append(lines, text)
	}

	sudoku, err := sudoku.NewSudoku([9][9]int{}).ReadFromStrings(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Sudoku: \n", sudoku.String(), "\n")

	s := solver.NewSolver(sudoku)
	solution, err := s.Solve()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Solution: \n", solution.String(), "\n")
}
