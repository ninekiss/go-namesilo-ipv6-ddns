package main

import (
	"io/ioutil"
	"log"
	"net"

	"gopkg.in/yaml.v2"
)

func GetLocalPubilcIPv6() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	var ipv6s []string
	for _, address := range addrs {
		if ipnet, _ := address.(*net.IPNet); ISPublicIPv6(ipnet.IP) {
			ipv6s = append(ipv6s, ipnet.IP.String())
		}
	}
	return ipv6s
}

func ISIPv6(ip net.IP) bool {
	isIpv6 := ip.To4()
	return isIpv6 == nil
}

func ISPublicIPv6(ip net.IP) bool {
	return ISIPv6(ip) && ip.IsGlobalUnicast()
}

func InList(list []string, target string) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}

func parseConfig(config *Config) {
	file, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatal(`读取配置文件 "config.yml" 出错`)
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatal(`解析配置文件 "config.yml" 出错`)
	}

	if len(config.Key) == 0 {
		log.Fatal(`CONFIGERROR: key 不能为空`)
	}

	if len(config.Domain) == 0 {
		log.Fatal(`ConfigError: domain 不能为空`)
	}

	if len(config.Version) == 0 {
		config.Version = "1"
	}

	if len(config.Type) == 0 {
		config.Type = "xml"
	}

	if len(config.Url.Base) == 0 {
		config.Url.Base = "https://www.namesilo.com/api/"
	}

	if len(config.Url.Get) == 0 {
		config.Url.Get = "dnsListRecords"
	}

	if len(config.Url.Add) == 0 {
		config.Url.Add = "dnsAddRecord"
	}

	if len(config.Url.Update) == 0 {
		config.Url.Update = "dnsUpdateRecord"
	}

	if len(config.Url.Delete) == 0 {
		config.Url.Delete = "dnsDeleteRecord"
	}
}
