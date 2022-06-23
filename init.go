package main

import (
	"log"
)

func init() {
	// 读取全局配置
	parseConfig(&GlobalConfig)

	IPs = GetLocalPubilcIPv6()
	log.Println("【获取本地最新 IPv6 公网地址】", IPs)

	// 将本地公网 IPv6 地址中的第一个设置为要解析的目标 IP
	TargetIP = IPs[0]

	result := &DNSRecordsResult{}
	GetDNSRecords(result)
	ResourceRecords = result.Reply.ResourceRecords
}
