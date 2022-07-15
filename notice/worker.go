package notice

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

type Notice struct {
	Url string
	// Method  string
	payload chan string
}

func New(url string) *Notice {
	return &Notice{
		Url:     url,
		payload: make(chan string),
	}
}

func (n *Notice) worker(ctx context.Context) {

	for requestBody := range n.payload {
		buffer := bytes.NewBufferString(requestBody)
		resp, err := http.Post(n.Url, "application/json", buffer)
		if err != nil {
			continue
			// log.Printf()
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		log.Printf("企微通知成功, resp: %s", respBody)
	}

}
