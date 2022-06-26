// This file is auto-generated, don't edit it. Thanks.
package agent

import (
	"fmt"

	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	env "github.com/alibabacloud-go/darabonba-env/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
* Initialization  初始化公共请求参数
 */
func Initialization(regionId *string) (_result *dns.Client, _err error) {
	config := &openapi.Config{}
	// 您的AccessKey ID
	config.AccessKeyId = env.GetEnv(tea.String("ACCESS_KEY_ID"))
	// 您的AccessKey Secret
	config.AccessKeySecret = env.GetEnv(tea.String("ACCESS_KEY_SECRET"))
	// 您的可用区ID
	config.RegionId = regionId
	_result = &dns.Client{}
	_result, _err = dns.NewClient(config)
	return _result, _err
}

/**
 * 获取主域名的所有解析记录列表
 */
func DescribeDomainRecords(client *dns.Client, domainName *string, RR *string, recordType *string) (_result *dns.DescribeDomainRecordsResponse, _err error) {
	req := &dns.DescribeDomainRecordsRequest{}
	// 主域名
	req.DomainName = domainName
	// 主机记录
	req.RRKeyWord = RR
	// 解析记录类型
	req.Type = recordType
	_, tryErr := func() (_r *dns.DescribeDomainRecordsResponse, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.DescribeDomainRecords(req)
		if _err != nil {
			return _result, _err
		}

		console.Log(tea.String("-------------------获取主域名的所有解析记录列表--------------------"))
		console.Log(util.ToJSONString(tea.ToMap(resp)))
		_result = resp
		return _result, _err
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	_result = nil
	return _result, _err
}

/**
 * 修改解析记录
 */
func UpdateDomainRecord(client *dns.Client, req *dns.UpdateDomainRecordRequest) (_err error) {
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.UpdateDomainRecord(req)
		if _err != nil {
			return _err
		}

		console.Log(tea.String("-------------------修改解析记录--------------------"))
		console.Log(util.ToJSONString(tea.ToMap(resp)))

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		console.Log(error.Message)
	}
	return _err
}

func _main(args []*string) (_err error) {
	regionid := args[0]
	currentHostIP := args[1]
	domainName := args[2]
	RR := args[3]
	recordType := args[4]
	fmt.Println(regionid, currentHostIP, domainName, RR, recordType)
	client, _err := Initialization(regionid)
	if _err != nil {
		return _err
	}

	resp, _err := DescribeDomainRecords(client, domainName, RR, recordType)
	if _err != nil {
		return _err
	}

	fmt.Println(resp)
	fmt.Println("2333333333")
	if tea.BoolValue(util.IsUnset(tea.ToMap(resp))) || tea.BoolValue(util.IsUnset(tea.ToMap(resp.Body.DomainRecords.Record[0]))) {
		console.Log(tea.String("错误参数！"))
		return _err
	}

	record := resp.Body.DomainRecords.Record[0]
	// 记录ID
	recordId := record.RecordId
	// 记录值
	recordsValue := record.Value
	console.Log(tea.String("-------------------当前主机公网IP为：" + tea.StringValue(currentHostIP) + "--------------------"))
	if !tea.BoolValue(util.EqualString(currentHostIP, recordsValue)) {
		// 修改解析记录
		req := &dns.UpdateDomainRecordRequest{}
		// 主机记录
		req.RR = RR
		// 记录ID
		req.RecordId = recordId
		// 将主机记录值改为当前主机IP
		req.Value = currentHostIP
		// 解析记录类型
		req.Type = recordType
		_err = UpdateDomainRecord(client, req)
		if _err != nil {
			return _err
		}
	}

	return _err
}

// func main() {
// 	err := _main(tea.StringSlice(os.Args[1:]))
// 	if err != nil {
// 		panic(err)
// 	}
// }
