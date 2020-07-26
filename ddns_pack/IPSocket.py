import socket
import json

def get_ip():
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect(('www.farmer233.top', 11451))
    data = s.recv(1024).decode('utf-8')
    data = json.loads(data)
    if data['status'] == 200:
        my_ip = data['ip']
    s.close()
    return my_ip

# print(get_ip())
# print(type(get_ip()))