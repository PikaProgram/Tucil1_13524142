package board

import (
	"errors"
	"sort"
	"strconv"
	"time"
)

var iterationCount int
var elapsedTime time.Time
var lastTime time.Time

func CreateSolvedBoard(b *Board) (*Board, error) {
	regionIDs := make([]string, 0, len(b.Regions))
	regions := make([][]*Cell, 0, len(b.Regions))
	iterationCount = 0
	elapsedTime = time.Now()
	lastTime = time.Now()

	for id, cells := range b.Regions {
		regionIDs = append(regionIDs, id)
		regions = append(regions, cells)
	}

	sort.Strings(regionIDs)

	sols := make([]*Cell, len(regions))

	if FindSolutions(b, regions, 0, sols) {
		println("Finished " + strconv.FormatInt(int64(iterationCount), 10) + " iterations.")
		println("Total Time Elapsed: " + time.Since(elapsedTime).String())
		return b, nil
	}

	println("Finished " + strconv.FormatInt(int64(iterationCount), 10) + " iterations.")
	println("Total Time Elapsed: " + time.Since(elapsedTime).String())
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

		iterationCount++

		if time.Since(lastTime) > 3*time.Second {
			println("Iteration #" + strconv.FormatInt(int64(iterationCount), 10) + " :")
			DisplayBoard(b)
			println("Average Iteration Per Second: " + strconv.FormatFloat(float64(iterationCount)/time.Since(elapsedTime).Seconds(), 'f', 2, 64) + "/s")
			println()
			lastTime = time.Now()
		}

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
