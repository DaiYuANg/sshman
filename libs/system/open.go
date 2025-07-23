package system

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Open 跨平台打开文件、目录、URL 或应用程序
func Open(target string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", target)
	case "windows":
		// start 是 cmd 内置命令，需要通过 cmd /c 调用
		// 需要避免 target 被当成窗口标题，故加空字符串 ""
		cmd = exec.Command("cmd", "/c", "start", "", target)
	case "linux":
		cmd = exec.Command("xdg-open", target)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Start()
}

func OpenDefaultTerminal() error {
	switch runtime.GOOS {
	case "windows":
		// 启动 cmd.exe 新窗口
		cmd := exec.Command("cmd", "/c", "start", "cmd")
		return cmd.Start()
	case "darwin":
		// macOS 用 open 命令打开 Terminal.app
		cmd := exec.Command("open", "-a", "Terminal")
		return cmd.Start()
	case "linux":
		// 尝试使用 x-terminal-emulator，常见发行版有这个软链接
		cmd := exec.Command("x-terminal-emulator")
		err := cmd.Start()
		if err != nil {
			// 如果上面失败，尝试 gnome-terminal 或者 xterm
			cmd = exec.Command("gnome-terminal")
			err = cmd.Start()
			if err != nil {
				cmd = exec.Command("xterm")
				err = cmd.Start()
			}
		}
		return err
	default:
		return fmt.Errorf("unsupported platform")
	}
}
