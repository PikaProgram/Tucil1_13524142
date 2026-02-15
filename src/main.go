package main

import (
	"tucil1/src/internal/board"
)

func main() {
	rawBoardData :=
		`ABCDEFGHIJKLM
BCDEFGHIJKLMA
CDEFGHIJKLMAB
DEFGHIJKLMABC
EFGHIJKLMABCD
FGHIJKLMABCDE
GHIJKLMABCDEF
HIJKLMABCDEFG
IJKLMABCDEFGH
JKLMABCDEFGHI
KLMABCDEFGHIJ
LMABCDEFGHIJK
MABCDEFGHIJKL
`

	boardData, err := board.ParseRawBoard(rawBoardData)
	if err != nil {
		println("Error parsing raw board data:", err.Error())
		return
	}

	b, err := board.NewBoard(boardData)
	println("Board created with dimensions:", b.Rows, "x", b.Cols)

	println("Initial Board:")
	board.DisplayBoard(b)

	solvedBoard, err := board.CreateSolvedBoard(b)
	if err != nil {
		println("Error solving board:", err.Error())
		return
	}

	println("Solved Board:")
	board.DisplayBoard(solvedBoard)
}
