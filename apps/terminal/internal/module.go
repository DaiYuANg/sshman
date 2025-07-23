package internal

import (
	"fmt"
	"github.com/rivo/tview"
	"go.uber.org/fx"
)

var Module = fx.Module("tui", fx.Provide(app, layout), fx.Invoke(lifecycle))

func app() *tview.Application {
	app := tview.NewApplication()
	return app
}

func layout() *tview.Flex {
	// Left: connection list
	connList := tview.NewList().
		AddItem("New Connection", "test", 'n', func() {
			fmt.Print("tesst")
		}).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil)

	// Right: tab view
	sshTabs := tview.NewPages()
	sshTabs.AddPage("test", tview.NewBox(), true, true)
	// Bottom: SSH real-time output view (e.g., textview)
	sshOutput := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			// Refresh on new output
			//sshOutput.ScrollToEnd()
		})
	sshOutput.SetBorder(true).SetTitle("SSH Output")

	// Right-side layout (top = tab, bottom = ssh output)
	rightPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(sshTabs, 0, 3, false).
		AddItem(sshOutput, 0, 1, false)

	mainLayout := tview.NewFlex().
		AddItem(connList, 30, 1, true).
		AddItem(rightPanel, 0, 3, false)

	return mainLayout
}

func lifecycle(app *tview.Application, root *tview.Flex, lc fx.Lifecycle) {
	lc.Append(fx.StartHook(func() error {
		if err := app.SetRoot(root, true).Run(); err != nil {
			return err
		}
		return nil
	}),
	)
}
