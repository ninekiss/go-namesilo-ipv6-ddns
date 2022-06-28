package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	for _, rr := range ResourceRecords {
		// 请求纪录列表后写入缓存
		CacheRecordValues = append(CacheRecordValues, rr.Value)

		rr.GetUpdateHost()
		ResourceRecordHosts = append(ResourceRecordHosts, rr.Host)
		rr.LocalPublicIPv6DDNS()
	}

	// InitAddDNSRecord()

	if CacheRecordValues != nil {
		err := ioutil.WriteFile("catch.gnd4", []byte(strings.Join(CacheRecordValues, ",")), 0644)
		if err != nil {
			log.Println("写入记录值缓存失败")
		}
	}
}
