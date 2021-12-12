package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "han-xuefeng/zookeeperAdmin"
	"han-xuefeng/zookeeperAdmin/infra"
)

func main() {
	// 入口
	file := kvs.GetCurrentFilePath("config.ini", 1)
	// 加载和解析配置
	conf := ini.NewIniFileConfigSource(file)
	app := infra.New(conf)
	app.Start()
}
