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
$ git clone https://github.com/greyli/sayhello.git # Github

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