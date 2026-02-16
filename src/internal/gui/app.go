package gui

import (
	"queenables/src/internal/board"
	"queenables/src/internal/gui/solver"
	"queenables/src/internal/gui/widgets"
	"queenables/src/internal/io"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/ncruces/zenity"
)

type App struct {
	appState         *GUIState
	boardWidget      *widgets.BoardWidget
	inputWidget      *widgets.InputWidget
	controlWidget    *widgets.ControlWidget
	statisticsWidget *widgets.StatisticsWidget
	mainLayout       layout.Widget
}

func RunGUI(window *app.Window, appState *GUIState) error {
	th := material.NewTheme()
	var ops op.Ops

	queenablesApp := &App{
		appState:         appState,
		boardWidget:      widgets.CreateBoardWidget(appState.Board),
		inputWidget:      widgets.CreateInputWidget(),
		controlWidget:    widgets.CreateControlWidget(),
		statisticsWidget: widgets.CreateStatisticsWidget(),
	}

	for {
		e := window.Event()
		switch e := e.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if queenablesApp.inputWidget.SubmitBtn.Clicked(gtx) {
				editorText := queenablesApp.inputWidget.Editor.Text()
				rawBoardData, err := board.ParseRawBoard(editorText)
				if err == nil {
					newBoard, err := board.NewBoard(rawBoardData)
					if err == nil {
						queenablesApp.appState.Board = newBoard
						queenablesApp.boardWidget.SetBoard(newBoard)
					}
				}
			}

			if queenablesApp.inputWidget.LoadFileBtn.Clicked(gtx) {
				go func() {
					filename, err := zenity.SelectFile(
						zenity.Title("Load Board File"),
						zenity.FileFilters{{
							Name:     "Text Files",
							Patterns: []string{"*.txt"},
						}},
					)
					if err == nil && filename != "" {
						newBoard, err := io.ReadBoardFromFile(filename)
						if err == nil {
							queenablesApp.appState.Board = newBoard
							queenablesApp.boardWidget.SetBoard(newBoard)
							queenablesApp.inputWidget.InputPath = filename
							queenablesApp.inputWidget.Editor.SetText(newBoard.String())
							window.Invalidate()
						}
					}
				}()
			}

			if queenablesApp.controlWidget.SaveFileBtn.Clicked(gtx) && queenablesApp.appState.Board != nil {
				go func() {
					filename, err := zenity.SelectFileSave(
						zenity.Title("Save Board File"),
						zenity.FileFilters{{
							Name:     "Text Files",
							Patterns: []string{"*.txt"},
						}},
					)
					if err == nil && filename != "" {
						err := io.WriteBoardToFile(filename, queenablesApp.appState.Board)
						if err != nil {
							println("Error saving board:", err.Error())
						}
					}
				}()
			}

			if queenablesApp.controlWidget.SolveBtn.Clicked(gtx) && queenablesApp.appState.Board != nil {
				queenablesApp.appState.SolveState = SolveStateSolving
				solver.SolveBoardAsync(
					queenablesApp.appState.Board,
					queenablesApp.appState.SolveProgressChan,
					300*time.Millisecond,
				)
			}

			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Max.X = gtx.Dp(200)
					return queenablesApp.inputWidget.Layout(gtx, th)

				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if queenablesApp.boardWidget != nil {
							return queenablesApp.boardWidget.Layout(gtx)
						}
						return layout.Dimensions{}
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					solvingBoard := queenablesApp.appState.SolveState == SolveStateSolving
					solved := queenablesApp.appState.SolveState == SolveStateSolved
					return queenablesApp.controlWidget.Layout(gtx, th, solvingBoard, solved)
				}),
			)

			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}

func StartGUI() {
	appState := &GUIState{
		Board:             nil,
		SolveState:        SolveStateIdle,
		SolveProgressChan: make(chan solver.SolveProgress),
	}

	go func() {
		w := new(app.Window)
		w.Option(app.Title("Queenables Solver"))
		w.Option(app.MinSize(1280, 720))
		if err := RunGUI(w, appState); err != nil {
			panic(err)
		}
	}()
	app.Main()
}
