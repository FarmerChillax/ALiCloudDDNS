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
	for counter <= (len(GetIpFuncs)+1)*3 {
		currentIp, err = currentNode.getIpFunc()
		if err == nil {
			return currentIp, nil
		}
		log.Printf("[INFO] 从 %v 节点获取本机 IP 出错: %v\n", currentNode, err)
		currentNode = currentNode.next
		counter++
	}

	log.Panic("获取本机 IP 出错，请检查网络")
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
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("[INFO] 从 %s 获取本机 IP 出错: %v", url, err)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("[INFO] 解析本机 IP 响应内容成二进制出错: %v", err)
	}

	respMap := make(map[string]string)
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return "", fmt.Errorf("[INFO] 解析 JSON 数据出错: %v", err)
	}
	return respMap["ip"], nil
}
