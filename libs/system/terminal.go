package system

import (
	"os"
	"os/exec"
	"runtime"
)

func detectTerminals() []string {
	var terminals []string
	switch runtime.GOOS {
	case "windows":
		candidates := []string{"cmd.exe", "powershell.exe", "wt.exe"}
		for _, c := range candidates {
			if path, err := exec.LookPath(c); err == nil {
				terminals = append(terminals, path)
			}
		}
	case "darwin":
		apps := []string{
			"/Applications/Utilities/Terminal.app",
			"/Applications/iTerm.app",
			"/Applications/Hyper.app",
			"/Applications/Alacritty.app",
		}
		for _, app := range apps {
			if exists(app) {
				terminals = append(terminals, app)
			}
		}
	case "linux":
		candidates := []string{"gnome-terminal", "konsole", "xfce4-terminal", "xterm", "alacritty", "tilix"}
		for _, c := range candidates {
			if path, err := exec.LookPath(c); err == nil {
				terminals = append(terminals, path)
			}
		}
	}
	return terminals
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
