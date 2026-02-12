package board

func DisplayBoard(b *Board) {
	for r := 0; r < b.Rows; r++ {
		PrintHorizontalSeparator(b)
		for c := 0; c < b.Cols; c++ {
			cell := b.Cells[r][c]
			if cell.HasQueen {
				print("| Q ")
			} else {
				print("| " + cell.RegionID + " ")
			}
		}
		println("|")
	}
	PrintHorizontalSeparator(b)
}

func PrintHorizontalSeparator(b *Board) {
	for range b.Cells[0] {
		print("+---")
	}
	println("+")
}
