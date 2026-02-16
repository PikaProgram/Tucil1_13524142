package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ControlWidget struct {
	SolveBtn    widget.Clickable
	ResetBtn    widget.Clickable
	SaveFileBtn widget.Clickable
	SaveImgBtn  widget.Clickable
}

func CreateControlWidget() *ControlWidget {
	return &ControlWidget{
		SolveBtn:    widget.Clickable{},
		ResetBtn:    widget.Clickable{},
		SaveFileBtn: widget.Clickable{},
		SaveImgBtn:  widget.Clickable{},
	}
}

func (cw *ControlWidget) Layout(gtx layout.Context, th *material.Theme, solving bool, solved bool) layout.Dimensions {
	solveTxt := "Solve Board"
	if solving {
		solveTxt = "Solving..."
	}

	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &cw.SolveBtn, solveTxt)
			if solving {
				gtx = gtx.Disabled()
			}
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: 8}.Layout),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &cw.ResetBtn, "Reset")
				if solving {
					gtx = gtx.Disabled()
				}
				return btn.Layout(gtx)
			},
		),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &cw.SaveFileBtn, "Save Board")
			if solving || !solved {
				gtx = gtx.Disabled()
			}
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: 8}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &cw.SaveImgBtn, "Save Image")
			if solving || !solved {
				gtx = gtx.Disabled()
			}
			return btn.Layout(gtx)
		}),
	)
}
