package main

import (
	"log"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (result *alidns20150109.Client, err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	result = &alidns20150109.Client{}
	result, err = alidns20150109.NewClient(config)
	return result, err
}

func main() {
	client, err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if err != nil {
		log.Println(err)
	}

	describeCustomLineRequest := &alidns20150109.DescribeCustomLineRequest{
		LineId: tea.Int64(1),
		Lang:   tea.String("test"),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, err = client.DescribeCustomLine(describeCustomLineRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println(err)
}
