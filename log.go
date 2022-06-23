package main

import (
	"log"
)

// DNS 记录值为最新 IP，无需操作
func (rr *ResourceRecord) RecordIsLatestLog() {
	log.Printf("【查询记录值是否为最新】Host: %s.%s, Value: %s, TTL: %s, 无需更新", rr.Host, GlobalConfig.Domain, rr.Value, rr.Ttl)
}

// DNS 记录操作失败
func (rs *ResponseStatus) RecordFailedLog() {
	if rs.Code != "300" || rs.Detail != "success" {
		log.Fatalf("%s %s DNS记录操作失败, 请重试", rs.Code, rs.Detail)
	}
}

// DNS 记录操作成功
func (dr *DNSRecordsOperateResult) RecordSuccessLog(ip string, operateType string) {
	log.Printf("【%s记录】latestIP: %s, RecordId: %s, 操作成功", operateType, ip, dr.Reply.RecordId)
}
