package board

import "errors"

type Cell struct {
	Row      int
	Col      int
	HasQueen bool
	RegionID string
}

type Board struct {
	Rows        int
	Cols        int
	Cells       [][]*Cell
	Adjacency   [][][]*Cell
	Regions     map[string][]*Cell
	RegionCount int
}

func NewBoard(data [][]string) (*Board, error) {
	rows := len(data)
	cols := len(data[0])

	if rows == 0 || cols == 0 {
		return nil, errors.New("Board data cannot be empty")
	}

	if rows != cols {
		return nil, errors.New("Board must be square (rows must equal columns)")
	}

	cells := make([][]*Cell, rows)
	adjacency := make([][][]*Cell, rows)

	for r := 0; r < rows; r++ {
		cells[r] = make([]*Cell, cols)
		for c := 0; c < cols; c++ {
			cell := &Cell{Row: r, Col: c, HasQueen: false, RegionID: data[r][c]}
			cells[r][c] = cell
		}
	}

	regions := make(map[string][]*Cell)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			regionID := cells[r][c].RegionID
			regions[regionID] = append(regions[regionID], cells[r][c])
		}
	}

	for r := 0; r < rows; r++ {
		adjacency[r] = make([][]*Cell, cols)
		for c := 0; c < cols; c++ {
			var adjCells []*Cell
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					nr, nc := r+dr, c+dc
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
						adjCells = append(adjCells, cells[nr][nc])
					}
				}
			}
			adjacency[r][c] = adjCells
		}
	}

	return &Board{
		Rows:        rows,
		Cols:        cols,
		Adjacency:   adjacency,
		Cells:       cells,
		Regions:     regions,
		RegionCount: len(regions),
	}, nil
}

func ParseRawBoard(raw string) ([][]string, error) {
	lines := []string{}
	for _, line := range splitAndTrim(raw) {
		if line != "" {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		return nil, errors.New("Raw board data cannot be empty")
	}

	boardData := make([][]string, len(lines))
	for i, line := range lines {
		boardData[i] = splitToChars(line)
	}

	return boardData, nil
}

func splitAndTrim(s string) []string {
	var result []string
	current := ""
	for _, r := range s {
		if r == '\n' {
			result = append(result, current)
			current = ""
		} else if r != '\r' {
			current += string(r)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

func splitToChars(s string) []string {
	var chars []string
	for _, r := range s {
		chars = append(chars, string(r))
	}
	return chars
}
