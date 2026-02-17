package widgets

import (
	"image"
	"image/color"
	"queenables/src/internal/board"
	"strings"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type BoardWidget struct {
	board *board.Board

	cellSize     int32
	borderSize   int32
	regionColors map[string]color.NRGBA
}

func CreateBoardWidget(board *board.Board) *BoardWidget {
	regionColors := GenerateRegionColors(strings.Split("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", ""))
	return &BoardWidget{
		board:        board,
		cellSize:     48,
		borderSize:   2,
		regionColors: regionColors,
	}
}

func (bw *BoardWidget) SetBoard(board *board.Board) {
	bw.board = board
}

func (bw *BoardWidget) Layout(gtx layout.Context) layout.Dimensions {
	if bw.board == nil {
		return layout.Dimensions{}
	}

	if bw.board.Rows > 0 || bw.board.Cols > 0 {
		mW := int32(bw.board.Cols) * bw.cellSize
		mH := int32(bw.board.Rows) * bw.cellSize
		if mW > 0 && mH > 0 {
			cW := int32(mW / int32(bw.board.Cols))
			cH := int32(mH / int32(bw.board.Rows))
			cellSize := max(min(cW, cH), 16)
			bw.cellSize = cellSize
		}
	}

	boardWidth := int32(bw.board.Cols) * bw.cellSize
	boardHeight := int32(bw.board.Rows) * bw.cellSize

	for r := 0; r < bw.board.Rows; r++ {
		for c := 0; c < bw.board.Cols; c++ {
			bw.drawCell(gtx, r, c)
		}
	}

	bw.drawLine(gtx, image.Pt(0, 0), image.Pt(int(boardWidth), 0), color.NRGBA{R: 0, G: 0, B: 0, A: 255}, float32(bw.borderSize))
	bw.drawLine(gtx, image.Pt(0, 0), image.Pt(0, int(boardHeight)), color.NRGBA{R: 0, G: 0, B: 0, A: 255}, float32(bw.borderSize))
	bw.drawLine(gtx, image.Pt(int(boardWidth), 0), image.Pt(int(boardWidth), int(boardHeight)), color.NRGBA{R: 0, G: 0, B: 0, A: 255}, float32(bw.borderSize))
	bw.drawLine(gtx, image.Pt(0, int(boardHeight)), image.Pt(int(boardWidth), int(boardHeight)), color.NRGBA{R: 0, G: 0, B: 0, A: 255}, float32(bw.borderSize))

	return layout.Dimensions{Size: image.Pt(int(boardWidth), int(boardHeight))}
}

func (bw *BoardWidget) drawCell(gtx layout.Context, row, col int) {
	x := int32(col) * bw.cellSize
	y := int32(row) * bw.cellSize
	cellRect := image.Rect(
		int(x),
		int(y),
		int(x+bw.cellSize),
		int(y+bw.cellSize),
	)

	regionID := bw.board.Cells[row][col].RegionID
	cellColor := bw.regionColors[regionID]

	cellStack := clip.Rect(cellRect).Push(gtx.Ops)

	paint.ColorOp{Color: cellColor}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	if bw.board.Cells[row][col].HasQueen {
		bw.drawQueen(gtx, cellRect)
	}

	cellStack.Pop()

	bw.drawCellBorder(gtx, row, col, cellRect)

}

func (bw *BoardWidget) drawCellBorder(gtx layout.Context, row, col int, cellRect image.Rectangle) {
	borderBlack := color.NRGBA{R: 0, G: 0, B: 0, A: 255}

	top := row > 0 && bw.board.Cells[row-1][col].RegionID != bw.board.Cells[row][col].RegionID
	bottom := row < bw.board.Rows-1 && bw.board.Cells[row+1][col].RegionID != bw.board.Cells[row][col].RegionID
	left := col > 0 && bw.board.Cells[row][col-1].RegionID != bw.board.Cells[row][col].RegionID
	right := col < bw.board.Cols-1 && bw.board.Cells[row][col+1].RegionID != bw.board.Cells[row][col].RegionID

	if top {
		bw.drawLine(gtx, image.Pt(cellRect.Min.X, cellRect.Min.Y), image.Pt(cellRect.Max.X, cellRect.Min.Y), borderBlack, float32(bw.borderSize))
	} else {
		bw.drawLine(gtx, image.Pt(cellRect.Min.X, cellRect.Min.Y), image.Pt(cellRect.Max.X, cellRect.Min.Y), borderBlack, float32(bw.borderSize/2))
	}

	if bottom {
		bw.drawLine(gtx, image.Pt(cellRect.Min.X, cellRect.Max.Y), image.Pt(cellRect.Max.X, cellRect.Max.Y), borderBlack, float32(bw.borderSize))
	} else {
		bw.drawLine(gtx, image.Pt(cellRect.Min.X, cellRect.Max.Y), image.Pt(cellRect.Max.X, cellRect.Max.Y), borderBlack, float32(bw.borderSize/2))
	}

	if left {
		bw.drawLine(gtx, image.Pt(cellRect.Min.X, cellRect.Min.Y), image.Pt(cellRect.Min.X, cellRect.Max.Y), borderBlack, float32(bw.borderSize))
	} else {
		bw.drawLine(gtx, image.Pt(cellRect.Min.X, cellRect.Min.Y), image.Pt(cellRect.Min.X, cellRect.Max.Y), borderBlack, float32(bw.borderSize/2))
	}

	if right {
		bw.drawLine(gtx, image.Pt(cellRect.Max.X, cellRect.Min.Y), image.Pt(cellRect.Max.X, cellRect.Max.Y), borderBlack, float32(bw.borderSize))
	} else {
		bw.drawLine(gtx, image.Pt(cellRect.Max.X, cellRect.Min.Y), image.Pt(cellRect.Max.X, cellRect.Max.Y), borderBlack, float32(bw.borderSize/2))
	}

}

func (bw *BoardWidget) drawQueen(gtx layout.Context, cellRect image.Rectangle) {
	centerX := float32(cellRect.Min.X+cellRect.Max.X) / 2
	centerY := float32(cellRect.Min.Y+cellRect.Max.Y) / 2
	radius := float32(bw.cellSize) * 0.25

	queenColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	borderBlack := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	borderWhite := color.NRGBA{R: 255, G: 255, B: 255, A: 255}

	outerBorderCircle := clip.Ellipse{
		Min: image.Pt(
			int(centerX-radius-4),
			int(centerY-radius-4),
		),
		Max: image.Pt(
			int(centerX+radius+4),
			int(centerY+radius+4),
		),
	}

	outerStack := outerBorderCircle.Push(gtx.Ops)
	paint.ColorOp{Color: borderBlack}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	outerStack.Pop()

	innerBorderCircle := clip.Ellipse{
		Min: image.Pt(
			int(centerX-radius-2),
			int(centerY-radius-2),
		),
		Max: image.Pt(
			int(centerX+radius+2),
			int(centerY+radius+2),
		),
	}

	innerStack := innerBorderCircle.Push(gtx.Ops)
	paint.ColorOp{Color: borderWhite}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	innerStack.Pop()

	queenCircle := clip.Ellipse{
		Min: image.Pt(
			int(centerX-radius),
			int(centerY-radius),
		),
		Max: image.Pt(
			int(centerX+radius),
			int(centerY+radius),
		),
	}

	cellStack := queenCircle.Push(gtx.Ops)
	paint.ColorOp{Color: queenColor}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	cellStack.Pop()
}

func (bw *BoardWidget) drawLine(gtx layout.Context, p1, p2 image.Point, c color.NRGBA, width float32) {
	lineRect := image.Rect(
		int(float32(p1.X)-width/2),
		int(float32(p1.Y)-width/2),
		int(float32(p2.X)+width/2),
		int(float32(p2.Y)+width/2),
	)
	cellStack := clip.Rect(lineRect).Push(gtx.Ops)
	paint.ColorOp{Color: c}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	cellStack.Pop()
}

func GenerateRegionColors(alphabet []string) map[string]color.NRGBA {
	regionColors := make(map[string]color.NRGBA)
	hashSeed := 6769420
	for i, regionID := range alphabet {
		regionColors[regionID] = color.NRGBA{
			R: uint8((i * hashSeed) % 256),
			G: uint8((i * hashSeed / 256) % 256),
			B: uint8((i * hashSeed / 65536) % 256),
			A: 255,
		}
		hashSeed += hashSeed * i % 0xFFFFFF
	}
	return regionColors
}
