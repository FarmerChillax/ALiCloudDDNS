from urllib import request
import socket
import time
import json

# url_1 = 'http://www.3322.org/dyndns/getip'

class ListNode():
    def __init__(self, x):
        self.val = x
        self.next = None

# from farmer socket 
def farmer_ip():
    try:    
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect(('www.farmer233.top', 11451))
        data = s.recv(1024).decode('utf-8')
        data = json.loads(data)
        if data['status'] == 200:
            my_ip = data['ip']
        s.close()
        return my_ip
    except Exception as base:
        print(base)
        return 1

# from xiaotao server
def xiaotao_ip():
    try:
        myip_html = request.urlopen('http://ip.xiaotao233.top/')
        myip = myip_html.read()
        myip = json.loads(str(myip, encoding='utf-8'))['ip']
        return myip
    except Exception as base:
        print(base)
        return 1

# from network
def network_ip():
    try:
        myip_html = request.urlopen('http://www.3322.org/dyndns/getip')
        myip = myip_html.read()
        myip = str(myip, encoding='utf-8')
        return myip
    except Exception as base:
        print(base)
        return 1

# listnode
def add_function():
    function = [farmer_ip, xiaotao_ip, network_ip]
    head = ListNode(function[0])
    function_point = head
    for key in function[1:]:
        function_point.next = ListNode(key)
        function_point = function_point.next
    
    # 循环链表
    function_point.next = head
    return head

def main(function_point):
    n = 0
    while n != 114514:
        myip = function_point.val()
        if myip != 1:
            return myip
        else:
            function_point = function_point.next
        time.sleep(3)
        # try 12 times
        n = n + 1 if n < 12 else 114514
    return None

# API function
function_point = add_function()
print(type(function_point))
def get_ip():
    myip = main(function_point)
    print("[INFO] GetIP msg: Get IP successfully, localIP :" + myip)
    return myip

# print(get_ip())
# print(type(get_ip()))