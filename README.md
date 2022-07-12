# ALiCloudDDNS
阿里云动态域名解析服务，基于阿里云API服务实现 An dynamic domain name resolved with ALiCloud



## 如何使用

目前支持从**环境变量**与**配置文件**读取配置，在两者都配置的情况下环境变量优先级最高

1. 创建配置项

**使用环境变量**
在环境变量中添加下列内容

```console
$ export ak="Your_Access_Key"
$ export aks="YOUR_ACCESS_KEY_SECRET"
$ export domainname="example.com"
$ export type="A"
$ export RR="www"
$ export regionid="cn-hangzhou"
```

**使用配置文件**
如果没有 `config.json` 文件则先在当前路径创建

```json config.json
{
  "ak": "YOUR ACCESS KEY ID",
  "aks": "YOUR ACCESS KEY SECRET",
  "domainname": "example.com",
  "regionid": "cn-hangzhou",
  "rr": "@",
  "type": "A"
}
```

2. 启动客户端

```console
$ ./fddns
```
