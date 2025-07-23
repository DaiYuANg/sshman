package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"gui",
	fx.Provide(
		newApp,
		newWindow,
		newGlobalMenu,
		component(newSide, "sidebar"),
		component(newLayout, "layout"),
	),
	fx.Invoke(
		lifecycle,
	),
)

func newApp() fyne.App {
	a := app.New()
	return a
}

func lifecycle(window fyne.Window, menu *fyne.MainMenu) {
	window.SetMainMenu(menu)
	window.RequestFocus()
	window.ShowAndRun()
}
