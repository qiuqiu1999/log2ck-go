package main

import (
	"github.com/qiuqiu1999/log2ck-go/src"
	"log"
)

var serverSetting *src.ServerSetting

func init() {
	initSetting()
}

func initSetting() {
	setting, err := src.NewSetting("config/config.yaml")
	if err != nil {
		log.Fatalf("init src.NewSetting err: %s", err)
	}
	serverSetting = &src.ServerSetting{}
	err = setting.ReadSection("ServerSetting", serverSetting)
	if err != nil {
		log.Fatalf("init setting.ReadSection err: %s", err)
	}
}

func main() {
	logManager := src.NewLogManager()
	logManager.Run()
}
