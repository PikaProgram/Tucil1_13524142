package main

import (
	"tucil1/src/internal/board"
)

func main() {
	rawBoardData :=
		`AAAAB
ACCBB
AACBB
DDCEE
DDCEE`

	boardData, err := board.ParseRawBoard(rawBoardData)
	if err != nil {
		println("Error parsing raw board data:", err.Error())
		return
	}

	b, err := board.NewBoard(boardData)
	println("Board created with dimensions:", b.Rows, "x", b.Cols)

	board.DisplayBoard(b)

	solvedBoard, err := board.CreateSolvedBoard(b)
	if err != nil {
		println("Error solving board:", err.Error())
		return
	}

	println()
	println("Solved Board:")
	board.DisplayBoard(solvedBoard)

	err = board.ValidateBoard(solvedBoard)
	if err != nil {
		println("Board validation failed:", err.Error())
	} else {
		println("Board is valid!")
	}
}
