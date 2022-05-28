
from aliyunsdkcore.client import AcsClient
from aliyunsdkcore.acs_exception.exceptions import ClientException
from aliyunsdkcore.acs_exception.exceptions import ServerException
from aliyunsdkalidns.request.v20150109.DescribeDomainRecordsRequest import DescribeDomainRecordsRequest 
from aliyunsdkalidns.request.v20150109.UpdateDomainRecordRequest import UpdateDomainRecordRequest
import json
from .settings import accessKeyId, accessSecret


# 获取解析记录，返回字符串
def get_cloud_ip(DomainName, target_RR):
    client = AcsClient(accessKeyId, accessSecret, 'cn-hangzhou')
    request = DescribeDomainRecordsRequest()
    request.set_accept_format('json')
    request.set_DomainName(DomainName)
    response = client.do_action_with_exception(request)
    cloud_data = json.loads(response)['DomainRecords']['Record']
    new_cloud_date = []
    for Record in cloud_data:
        if Record['RR'] in target_RR:
            cloud_ip = Record['Value']
            new_cloud_date.append(Record)

    return cloud_ip, new_cloud_date

# 需要->RecordId，RR类型，Type，Value
# 返回 RecordId，RequestId
def updata_ip(new_cloud_date,my_ip):
    client = AcsClient(accessKeyId, accessSecret, 'cn-hangzhou')
    request = UpdateDomainRecordRequest()
    request.set_accept_format('json')
    request.set_RecordId(new_cloud_date["RecordId"])
    request.set_RR(new_cloud_date['RR'])
    request.set_Type(new_cloud_date["Type"])
    request.set_Value(str(my_ip))
    
    response = client.do_action_with_exception(request)

    print('Host-RR:' + new_cloud_date['RR'],"IP updata ip success!")
