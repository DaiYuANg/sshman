package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fyne-io/terminal"
	"go.uber.org/fx"
	"log"
)

type WindowParameter struct {
	fx.In
	App    fyne.App
	Layout *fyne.Container `name:"layout"`
}

func newWindow(param WindowParameter) fyne.Window {
	w := param.App.NewWindow("SSHMAN")
	w.Resize(fyne.NewSize(1000, 600))
	w.SetContent(param.Layout)
	return w
}

func newMainContent() *fyne.Container {
	// 主内容
	t := terminal.New()
	go func() {
		_ = t.RunLocalShell()
		log.Printf("Terminal's shell exited with exit code: %d", t.ExitCode())
	}()

	terminalTab := container.NewTabItem("终端", container.NewStack(container.NewScroll(t)))

	welcomeTab := container.NewTabItem("欢迎", widget.NewLabel("欢迎使用 SSH 管理器"))
	settingsTab := container.NewTabItem("设置", widget.NewLabel("设置内容"))

	tabs := container.NewAppTabs(
		welcomeTab,
		terminalTab,
		settingsTab,
	)
	tabs.SetTabLocation(container.TabLocationTop)
	tabs.Select(terminalTab)

	content := container.NewStack(tabs)
	return content
}
