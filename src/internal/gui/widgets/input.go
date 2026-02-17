package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type InputWidget struct {
	Editor       widget.Editor
	SubmitBtn    widget.Clickable
	LoadFileBtn  widget.Clickable
	InputPath    string
	RawBoardData string
}

func CreateInputWidget() *InputWidget {
	inputWidget := &InputWidget{
		InputPath:    "",
		RawBoardData: "",
		Editor:       widget.Editor{},
		SubmitBtn:    widget.Clickable{},
		LoadFileBtn:  widget.Clickable{},
	}

	inputWidget.Editor.SingleLine = false
	inputWidget.Editor.Submit = true

	return inputWidget
}

func (iw *InputWidget) Layout(gtx layout.Context, th *material.Theme, boardExists bool) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			h := gtx.Dp(unit.Dp(120))
			gtx.Constraints.Min.Y = h
			gtx.Constraints.Max.Y = h

			return material.Editor(th, &iw.Editor, "Enter board data here...").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &iw.SubmitBtn, "Submit")
			if boardExists {
				gtx = gtx.Disabled()
			}
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: 8}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &iw.LoadFileBtn, "Input From File")
			if boardExists {
				gtx = gtx.Disabled()
			}
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: 8}.Layout),
		layout.Rigid(material.Body2(th, iw.InputPath).Layout),
	)
}
