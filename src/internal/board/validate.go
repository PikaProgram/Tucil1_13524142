package board

import (
	"errors"
	"fmt"
)

func ValidateBoard(b *Board) error {
	// Check if each region has exactly one queen
	for regionID, cells := range b.Regions {
		queenCount := 0
		for _, cell := range cells {
			if cell.HasQueen {
				queenCount++
			}
		}
		if queenCount != 1 {
			return errors.New("Region " + regionID + " must have exactly one queen")
		}
	}

	// Check for adjacent queens
	for r := 0; r < b.Rows; r++ {
		for c := 0; c < b.Cols; c++ {
			cell := b.Cells[r][c]
			for _, regionCells := range b.Regions {
				for _, rc := range regionCells {
					if rc.Row == r && rc.Col == c {
						cell.HasQueen = rc.HasQueen
						break
					}
				}
				if cell.HasQueen {
					break
				}
			}
			if cell.HasQueen {
				adjacentCells := b.Adjacency[r][c]
				for _, adj := range adjacentCells {
					if adj.HasQueen {
						return errors.New("Adjacent queens found at (" + fmt.Sprint(r) + "," + fmt.Sprint(c) + ") and (" + fmt.Sprint(adj.Row) + "," + fmt.Sprint(adj.Col) + ")")
					}
				}
			}
		}
	}

	// Check rows and columns for multiple queens
	for r := 0; r < b.Rows; r++ {
		queenCount := 0
		for c := 0; c < b.Cols; c++ {
			if b.Cells[r][c].HasQueen {
				queenCount++
			}
		}
		if queenCount > 1 {
			return errors.New("Row " + fmt.Sprint(r) + " has more than one queen")
		}
	}

	for c := 0; c < b.Cols; c++ {
		queenCount := 0
		for r := 0; r < b.Rows; r++ {
			if b.Cells[r][c].HasQueen {
				queenCount++
			}
		}
		if queenCount > 1 {
			return errors.New("Column " + fmt.Sprint(c) + " has more than one queen")
		}
	}

	return nil
}
