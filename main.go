package main

import (
	crawler "./crawler/icorating"
	"./misc"
)

func main() {
	misc.InitLog()
	config := misc.ReadConfig("config.json")
	manager := crawler.GamebombCrawler{}
	err := manager.Init(config)
	if err != nil {
		misc.LogError(err)
	}
}
