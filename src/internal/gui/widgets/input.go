package widgets

import (
	"gioui.org/layout"
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

func (iw *InputWidget) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Editor(th, &iw.Editor, "Enter board data here...").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &iw.SubmitBtn, "Submit")
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: 10}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(material.Button(th, &iw.LoadFileBtn, "Input From File").Layout),
				layout.Rigid(layout.Spacer{Width: 10}.Layout),
				layout.Flexed(1, material.Body2(th, iw.InputPath).Layout),
			)
		}),
	)
}
