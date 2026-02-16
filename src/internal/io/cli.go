package io

import (
	"fmt"
	"queenables/src/internal/board"
)

func ReadBoardFromStdin() *board.Board {
	println("Enter board data (end with an empty line):")
	var rawBoardData string
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		if line == "" {
			break
		}
		rawBoardData += line + "\n"
	}

	boardData, err := board.ParseRawBoard(rawBoardData)
	if err != nil {
		println("Error parsing board data:", err.Error())
		return nil
	}

	b, err := board.NewBoard(boardData)
	if err != nil {
		println("Error creating board:", err.Error())
		return nil
	}

	return b
}
