package agent

import (
	"fmt"
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/config"
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

const AGENT_NAME = "ali"

type ALiDNSAgent struct {
	agentNickName string
	dnsClient     *dns.Client
	domainName    string
	recordType    string
	RR            string
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

func (a *ALiDNSAgent) SetName(userNickName string) {
	a.agentNickName = userNickName
}

func (a *ALiDNSAgent) GetName() string {
	if a.agentNickName != "" {
		return fmt.Sprintf("[%s - %s]", AGENT_NAME, a.agentNickName)
	}
	return fmt.Sprintf("[%s]", AGENT_NAME)

}

func NewALiAgent(conf *config.DDNSConfig) *ALiDNSAgent {
	client, err := dns.NewClient(&openapi.Config{
		AccessKeyId:     &conf.AccessKey,
		AccessKeySecret: &conf.AccessKeySecret,
		RegionId:        &conf.RegionId,
	})
	if err != nil {
		log.Fatalf("[ERR] 初始化阿里云 Agent 错误: %v\n", err)
	}

	return &ALiDNSAgent{
		dnsClient:  client,
		domainName: conf.DomainName,
		recordType: conf.Type,
		RR:         conf.RR,
	}
}
