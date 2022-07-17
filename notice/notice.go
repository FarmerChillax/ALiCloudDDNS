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

type Notice struct {
	Url string
	// Method  string
}

func New(url string) *Notice {
	return &Notice{
		Url: url,
	}
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

// func (n *Notice) worker(ctx context.Context) {

// 	for requestBody := range n.payload {
// 		buffer := bytes.NewBufferString(requestBody)
// 		resp, err := http.Post(n.Url, "application/json", buffer)
// 		if err != nil {
// 			continue
// 			// log.Printf()
// 		}
// 		defer resp.Body.Close()
// 		respBody, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			continue
// 		}
// 		log.Printf("企微通知成功, resp: %s", respBody)
// 	}

// }
