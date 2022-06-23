package main

// func main() {
// 	for _, rr := range ResourceRecords {
// 		rr.LocalPublicIPv6DDNS()

// 		// 更新单个记录值
// 		if rr.Host == "scpc.games."+GlobalConfig.Domain {
// 			rr.Value = IPs[0]
// 			rr.Host = ""
// 			rr.UpdateDNSRecord()
// 		}

// 		// 删除单个记录
// 		if rr.Host == "bad."+GlobalConfig.Domain {
// 			rr.DeleteDNSRecord()
// 		}
// 	}

// 	// 添加记录值
// 	rr := &ResourceRecord{
// 		Type:  "AAAA",
// 		Host:  "Bad3",
// 		Value: "240e:3b7:3223:da50:881d:b582:8fae:1cbd",
// 		Ttl:   "3600",
// 	}
// 	rr.AddDNSRecord()
// }
