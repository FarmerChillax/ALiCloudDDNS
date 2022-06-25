package client

import "fmt"

type ALiDNSClient struct {
}

func (a *ALiDNSClient) GetIp(rrType ...string) string {
	fmt.Println("获取解析:", rrType)
	return "解析IP: 1.1.1.1"
}

func (a *ALiDNSClient) Update(ip string) (bool, error) {
	return true, nil
}
