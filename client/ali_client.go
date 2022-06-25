package client

type ALiDNSClient struct {
}

func (a *ALiDNSClient) GetIp(rrType string) string {
	return ""
}

func (a *ALiDNSClient) Update(ip string) (bool, error) {
	return true, nil
}
