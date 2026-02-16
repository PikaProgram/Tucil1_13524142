package gui

import (
	"queenables/src/internal/board"
	"queenables/src/internal/gui/solver"
)

type SolveState int

const (
	SolveStateIdle = iota
	SolveStateSolving
	SolveStateSolved
	SolveStateError
)

var SolveStateName = map[SolveState]string{
	SolveStateIdle:    "Idle",
	SolveStateSolving: "Solving",
	SolveStateSolved:  "Solved",
	SolveStateError:   "Error",
}

type GUIState struct {
	Board             *board.Board
	SolveState        SolveState
	SolveProgressChan chan solver.SolveProgress
}
