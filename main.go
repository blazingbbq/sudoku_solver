package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var lines []string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter sudoku:\n")
	for i := 0; i < 9; i++ {
		text, _ := reader.ReadString('\n')
		lines = append(lines, text)
	}

	sudoku, err := NewSudoku().ReadFromStrings(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Sudoku: \n", sudoku.String(), "\n")

	solution, err := sudoku.Solve()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Solution: \n", solution.String(), "\n")
}
