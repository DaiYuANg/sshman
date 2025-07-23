package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/fx"
)

type LayoutParameter struct {
	fx.In
	Sidebar *fyne.Container `name:"sidebar"`
}

func newLayout(parameter LayoutParameter) *fyne.Container {
	content := newMainContent()
	split := container.NewHSplit(parameter.Sidebar, content)
	split.Offset = 0.2

	line := canvas.NewLine(theme.SeparatorColor())
	line.StrokeWidth = 1

	status := widget.NewLabel("状态栏：Ready")
	bottom := container.NewVBox(
		line,
		status,
	)
	return container.New(
		layout.NewBorderLayout(nil, bottom, nil, nil),
		bottom,
		split,
	)
}
