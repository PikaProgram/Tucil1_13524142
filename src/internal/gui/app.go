package gui

import (
	"queenables/src/internal/board"
	"queenables/src/internal/gui/solver"
	"queenables/src/internal/gui/widgets"
	"queenables/src/internal/io"
	"sync"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/ncruces/zenity"
)

type App struct {
	appState         *GUIState
	boardWidget      *widgets.BoardWidget
	inputWidget      *widgets.InputWidget
	controlWidget    *widgets.ControlWidget
	statisticsWidget *widgets.StatisticsWidget
	boardKey         sync.RWMutex
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
		boardKey:         sync.RWMutex{},
	}

	for {
		e := window.Event()
		switch e := e.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if queenablesApp.inputWidget.SubmitBtn.Clicked(gtx) && queenablesApp.appState.Board == nil {
				editorText := queenablesApp.inputWidget.Editor.Text()
				rawBoardData, err := board.ParseRawBoard(editorText)
				if err == nil {
					newBoard, err := board.NewBoard(rawBoardData)
					if err == nil {
						queenablesApp.boardKey.Lock()
						queenablesApp.appState.Board = newBoard
						queenablesApp.boardWidget.SetBoard(newBoard)
						queenablesApp.boardKey.Unlock()
					}
				}
			}

			if queenablesApp.inputWidget.LoadFileBtn.Clicked(gtx) && queenablesApp.appState.Board == nil {
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

			if queenablesApp.controlWidget.SaveFileBtn.Clicked(gtx) && queenablesApp.appState.Board != nil && queenablesApp.appState.SolveState == solver.SolveStateSolved {
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

			if queenablesApp.controlWidget.SolveBtn.Clicked(gtx) && queenablesApp.appState.Board != nil && queenablesApp.appState.SolveState == solver.SolveStateIdle {
				queenablesApp.appState.SolveState = solver.SolveStateSolving
				queenablesApp.appState.SolveProgressChan = make(chan solver.SolveProgress)

				queenablesApp.boardKey.RLock()
				boardToSolve := queenablesApp.appState.Board
				queenablesApp.boardKey.RUnlock()

				newProgressChan := make(chan solver.SolveProgress)

				go func(progressChan chan solver.SolveProgress) {
					for progress := range progressChan {
						queenablesApp.boardKey.Lock()
						queenablesApp.appState.Board = progress.Board
						queenablesApp.boardWidget.SetBoard(progress.Board)
						queenablesApp.statisticsWidget.UpdateStats(progress.IterationCount, progress.ElapsedTime)

						if progress.IsComplete {
							if progress.Err != nil {
								queenablesApp.appState.SolveState = solver.SolveStateError
							} else {
								queenablesApp.appState.SolveState = solver.SolveStateSolved
							}
						} else {
							queenablesApp.appState.SolveState = solver.SolveStateSolving
						}
						queenablesApp.boardKey.Unlock()

						window.Invalidate()
					}
				}(newProgressChan)

				solver.SolveBoardAsync(
					boardToSolve,
					newProgressChan,
					67*time.Millisecond,
					0,
				)
			}

			// if queenablesApp.controlWidget.SolveBruteForceBtn.Clicked(gtx) && queenablesApp.appState.Board != nil && queenablesApp.appState.SolveState == solver.SolveStateIdle {
			// 	queenablesApp.appState.SolveState = solver.SolveStateSolving
			// 	queenablesApp.appState.SolveProgressChan = make(chan solver.SolveProgress)

			// 	queenablesApp.boardKey.RLock()
			// 	boardToSolve := queenablesApp.appState.Board
			// 	queenablesApp.boardKey.RUnlock()

			// 	newProgressChan := make(chan solver.SolveProgress)

			// 	go func(progressChan chan solver.SolveProgress) {
			// 		for progress := range progressChan {
			// 			queenablesApp.boardKey.Lock()
			// 			queenablesApp.appState.Board = progress.Board
			// 			queenablesApp.boardWidget.SetBoard(progress.Board)
			// 			queenablesApp.statisticsWidget.UpdateStats(progress.IterationCount, progress.ElapsedTime)

			// 			if progress.IsComplete {
			// 				if progress.Err != nil {
			// 					queenablesApp.appState.SolveState = solver.SolveStateError
			// 				} else {
			// 					queenablesApp.appState.SolveState = solver.SolveStateSolved
			// 				}
			// 			} else {
			// 				queenablesApp.appState.SolveState = solver.SolveStateSolving
			// 			}
			// 			queenablesApp.boardKey.Unlock()

			// 			window.Invalidate()
			// 		}
			// 	}(newProgressChan)

			// 	solver.SolveBoardAsync(
			// 		boardToSolve,
			// 		newProgressChan,
			// 		67*time.Millisecond,
			// 		1,
			// 	)
			// }

			if queenablesApp.controlWidget.ResetBtn.Clicked(gtx) && queenablesApp.appState.Board != nil && (queenablesApp.appState.SolveState == solver.SolveStateIdle || queenablesApp.appState.SolveState == solver.SolveStateSolved || queenablesApp.appState.SolveState == solver.SolveStateError) {
				queenablesApp.boardKey.Lock()
				queenablesApp.appState.Board = nil
				queenablesApp.boardWidget.SetBoard(nil)
				queenablesApp.appState.SolveState = solver.SolveStateIdle
				queenablesApp.boardKey.Unlock()
			}

			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Rigid(layout.Spacer{Width: 32}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					sidebarWidth := gtx.Dp(unit.Dp(240))
					gtx.Constraints.Min.X = sidebarWidth
					gtx.Constraints.Max.X = sidebarWidth

					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return queenablesApp.inputWidget.Layout(gtx, th, queenablesApp.appState.Board != nil)
						}),
						layout.Rigid(layout.Spacer{Height: 12}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return queenablesApp.controlWidget.Layout(gtx, th, queenablesApp.appState.SolveState, queenablesApp.appState.Board != nil)
						}),
						layout.Rigid(layout.Spacer{Height: 12}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if queenablesApp.appState.Board == nil {
								return layout.Dimensions{}
							}

							return queenablesApp.statisticsWidget.Layout(gtx, th, queenablesApp.appState.SolveState)
						}),
					)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if queenablesApp.appState.Board == nil {
							return layout.Dimensions{}
						}
						return queenablesApp.boardWidget.Layout(gtx)
					})
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
		SolveState:        solver.SolveStateIdle,
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
