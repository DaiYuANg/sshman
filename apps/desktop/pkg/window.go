package pkg

import "fyne.io/fyne/v2"

func Current() fyne.Window {
	return fyne.CurrentApp().Driver().AllWindows()[0]
}
