package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fyne-io/terminal"
	"go.uber.org/fx"
)

var Module = fx.Module("gui", fx.Provide(newFyne), fx.Invoke(lifecycle))

func newFyne() fyne.Window {
	a := app.New()
	w := a.NewWindow("Hello")
	t := terminal.New()
	go func() {
		_ = t.RunLocalShell()
		a.Quit()
	}()
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		t,
	))
	return w
}

func lifecycle(window fyne.Window) {
	window.ShowAndRun()
}
