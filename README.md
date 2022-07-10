# Go Namesilo IPv6 DDNS

动态解析本地 IPv6 地址到 namesilo.com 托管的域名

## 特性

- namesilo.com DNS 记录新增（批量新增）、删除、修改、批量查，已实现
- 获取本地公共 IPv6 地址，已实现
- 动态 IPv6 地址解析，已实现
- 域名批量解析，已实现
- 定时任务自动域名解析， 未实现

## 安装

### 从源码编译

##### 1. 克隆源码

```bash
git clone https://github.com/lankeo/go-namesilo-ipv6-ddns.git
```

##### 2. 修改配置文件

详见配置

##### 3. 编译

```
cd go-namesilo-ipv6-ddns
go build .

# 编译时可指定配置文件名称, 如：这里使用 user-config.yml 作为运行程序时的配置文件
go build -ldflags "-X main.ConfigFile=user-config.yml" .

# 编译时自定义程序名（之后运行程序使用该名称），Windows下需要加 .exe 后缀
go build -o xxx .
go build -o xxx.exe .
```

### 下载二进制文件

#### 1. 下载

[ddns.rar](https://github.com/ninekiss/go-namesilo-ipv6-ddns/releases/download/v0.0.1/ddns.rar)

#### 2. 解压后修改配置文件 `config.ymal`，然后运行 `ddns.exe`

## 运行

##### Windows

- 直接运行 `go-namesilo-ipv6-ddns.exe`

- 从命令行运行

  ```bash
  go-namesilo-ipv6-ddns.exe

  # -c xxx.yml 可以指定运行时的配置文件
  go-namesilo-ipv6-ddns.exe -c xxx.yml
  ```

- 双击运行提供的 `ddns-run.bat`

##### Linux

- 从命令行运行

  ```bash
  ./go-namesilo-ipv6-ddns

  # -c xxx.yml 可以指定运行时的配置文件
  ./go-namesilo-ipv6-ddns -c xxx.yml
  ```

- 运行 `ddns-run` 脚本
  ```bash
  chomod +x ./ddns-run
  ./ddns-run
  ```

## 配置

程序根目录下 `config.yaml` 文件

```yaml
# namesilo API 请求 url
url:
  base: "https://www.namesilo.com/api/"
  get: "dnsListRecords"
  add: "dnsAddRecord"
  update: "dnsUpdateRecord"
  delete: "dnsDeleteRecord"

# namesilo API 要求字段
version: "1"
type: "xml"

# 账户设定的 key
key: "dd911ac49199bb691795a"

# 需要解析的域名（目前每次仅支持单个域名）
domain: "example.com"

# 需要添加的 DNS 解析记录（不添加请忽略）
records:
  - # 记录值类型，值为 "A"、"AAAA"、"CNAME"、"MX"、"TXT"，"AAAA" 支持 IPv6
    type: "AAAA"

    # 主机名，修改时不用包含域名
    host: "www"

    # 记录值，IPv4 或 IPv6 地址
    value: "260e:3b7:3223:d50::1000"

    # DNS记录在DNS服务器上的缓存时间，最小可设置为 3600 秒, 默认 7200 秒
    ttl: "3600"

    # 仅用于 MX 类型（如果未提供，则默认为 10），其他类型默认为 0
    distance: "0"
  - type: "AAAA"
    host: ""
    value: "260e:3b7:3223:d50::1000"
    ttl: "3600"
    distance: "0"
```

## Example

##### 1. 动态解析本地 IPv6 地址到 namesilo，所有的解析记录使用同一个本地 IPv6 地址

```go
  # `main.go` 文件
  package main

  func main() {
    for _, rr := range ResourceRecords {
      rr.GetUpdateHost()
      rr.LocalPublicIPv6DDNS()
    }
  }
```

- ResourceRecords 为所有记录值列表

##### 2. 根据配置文件 `config.yml` 中配置的 `records` 添加 DNS 解析记录，已存在则忽略

```go
  # `main.go` 文件
  package main

  func main() {
    for _, rr := range ResourceRecords {
      rr.GetUpdateHost()
    }

    InitAddDNSRecord()
  }
```

> Tips: 程序默认会动态解析本地 IPv6 地址到 namesilo，并根据配置添加 DNS 解析记录

## API

##### Variables

- var IPs [ ]string
- var TargetIP string
- var ResourceRecords [ ]ResourceRecord

##### func GetDNSRecords(result \*DNSRecordsResult)

##### type ResourceRecord

- ##### func (rr \*ResourceRecord) GetUpdateHost()
- ##### func (rr \*ResourceRecord) UpdateDNSRecord()
- ##### func (rr \*ResourceRecord) AddDNSRecord()
- ##### func (rr \*ResourceRecord) DeleteDNSRecord()
- ##### func (rr \*ResourceRecord) DeleteDNSRecord()
