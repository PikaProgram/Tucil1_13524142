package solver

import (
	"queenables/src/internal/board"
	"time"
)

type SolveProgress struct {
	Board          *board.Board
	IterationCount int
	ElapsedTime    time.Duration
	IsComplete     bool
	Err            error
}

func SolveBoardAsync(b *board.Board, progressChan chan<- SolveProgress, updateDuration time.Duration) {
	if updateDuration <= 0 {
		updateDuration = 5_000 * time.Millisecond
	}

	go func() {
		defer close(progressChan)

		completed := make(chan error, 1)

		go func() {
			_, err := board.CreateSolvedBoard(b)
			completed <- err
		}()

		ticker := time.NewTicker(updateDuration)
		defer ticker.Stop()

		for {
			select {
			case err := <-completed:
				iterationCount, elapsed := board.GetCurrentSolveStats()
				progressChan <- SolveProgress{
					Board:          b,
					IterationCount: iterationCount,
					ElapsedTime:    elapsed,
					IsComplete:     true,
					Err:            err,
				}
				if err != nil {
					println("Error solving board:", err.Error())
				}
				return
			case <-ticker.C:
				iterationCount, elapsed := board.GetCurrentSolveStats()
				progressChan <- SolveProgress{
					Board:          nil,
					IterationCount: iterationCount,
					ElapsedTime:    elapsed,
					IsComplete:     false,
					Err:            nil,
				}
			}
		}
	}()
}
