package main

import (
	"fmt"
	"net"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pion/stun"
)

// 内置的STUN服务器列表
var stunServers = []string{
	// 国际 STUN 服务器
	"stun.l.google.com:19302",
	"stun1.l.google.com:19302",
	"stun2.l.google.com:19302",
	"stun3.l.google.com:19302",
	"stun4.l.google.com:19302",
	"stun.ekiga.net:3478",
	"stun.ideasip.com:3478",
	"stun.schlund.de:3478",
	"stun.stunprotocol.org:3478",
	"stun.voipbuster.com:3478",
	"stun.voipstunt.com:3478",
	"stun.services.mozilla.com:3478",

	// 中国大陆 STUN 服务器
	"stun.qq.com:3478",
	"stun.miwifi.com:3478",
	"stun.wide.net.cn:3478",
	"stun.chat.bilibili.com:3478",

	// 其他 STUN 服务器
	"turn.cloudflare.com:3478",
	"fwa.lifesizecloud.com:3478",
	"stun.isp.net.au:3478",
	"stun.freeswitch.org:3478",
	"stun.voip.blackberry.com:3478",
	"stun.nextcloud.com:3478",
	"stun.sipnet.com:3478",
	"stun.radiojar.com:3478",
	"stun.sonetel.com:3478",
	"stun.voipgate.com:3478",
	"stun.cloudflare.com:53",
	"stun.cloudflare.com:80",
}

// 检测NAT类型的函数
func checkNatType(stunServer string) (string, error) {
	conn, err := net.Dial("udp", stunServer)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	c, err := stun.NewClient(conn)
	if err != nil {
		return "", err
	}

	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	var mappedAddr stun.XORMappedAddress
	err = c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			err = res.Error
			return
		}
		err = mappedAddr.GetFrom(res.Message)
	})
	if err != nil {
		return "", err
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	if localAddr.IP.Equal(mappedAddr.IP) {
		return "Open Internet", nil
	}

	if localAddr.IP.To4() != nil && mappedAddr.IP.To4() != nil &&
		!localAddr.IP.Equal(mappedAddr.IP) {
		return "Full Cone NAT", nil
	}

	if strings.Contains(localAddr.String(), ":") && strings.Contains(mappedAddr.String(), ":") {
		return "Symmetric NAT", nil
	}

	return "Port Restricted Cone NAT", nil
}

func main() {
	a := app.New()
	w := a.NewWindow("NAT Type Checker")

	// 创建选择STUN服务器的下拉菜单
	stunServerSelect := widget.NewSelect(stunServers, nil)
	stunServerSelect.PlaceHolder = "Select a STUN server"

	// 创建输入框让用户自定义STUN服务器
	customStunEntry := widget.NewEntry()
	customStunEntry.SetPlaceHolder("Or enter a custom STUN server")

	// 创建输出标签来显示NAT类型
	output := widget.NewLabel("")

	// 创建按钮，当点击时检测NAT类型
	var checkButton *widget.Button
	checkButton = widget.NewButton("Check NAT Type", func() {
		checkButton.Disable()
		var stunServer string
		if customStunEntry.Text != "" && stunServerSelect.Selected != "" {
			output.SetText("Using custom STUN server.")
			stunServer = customStunEntry.Text
		} else if customStunEntry.Text != "" {
			output.SetText("Using custom STUN server.")
			stunServer = customStunEntry.Text
		} else if stunServerSelect.Selected != "" {
			output.SetText(fmt.Sprintf("Using selected STUN server: %s", stunServerSelect.Selected))
			stunServer = stunServerSelect.Selected
		} else {
			output.SetText("Please select or enter a STUN server.")
			checkButton.Enable() // 如果没有选择或输入服务器，立即启用按钮
			return
		}
		defer checkButton.Enable()

		natType, err := checkNatType(stunServer)
		if err != nil {
			output.SetText(fmt.Sprintf("Error: %v", err))
			return
		}
		output.SetText(fmt.Sprintf("NAT Type: %s", natType))
	})

	// 布局组件
	w.SetContent(
		container.NewVBox(
			widget.NewLabel("Choose a STUN Server or enter a custom one:"),
			stunServerSelect,
			customStunEntry,
			checkButton,
			output,
		),
	)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
