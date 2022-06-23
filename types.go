package main

import "encoding/xml"

// 记录列表
type ResourceRecord struct {
	RecordId string `xml:"record_id"` // 记录 id，更新记录和删除记录都会用到
	Type     string `xml:"type"`      // 记录值类型，值为 "A"、"AAAA"、"CNAME"、"MX"、"TXT"，"AAAA" 支持 IPv6
	Host     string `xml:"host"`      // 主机名，修改时不用包含域名
	Value    string `xml:"value"`     // 记录值，IPv4 或 IPv6 地址
	Ttl      string `xml:"ttl"`       // DNS记录在DNS服务器上的缓存时间，最小可设置为 3600 秒, 默认 7200 秒
	Distance string `xml:"distance"`  // 仅用于 MX 类型（如果未提供，则默认为 10），其他类型默认为 0
}

// API 响应状态
type ResponseStatus struct {
	Code   string `xml:"code"`   // 状态码，300 为成功
	Detail string `xml:"detail"` // 状态, success 为成功
}

type RecordReply struct {
	ResponseStatus
	ResourceRecords []ResourceRecord `xml:"resource_record"` // 记录列表
}

// 查询记录
type DNSRecordsResult struct {
	XMLName xml.Name    `xml:"namesilo"`
	Reply   RecordReply `xml:"reply"`
}

type OperatRecordReply struct {
	ResponseStatus
	RecordId string `xml:"record_id"` // 已更新资源记录的新唯一 ID。任何成功的更新都会产生一个新的record_id
}

// 增删改记录
type DNSRecordsOperateResult struct {
	XMLName xml.Name          `xml:"namesilo"`
	Reply   OperatRecordReply `xml:"reply"`
}

// 记录配置
type RecordsConfig struct {
	Type     string `yml:"type"`
	Host     string `yml:"host"`
	Value    string `yml:"value"`
	Ttl      string `yml:"ttl"`
	Distance string `yml:"distance"`
}

// 记录配置
type URLConfig struct {
	Base   string `yml:"base"`
	Get    string `yml:"get"`
	Add    string `yml:"add"`
	Update string `yml:"update"`
	Delete string `yml:"delete"`
}

// 全局配置
type Config struct {
	Url     URLConfig
	Version string `yml:"version"`
	Type    string `yml:"type"`
	Key     string `yml:"key"`
	Domain  string `yml:"domain"`
	Records []RecordsConfig
}
