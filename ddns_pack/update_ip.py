# # 修改解析记录

# from aliyunsdkcore.client import AcsClient
# from aliyunsdkcore.acs_exception.exceptions import ClientException
# from aliyunsdkcore.acs_exception.exceptions import ServerException
# from aliyunsdkalidns.request.v20150109.UpdateDomainRecordRequest import UpdateDomainRecordRequest

# # 需要->RecordId，RR类型，Type，Value
# # 返回 RecordId，RequestId
# def update_ip(accessKeyId, accessSecret, new_cloud_date):
    
#     client = AcsClient(accessKeyId, accessSecret, 'cn-hangzhou')

#     request = UpdateDomainRecordRequest()
#     request.set_accept_format('json')

#     request.set_RecordId(new_cloud_date["RecordId"])
#     request.set_RR(new_cloud_date["RR"])
#     request.set_Type(new_cloud_date["Type"])
#     request.set_Value(new_cloud_date["Value"])
    
#     response = client.do_action_with_exception(request)
#     # python2:  print(response) 
#     print(str(response, encoding='utf-8'))
#     return response
