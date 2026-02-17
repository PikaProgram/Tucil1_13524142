package solver

import (
	"queenables/src/internal/board"
	"time"

	"github.com/huandu/go-clone"
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

type SolveProgress struct {
	Board          *board.Board
	IterationCount int
	ElapsedTime    time.Duration
	IsComplete     bool
	Err            error
	status         string
}

func SolveBoardAsync(b *board.Board, progressChan chan<- SolveProgress, updateDuration time.Duration, algo int) {
	if updateDuration <= 0 {
		updateDuration = 5_000 * time.Millisecond
	}

	go func() {
		defer close(progressChan)

		board.ResetSolveStats()

		completed := make(chan error, 1)

		go func() {
			switch algo {
			case 0:
				_, err := board.CreateSolvedBoard(b)
				completed <- err
			}
		}()

		ticker := time.NewTicker(updateDuration)
		defer ticker.Stop()

		for {
			select {
			case err := <-completed:
				iterationCount, elapsed := board.GetCurrentSolveStats()
				progressBoard := clone.Clone(b).(*board.Board)
				progressChan <- SolveProgress{
					Board:          progressBoard,
					IterationCount: iterationCount,
					ElapsedTime:    elapsed,
					IsComplete:     true,
					Err:            err,
					status:         "Complete",
				}
				if err != nil {
					progressChan <- SolveProgress{
						Board:          progressBoard,
						IterationCount: iterationCount,
						ElapsedTime:    elapsed,
						IsComplete:     true,
						Err:            err,
						status:         "Error",
					}
					return
				}
				return
			case <-ticker.C:
				iterationCount, elapsed := board.GetCurrentSolveStats()
				progressBoard := clone.Clone(b).(*board.Board)
				progressChan <- SolveProgress{
					Board:          progressBoard,
					IterationCount: iterationCount,
					ElapsedTime:    elapsed,
					IsComplete:     false,
					Err:            nil,
					status:         "Solving",
				}
			}
		}
	}()
}
