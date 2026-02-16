package gui

import (
	"gioui.org/layout"
)

func (app *App) MainLayout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }),
		layout.Flexed(0.6, func(gtx layout.Context) layout.Dimensions {
			if app.boardWidget != nil {
				return app.boardWidget.Layout(gtx)
			}
			return layout.Dimensions{}
		}),
		layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions { return layout.Dimensions{} }),
	)
}
