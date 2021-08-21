# ALiCloudDDNS
阿里云动态域名解析服务端，基于阿里云API服务实现

## 目前实现的功能
- [X] 基本的域名更新
- [X] 多个子域名更新

## 更多
1. 后期可能会写一个webAPP来管理，但本项目并不涉及。
2. 目前仅有这点需求，更多功能欢迎在issue提出
3. 快来提需求~🦆

## Installation
clone:

```shell
$ git clone https://gitee.com/Farmer-chong/ALiCloudDDNS.git # Gitee

```

创建&&安装虚拟环境:
with venv/virtualenv + pip:

```shell
$ python -m venv env  # use `virtualenv env` for Python2, use `python3 ...` for Python3 on Linux & macOS
$ source env/bin/activate  # use `env\Scripts\activate` on Windows
$ pip install -r requirements.txt
```

运行DDNS:
```shell
$ python3 ddnsCore.py # use `python ddnsCore.py` on windows
```

## 进程守护

1. 复制`ddns.service`文件到 `/etc/systemd/system/`目录下
2. 编辑`ddns.service`
   1. 修改里面的运行目录`WorkingDirectory`为你项目的根目录
   2. 修改里面的运行命令`ExecStart`为启动命令
3. 设置守护进程`systemctl enable ddns.service`
4. 启动守护进程`systemctl start ddns.service`

**注意⚠**如果python使用了虚拟环境，请使用虚拟环境中的python,如`/<env directory>/bin/python3`

ddns.service.bak
```shell
[Unit]
Description=My ALiDDNS Client

[Service]
Type=simple
WorkingDirectory=<your work directory> 
ExecStart= <Commands to run>  #  e.g python3 ddnsCore.py
KillMode=process
Restart=on-failure
RestartSec=3s

[Install]
WantedBy=multi-user.target
```
