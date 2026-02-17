package board

import (
	"errors"
	"sort"
	"sync"
	"time"
)

var iterationCount int
var elapsedTime time.Time
var lastTime time.Time

var solveStatsKey sync.RWMutex

func GetCurrentSolveStats() (iteration int, elapsed time.Duration) {
	solveStatsKey.RLock()
	defer solveStatsKey.RUnlock()
	return iterationCount, time.Since(elapsedTime)
}

func ResetSolveStats() {
	solveStatsKey.Lock()
	iterationCount = 0
	elapsedTime = time.Now()
	lastTime = time.Now()
	solveStatsKey.Unlock()
}

func addIteration() {
	solveStatsKey.Lock()
	iterationCount++
	solveStatsKey.Unlock()
}

func CreateSolvedBoard(b *Board) (*Board, error) {
	regionIDs := make([]string, 0, len(b.Regions))
	regions := make([][]*Cell, 0, len(b.Regions))
	ResetSolveStats()

	for id, cells := range b.Regions {
		regionIDs = append(regionIDs, id)
		regions = append(regions, cells)
	}

	sort.Strings(regionIDs)

	sols := make([]*Cell, len(regions))

	if FindSolutions(b, regions, 0, sols) {
		return b, nil
	}

	return nil, errors.New("No Solution Found")
}

func FindSolutions(b *Board, regs [][]*Cell, idx int, sols []*Cell) bool {
	if idx == len(regs) {
		for _, region := range regs {
			for _, cell := range region {
				cell.HasQueen = false
			}
		}

		for _, cell := range sols {
			cell.HasQueen = true
		}

		addIteration()

		if ValidateBoard(b) == nil {
			return true
		}
		return false
	}

	for _, cell := range regs[idx] {
		sols[idx] = cell
		if FindSolutions(b, regs, idx+1, sols) {
			return true
		}
	}

	return false
}

func CreateSolvedBoardBruteForce(b *Board) (*Board, error) {
	ResetSolveStats()

	queenPositions := make([][]bool, 0)
	for r := 0; r < b.Rows; r++ {
		row := make([]bool, b.Cols)
		queenPositions = append(queenPositions, row)
	}

	if FindSolutionsBruteForce(b, queenPositions, 0) {
		return b, nil
	}
	return nil, errors.New("No Solution Found")
}

func FindSolutionsBruteForce(b *Board, queenPositions [][]bool, row int) bool {
	if row == b.Rows {
		for r := 0; r < b.Rows; r++ {
			for c := 0; c < b.Cols; c++ {
				b.Cells[r][c].HasQueen = queenPositions[r][c]
			}
		}

		addIteration()
		if ValidateBoard(b) == nil {
			return true
		}
		return false
	}

	for c := 0; c < b.Cols; c++ {
		queenPositions[row][c] = true
		if FindSolutionsBruteForce(b, queenPositions, row+1) {
			return true
		}
		queenPositions[row][c] = false
	}
	return false
}
