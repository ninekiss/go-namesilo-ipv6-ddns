package main

func main() {
	for _, rr := range ResourceRecords {
		rr.GetUpdateHost()
		ResourceRecordHosts = append(ResourceRecordHosts, rr.Host)

		rr.LocalPublicIPv6DDNS()
	}

	InitAddDNSRecord()
}
