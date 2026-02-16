package board

func DisplayBoard(b *Board) {
	println(BoardToString(b))
}

func BoardToString(b *Board) string {
	var boardString string
	for r := 0; r < b.Rows; r++ {
		boardString += GenerateSeparator(b)
		for c := 0; c < b.Cols; c++ {
			cell := b.Cells[r][c]
			if cell.HasQueen {
				boardString += "| @ "
			} else {
				boardString += "| " + cell.RegionID + " "
			}
		}
		boardString += "|\n"
	}
	boardString += GenerateSeparator(b)
	return boardString
}

func GenerateSeparator(b *Board) string {
	var separator string
	for range b.Cells[0] {
		separator += "+---"
	}
	separator += "+\n"
	return separator
}
