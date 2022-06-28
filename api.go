package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// 获取 DNS 记录列表
func GetDNSRecords(result *DNSRecordsResult) {
	url := GlobalConfig.Url.Base + GlobalConfig.Url.Get
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("version", GlobalConfig.Version)
	q.Add("type", GlobalConfig.Type)
	q.Add("key", GlobalConfig.Key)
	q.Add("domain", GlobalConfig.Domain)
	req.URL.RawQuery = q.Encode()
	log.Println("【查询】", req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(body, result)
	if err != nil {
		log.Fatal(err)
	}

	result.Reply.RecordListLog()
}

// 获取更新记录值时需要传的 host, 保持原 host 不变
func (rr *ResourceRecord) GetUpdateHost() {
	if rr.Host == GlobalConfig.Domain {
		rr.Host = ""
	} else {
		rr.Host = strings.Split(rr.Host, "."+GlobalConfig.Domain)[0]
	}
}

// 更新 DNS 记录
func (rr *ResourceRecord) UpdateDNSRecord() {
	url := GlobalConfig.Url.Base + GlobalConfig.Url.Update
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("version", GlobalConfig.Version)
	q.Add("type", GlobalConfig.Type)
	q.Add("key", GlobalConfig.Key)
	q.Add("domain", GlobalConfig.Domain)

	q.Add("rrid", rr.RecordId)
	q.Add("rrhost", rr.Host)
	q.Add("rrvalue", rr.Value)
	q.Add("rrttl", rr.Ttl)
	q.Add("rrdistance", rr.Distance)
	req.URL.RawQuery = q.Encode()
	log.Println("【更新】", req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := &DNSRecordsOperateResult{}
	err = xml.Unmarshal(body, result)
	if err != nil {
		log.Fatal(err)
	}

	result.RecordOperateLog(rr, "Update")
}

// 添加 DNS 记录
func (rr *ResourceRecord) AddDNSRecord() {
	url := GlobalConfig.Url.Base + GlobalConfig.Url.Add
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("version", GlobalConfig.Version)
	q.Add("type", GlobalConfig.Type)
	q.Add("key", GlobalConfig.Key)
	q.Add("domain", GlobalConfig.Domain)

	q.Add("rrtype", rr.Type)
	q.Add("rrhost", rr.Host)
	q.Add("rrvalue", rr.Value)
	q.Add("rrttl", rr.Ttl)
	q.Add("rrdistance", rr.Distance)
	req.URL.RawQuery = q.Encode()
	log.Println("【新增】", req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := &DNSRecordsOperateResult{}
	err = xml.Unmarshal(body, result)
	if err != nil {
		log.Fatal(err)
	}

	result.RecordOperateLog(rr, "Add")
}

// 删除 DNS 记录
func (rr *ResourceRecord) DeleteDNSRecord() {
	url := GlobalConfig.Url.Base + GlobalConfig.Url.Delete
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("version", GlobalConfig.Version)
	q.Add("type", GlobalConfig.Type)
	q.Add("key", GlobalConfig.Key)
	q.Add("domain", GlobalConfig.Domain)

	q.Add("rrid", rr.RecordId)

	req.URL.RawQuery = q.Encode()
	log.Println("【删除】", req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := &DNSRecordsOperateResult{}
	err = xml.Unmarshal(body, result)
	if err != nil {
		log.Fatal(err)
	}

	result.RecordOperateLog(rr, "Delete")
}

// 动态解析本地 IPv6 地址到 namesilo
func (rr *ResourceRecord) LocalPublicIPv6DDNS() {
	// ip 变化则更新记录值
	if isLatest := InList(IPs, rr.Value); isLatest {
		// 记录值已是最新 ip
		rr.RecordIsLatestLog()
	} else {
		fmt.Println(rr)
		// 更新记录值
		rr.Value = TargetIP
		rr.UpdateDNSRecord()
	}
}

// 配置中需要初始化添加的记录值
func InitAddDNSRecord() {
	for _, record := range GlobalConfig.Records {
		if InList(ResourceRecordHosts, record.Host) {
			log.Printf(`【配置记录值初始化】Host: "%s" 已存在于记录值列表，将自动忽略`, record.Host)
			continue
		}
		if len(record.Type) == 0 || len(record.Value) == 0 || len(record.Distance) == 0 {
			log.Fatalln("ConfigError: record 中 type, value, ttl 为必填项")
		}
		r := ResourceRecord{}
		r.Type = record.Type
		r.Host = record.Host
		r.Value = record.Value
		r.Ttl = record.Ttl
		r.Distance = record.Distance
		r.AddDNSRecord()
	}
}
