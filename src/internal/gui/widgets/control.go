package widgets

import (
	"queenables/src/internal/gui/solver"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ControlWidget struct {
	SolveBtn           widget.Clickable
	SolveBruteForceBtn widget.Clickable
	ResetBtn           widget.Clickable
	SaveFileBtn        widget.Clickable
	SaveImgBtn         widget.Clickable
}

func CreateControlWidget() *ControlWidget {
	return &ControlWidget{
		SolveBtn:           widget.Clickable{},
		SolveBruteForceBtn: widget.Clickable{},
		ResetBtn:           widget.Clickable{},
		SaveFileBtn:        widget.Clickable{},
		SaveImgBtn:         widget.Clickable{},
	}
}

func (cw *ControlWidget) Layout(gtx layout.Context, th *material.Theme, solveState solver.SolveState, boardExists bool) layout.Dimensions {
	solveTxt := "Solve Board"
	switch solveState {
	case solver.SolveStateSolving:
		solveTxt = "Solving..."
	case solver.SolveStateSolved:
		solveTxt = "Board Solved"
	}

	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &cw.SolveBtn, solveTxt)
			if solveState == solver.SolveStateIdle && !boardExists {
				gtx = gtx.Disabled()
			}
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: 8}.Layout),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &cw.ResetBtn, "Reset")
				if solveState == solver.SolveStateSolving || !boardExists {
					gtx = gtx.Disabled()
				}
				return btn.Layout(gtx)
			},
		),
		layout.Rigid(layout.Spacer{Height: 8}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &cw.SaveFileBtn, "Save Board")
			if solveState == solver.SolveStateSolving || solveState == solver.SolveStateIdle || solveState == solver.SolveStateError {
				gtx = gtx.Disabled()
			}
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: 8}.Layout),
	)
}
