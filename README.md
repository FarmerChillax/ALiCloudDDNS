# ALiCloudDDNS
阿里云动态域名解析服务，基于阿里云API服务实现 An dynamic domain name resolved with ALiCloud



## 如何使用

目前支持从**环境变量**与**配置文件**读取配置，在两者都配置的情况下环境变量优先级最高

1. 创建配置项

**使用环境变量**
在环境变量中添加下列内容

```console
$ export ACCESS_KEY="Your_Access_Key"
$ export ACCESS_KEY_SECRET="YOUR_ACCESS_KEY_SECRET"
$ export DOMAIN_NAME="example.com"
$ export TYPE="A"
$ export RR="www"
$ export REGION_ID="cn-hangzhou"
$ export NOTICE_URL=""
$ export SERVER_ADDR="gRPC 服务地址"
```

**使用配置文件**
如果没有 `config.json` 文件则先在当前路径创建

```json config.json
{
  "access_key": "YOUR ACCESS KEY ID",
  "access_key_secret": "YOUR ACCESS KEY SECRET",
  "domain_name": "example.com",
  "notice_url": "",
  "region_id": "cn-hangzhou",
  "rr": "@",
  "server_addr": "",
  "type": "A"
}
```

2. 启动客户端

```console
$ ./fddns
```
