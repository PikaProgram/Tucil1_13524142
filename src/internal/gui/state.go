package gui

import (
	"queenables/src/internal/board"
	"queenables/src/internal/gui/solver"
)

type GUIState struct {
	Board             *board.Board
	SolveState        solver.SolveState
	SolveProgressChan chan solver.SolveProgress
}
