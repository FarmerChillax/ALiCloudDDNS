package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type GetIpFunc func() (string, error)

type GetIpClient struct {
	getIpFunc GetIpFunc
	next      *GetIpClient
}

var (
	GetIpFuncs  = []GetIpFunc{FromXiaoTao}
	once        sync.Once
	getIpClient *GetIpClient
)

func (g *GetIpClient) Get() (currentIp string, err error) {
	var counter int
	currentNode := g

	// 每个节点重复尝试获取 3 次
	for counter <= len(GetIpFuncs)+1 {
		currentIp, err = currentNode.getIpFunc()
		if err == nil {
			return currentIp, nil
		}
		log.Printf("获取本机 IP 出错: %v", err)
		currentNode = currentNode.next
		counter++
	}

	return "", fmt.Errorf("所获获取节点均无法获取本机 IP, 请检查网络连接")
}

func newGetIpClient() *GetIpClient {
	head := &GetIpClient{
		getIpFunc: FromXiaoTao,
	}
	point := head
	for _, item := range GetIpFuncs {
		point.next = &GetIpClient{getIpFunc: item}
		point = point.next
	}
	point.next = head
	return head
}

func NewGetIpClient() *GetIpClient {
	if getIpClient == nil {
		once.Do(func() {
			getIpClient = newGetIpClient()
		})
	}
	return getIpClient
}

func FromXiaoTao() (string, error) {
	url := "http://ip.xiaotao233.top/"
	log.Printf("正在从 %s 获取本机 IP", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	respMap := make(map[string]string)
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return "", err
	}
	return respMap["ip"], nil
}
