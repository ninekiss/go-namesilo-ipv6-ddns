package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	// 读取记录值缓存
	content, err := ioutil.ReadFile("cache.gnd4")
	if err != nil && content != nil {
		log.Println("读取缓存失败，将重新获取")
	}
	for _, cacheValue := range strings.Split(string(content), ",") {
		if isLatest := InList(IPs, cacheValue); isLatest {
			log.Printf("【获取缓存记录值】记录值 %s 为最新, 无需更新\n", cacheValue)
		} else {
			log.Println("缓存记录值已失效, 重新获取", cacheValue)
			result := &DNSRecordsResult{}
			GetDNSRecords(result)
			ResourceRecords = result.Reply.ResourceRecords
			break
		}
	}
}

func init() {
	content, err := ioutil.ReadFile("error.gnd4")
	if err != nil && content != nil {
		log.Println("读取错误文件失败")
	}

	for _, errorRecordId := range strings.Split(string(content), "\n") {
		if len(errorRecordId) > 0 {
			rr := ResourceRecord{
				RecordId: errorRecordId,
			}
			rr.DeleteDNSRecord()
		}
	}
}
