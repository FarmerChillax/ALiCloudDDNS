package notice

import (
	"bytes"
	"fmt"
	"net/http"
)

const WechatTemplate = `{
    "msgtype": "markdown",
    "markdown": {
        "content": "检测到 IP 地址发生变动\n
         > 更新变化: <font color=\"comment\">%s -> %s</font>
         > 更新状态: <font color=\"warning\"> %s </font>"
    }
}
`

const OnlineTemplate = `{
    "msgtype": "markdown",
    "markdown": {
        "content": "新客户端<font color=\"success\">上线</font>通知\n
         > 客户端 IP: <font color=\"comment\">%s</font>
		 > 上线时间: <font color=\"comment\">%s</font>"
    }
}
`

const OfflineTemplate = `{
    "msgtype": "markdown",
    "markdown": {
        "content": "客户端<font color=\"error\">连接出错</font>\n
         > 客户端 UUID: <font color=\"warning\">%s</font>
         > 客户端 IP: <font color=\"comment\">%s</font>
		 > 离线时间: <font color=\"comment\">%s</font>"
    }
}
`

type Notice struct {
	Url string
	// Method  string
}

func New(url string) *Notice {
	return &Notice{
		Url: url,
	}
}

func (n *Notice) Success() {

}

func (n *Notice) Error() {

}

func (n *Notice) Push(preIp, currentIp, status string) error {
	messageContent := fmt.Sprintf(WechatTemplate, preIp, currentIp, status)
	buffer := bytes.NewBufferString(messageContent)
	resp, err := http.Post(n.Url, "application/json", buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (n *Notice) Online(ip, time string) error {
	messageContent := fmt.Sprintf(OnlineTemplate, ip, time)
	buffer := bytes.NewBufferString(messageContent)
	resp, err := http.Post(n.Url, "application/json", buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (n *Notice) Offline(uuid, ip, time string) error {
	messageContent := fmt.Sprintf(OfflineTemplate, uuid, ip, time)
	buffer := bytes.NewBufferString(messageContent)
	resp, err := http.Post(n.Url, "application/json", buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
