package widgets

import (
	"fmt"
	"queenables/src/internal/gui/solver"
	"time"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type SolveState int

const (
	SolveStateSolving SolveState = iota
	SolveStateSolved
	SolveStateError
)

type StatisticsWidget struct {
	elapsedTime    time.Duration
	iterationCount int
}

func CreateStatisticsWidget() *StatisticsWidget {
	return &StatisticsWidget{
		elapsedTime:    0,
		iterationCount: 0,
	}
}

func (sw *StatisticsWidget) UpdateStats(iterations int, elapsed time.Duration) {
	sw.iterationCount = iterations
	sw.elapsedTime = elapsed
}

func (sw *StatisticsWidget) Layout(gtx layout.Context, th *material.Theme, status solver.SolveState) layout.Dimensions {
	speed := 0.0
	if sw.elapsedTime > 0 {
		speed = float64(sw.iterationCount) / sw.elapsedTime.Seconds()
	}

	statusMsg := ""
	switch status {
	case solver.SolveStateIdle:
		statusMsg = "Ready to Solve"
	case solver.SolveStateSolving:
		statusMsg = "Solving..."
	case solver.SolveStateSolved:
		statusMsg = "Solution Found!"
	case solver.SolveStateError:
		statusMsg = "No Solution Found"
	}

	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceAround}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Iteration Count").Layout),
				layout.Rigid(material.Body1(th, fmt.Sprintf("%d", sw.iterationCount)).Layout),
			)
		}),
		layout.Rigid(layout.Spacer{Height: 16}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Elapsed Time").Layout),
				layout.Rigid(material.Body1(th, sw.elapsedTime.Round(time.Millisecond).String()).Layout),
			)
		}),
		layout.Rigid(layout.Spacer{Height: 16}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Iteration Speed").Layout),
				layout.Rigid(material.Body1(th, fmt.Sprintf("%.2f it/s", speed)).Layout),
			)
		}),
		layout.Rigid(layout.Spacer{Height: 16}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Status").Layout),
				layout.Rigid(material.Body1(th, statusMsg).Layout),
			)
		}),
	)
}
