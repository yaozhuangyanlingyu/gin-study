package main

import (
	"flag"
	"fmt"
	"monitor_server/setting"
	. "monitor_server/tools"
	"os"
)

func init() {
	// 切换执行目录
	pwd, _ := os.Getwd()
	exeDir := flag.String("d", pwd, "execute directory")
	os.Chdir(*exeDir)

	// 初始化配置
	IniConf = LoadConfig()
}

func main() {
	r := setting.GetGinDefault()
	host := IniConf.String("Server::Host")
	port := IniConf.String("Server::Port")
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
