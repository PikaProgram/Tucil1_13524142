package widgets

import (
	"fmt"
	"time"

	"gioui.org/layout"
	"gioui.org/widget/material"
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

func (sw *StatisticsWidget) Layout(gtx layout.Context, th *material.Theme, iterations int, elapsed time.Duration) layout.Dimensions {
	speed := 0.0
	if elapsed.Seconds() > 0 {
		speed = float64(iterations) / elapsed.Seconds()
	}

	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Iteration Count").Layout),
				layout.Rigid(material.Body1(th, fmt.Sprintf("%d", iterations)).Layout),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Iteration Count").Layout),
				layout.Rigid(material.Body1(th, fmt.Sprintf("%d", iterations)).Layout),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Elapsed Time").Layout),
				layout.Rigid(material.Body1(th, elapsed.Round(time.Millisecond).String()).Layout),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(material.Caption(th, "Iteration Speed").Layout),
				layout.Rigid(material.Body1(th, fmt.Sprintf("%.2f it/s", speed)).Layout),
			)
		}),
	)
}
