package client

import "testing"

func TestGetIpFromeXiaotao(t *testing.T) {
	ip, err := FromXiaoTao()
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
	t.Logf("Current Host Ip: %s\n", ip)
}
