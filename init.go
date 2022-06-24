package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("ddns.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("日志文件初始化失败")
	}
	log.SetOutput(file)
	log.Print("开始动态解析")
}

func init() {
	if len(os.Args) == 3 && os.Args[1] == "-c" && len(os.Args[2]) > 0 {
		ConfigFile = os.Args[2]
	}
}

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
