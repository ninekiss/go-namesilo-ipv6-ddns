# Go Namesilo DDNS
动态解析本地 IPv6 地址到 namesilo.com 托管的域名

### 特性
- namesilo.com DNS 记录新增（批量新增）、删除、修改、批量查，已实现
- 获取本地公共 IPv6 地址，已实现
- 动态 IPv6 地址解析，已实现
- 域名批量解析，已实现
- 定时任务自动域名解析， 未实现

### 配置
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
  -
    # 记录值类型，值为 "A"、"AAAA"、"CNAME"、"MX"、"TXT"，"AAAA" 支持 IPv6
    type: "AAAA"

    # 主机名，修改时不用包含域名
    host: "www" 

    # 记录值，IPv4 或 IPv6 地址
    value: "260e:3b7:3223:d50::1000"

    # DNS记录在DNS服务器上的缓存时间，最小可设置为 3600 秒, 默认 7200 秒
    ttl: "3600"

    # 仅用于 MX 类型（如果未提供，则默认为 10），其他类型默认为 0
    distance: "0"
  -
    type: "AAAA"
    host: ""
    value: "260e:3b7:3223:d50::1000"
    ttl: "3600"
    distance: "0"
```

### Example
动态解析本地 IPv6 地址到 namesilo，所有的解析记录使用同一个本地 IPv6 地址

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

根据配置文件 `config.yml` 中配置的 `records` 添加 DNS 解析记录，已存在则忽略
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
### API
#### Variables
  - var IPs [ ]string
  - var TargetIP string
  - var ResourceRecords [ ]ResourceRecord
#### func GetDNSRecords(result *DNSRecordsResult)
#### type ResourceRecord
  - #### func (rr *ResourceRecord) GetUpdateHost()
  - #### func (rr *ResourceRecord) UpdateDNSRecord()
  - #### func (rr *ResourceRecord) AddDNSRecord()
  - #### func (rr *ResourceRecord) DeleteDNSRecord()
  - #### func (rr *ResourceRecord) DeleteDNSRecord()