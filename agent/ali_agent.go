package agent

import (
	"fmt"
	"log"

	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	env "github.com/alibabacloud-go/darabonba-env/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"

	"github.com/alibabacloud-go/tea/tea"
)

type ALiDNSAgent struct {
	dnsClient  *dns.Client
	domainName string
	recordType string
	RR         string
	// 解析记录
	Record *dns.DescribeDomainRecordsResponseBodyDomainRecordsRecord
}

func (a *ALiDNSAgent) GetRecordIp() (string, error) {
	req := &dns.DescribeDomainRecordsRequest{
		DomainName: &a.domainName,
		Type:       &a.recordType,
	}
	// 请求阿里云记录
	log.Printf("请求阿里云解析记录中, 请求体: %v", req)
	resp, err := a.dnsClient.DescribeDomainRecords(req)
	if err != nil {
		return "", fmt.Errorf("请求阿里云解析记录出错, err: %v", err)
	}

	records := resp.Body.DomainRecords.Record
	for _, record := range records {
		if *record.RR == a.RR {
			a.Record = record
			return *record.Value, nil
		}
	}

	return "", fmt.Errorf("找不到 RR: %s; Type: %v 的记录", a.RR, a.recordType)
}

func (a *ALiDNSAgent) Update(ip string) (bool, error) {
	updateReq := &dns.UpdateDomainRecordRequest{
		RR:       a.Record.RR,
		RecordId: a.Record.RecordId,
		Type:     a.Record.Type,
		Value:    &ip,
	}
	_, err := a.dnsClient.UpdateDomainRecord(updateReq)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewALiAgent() *ALiDNSAgent {
	// fmt.Println(*env.GetEnv(tea.String("ACCESS_KEY_ID")),
	// 	*env.GetEnv(tea.String("ACCESS_KEY_SECRET")),
	// 	*env.GetEnv(tea.String("RegionId")))
	config := &openapi.Config{
		AccessKeyId:     env.GetEnv(tea.String("ACCESS_KEY_ID")),
		AccessKeySecret: env.GetEnv(tea.String("ACCESS_KEY_SECRET")),
		RegionId:        env.GetEnv(tea.String("RegionId")),
	}

	client, err := dns.NewClient(config)
	if err != nil {
		log.Fatalf("[ERR] 初始化阿里云 Agent 错误: %v\n", err)
	}

	return &ALiDNSAgent{
		dnsClient:  client,
		domainName: *env.GetEnv(tea.String("DomainName")),
		recordType: *env.GetEnv(tea.String("Type")),
		RR:         *env.GetEnv(tea.String("RR")),
	}
}
