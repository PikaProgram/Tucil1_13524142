package board

import (
	"github.com/huandu/go-clone"
)

func CreateSolvedBoard(b *Board) (*Board, error) {
	solvedBoard := clone.Slowly(b).(*Board)
}
