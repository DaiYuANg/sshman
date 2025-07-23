package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/daiyuang/sshman/core/model"
	"github.com/daiyuang/sshman/core/ssh"
	"github.com/daiyuang/sshman/desktop/pkg"
	"github.com/samber/lo"
	"log"
	"strconv"
	"time"
)

func newSide(manager *ssh.Manager) *fyne.Container {
	var sshList = manager.ListConnections()

	listWidget := widget.NewList(
		func() int { return len(sshList) },
		func() fyne.CanvasObject {
			return widget.NewLabel("item")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(sshList[i].Name + " (" + sshList[i].Host + ")")
		},
	)

	listWidget.OnSelected = func(i int) {
		conn := sshList[i]
		log.Printf("你选择了：%s@%s:%d", conn.Username, conn.Host, conn.Port)
		// TODO: 你可以在这里启动 terminal 连接
	}

	// 添加按钮行为：弹出 Modal 表单
	addButton := widget.NewButtonWithIcon("添加 SSH 连接", theme.Icon(theme.IconNameContentAdd), func() {
		nameEntry := widget.NewEntry()
		nameEntry.Size().Min(fyne.NewSize(400, 30))
		hostEntry := widget.NewEntry()
		portEntry := widget.NewEntry()
		usernameEntry := widget.NewEntry()
		passwordEntry := widget.NewPasswordEntry()
		formItems :=
			[]*widget.FormItem{
				widget.NewFormItem("名称", nameEntry),
				widget.NewFormItem("主机", hostEntry),
				widget.NewFormItem("端口", portEntry),
				widget.NewFormItem("用户名", usernameEntry),
				widget.NewFormItem("密码", passwordEntry),
			}
		w := pkg.Current()
		dialog.ShowForm("新建 SSH 连接", "添加", "取消", formItems, func(ok bool) {
			if ok {
				port, _ := strconv.Atoi(portEntry.Text)
				conn := model.SSHConnection{
					ID:         fmt.Sprintf("%d", time.Now().UnixNano()),
					Name:       nameEntry.Text,
					Host:       hostEntry.Text,
					Port:       port,
					Username:   usernameEntry.Text,
					Password:   passwordEntry.Text,
					AuthMethod: "password",
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				if !lo.Try1(func() error {
					return manager.Add(&conn)
				}) {
					dialog.ShowError(fmt.Errorf("添加失败"), w)
				}
				listWidget.Refresh()
			}
		}, w)
	})

	// 左侧栏容器：按钮 + List
	return container.NewVBox(
		addButton,
		widget.NewSeparator(),
		listWidget,
		layout.NewSpacer(),
	)
}
