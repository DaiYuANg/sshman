package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func newGlobalMenu(window fyne.Window) *fyne.MainMenu {
	appMenu := fyne.NewMenu("APP")
	infoMenu := fyne.NewMenu("Help", fyne.NewMenuItem("INFO", func() {
		dialog.ShowInformation("Info", "i", window)
	}))
	mainMenu := fyne.NewMainMenu(appMenu, infoMenu)
	return mainMenu
}
