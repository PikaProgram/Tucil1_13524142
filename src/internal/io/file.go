package io

import (
	"os"
	"queenables/src/internal/board"
)

func ReadBoardFromFile(filePath string) (*board.Board, error) {
	rawBoardData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	boardData, err := board.ParseRawBoard(string(rawBoardData))
	if err != nil {
		return nil, err
	}

	return board.NewBoard(boardData)
}

func WriteBoardToFile(filePath string, b *board.Board) error {
	boardString := b.String()
	return os.WriteFile(filePath, []byte(boardString), 0644)
}
