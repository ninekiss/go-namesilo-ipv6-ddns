package main

var GlobalConfig Config
var ConfigFile string = "config.yml"
var IPs []string
var TargetIP string
var ResourceRecords []ResourceRecord
var ResourceRecordHosts []string

var CacheRecordValues []string // 缓存记录值
