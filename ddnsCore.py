import time
import json
import os
import sys
from ddns_pack import domain_name_control, GetIP

# 1. 获取解析ip
# 2. 获取本机ip
# 3. 比较两者
#     - 不同更新解析
#     - 本机ip = 解析ip
#     - 相同进入阻塞
#     - 本机ip = 解析ip
# 4. 阻塞一段时间后，再次获取本机ip，判断新的本机ip与旧的是否相同。

# 日志
def write_log(My_ip):
    Time = time.strftime("%Y-%m-%d %H:%M:%S",time.localtime(time.time()))
    # 写入loacl ip 日志
    with open(path + '/local_ip.log', 'a', encoding='utf-8') as local_ip_file:
        local_ip_file.write(Time + '\n' + My_ip + '\n')
    print("[INFO] Writting log ...")

# main
def main(DomainName):
    cloud_ip, cloud_data = domain_name_control.get_cloud_ip(DomainName)
    print("my cloud ip :" + cloud_ip)
    while True:
        try:
            my_ip = GetIP.get_ip()
            if cloud_ip != my_ip:
                domain_name_control.updata_ip(cloud_data, my_ip)
                cloud_ip = my_ip
                write_log(my_ip)
        except Exception as base:
            print(base)
            time.sleep(10)
        
        print("Waiting for IP change...")
        time.sleep(6)

DomainName = 'farmer233.xyz'  # 需要动态解析的域名
path = os.getcwd()
main(DomainName)