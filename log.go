package main

import (
	"log"
	"os"
)

// DNS 记录值为最新 IP，无需操作
func (rr *ResourceRecord) RecordIsLatestLog() {
	log.Printf("【查询记录值是否为最新】Host: %s.%s, Value: %s, TTL: %s, 无需更新", rr.Host, GlobalConfig.Domain, rr.Value, rr.Ttl)
}

// DNS 记录查询
func (rs *ResponseStatus) RecordListLog() {
	if rs.Code != "300" || rs.Detail != "success" {
		log.Printf("%s %s DNS 记录列表查询失败, 请重试", rs.Code, rs.Detail)
	} else {
		log.Printf("查询 DNS 记录列表成功")
	}
}

// DNS 记录操作
func (dr *DNSRecordsOperateResult) RecordOperateLog(rr *ResourceRecord, operateType string) {
	if dr.Reply.Code != "300" || dr.Reply.Detail != "success" {
		log.Println("DNS记录操作失败, 请重试")

		if operateType == "Update" {
			// 操作失败，写入 id，下次启动自动删除
			file, err := os.OpenFile("error.gnd4", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Println("打开 error 文件失败")
			}

			if _, err := file.Write([]byte(rr.RecordId + "\n")); err != nil {
				log.Println("写入 error 文件失败", rr.RecordId)
			}
		}

	} else {
		log.Printf("【%s记录】latestIP: %s, RecordId: %s, 操作成功", operateType, rr.Value, dr.Reply.RecordId)
	}
}
