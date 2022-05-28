import time
import json
import os
from ddns_pack import domain_name_control, GetIP, settings

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
def main(DomainName, RR_list):
    # 获取解析的数据(json格式)， 仅返回用户指定的主机记录
    cloud_ip, cloud_data = domain_name_control.get_cloud_ip(DomainName, RR_list)
    print("my cloud ip :" + cloud_ip)
    while True:
        try:
            my_ip = GetIP.get_ip() # 通过循环链表获取本机ip地址
            # 判断解析的ip与local ip是否一致， 不一致就更新
            if cloud_ip != my_ip:
                for data in cloud_data:
                    if(data['RR'] in RR_list):
                        domain_name_control.updata_ip(data, my_ip)
                cloud_ip = my_ip # 更新本机ip
                # 写日志
                write_log(my_ip)
        except Exception as base:
            print(base)
            time.sleep(10)
        
        print("Waiting for IP change...")
        time.sleep(600)


# 读取数据





DomainName = settings.DomainName
RR_list = settings.RR_list
path = os.getcwd()

# run
if __name__ == "__main__":
    main(DomainName, RR_list)
    