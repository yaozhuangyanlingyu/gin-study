package main

import (
	"flag"
	"fmt"
	"monitor_server/setting"
	. "monitor_server/tools"
    "github.com/astaxie/beego/logs"
	"os"
)

func init() {
	// 切换执行目录
	pwd, _ := os.Getwd()
	exeDir := flag.String("d", pwd, "execute directory")
	os.Chdir(*exeDir)

	// 初始化配置
	IniConf = LoadConfig()

    // 日志配置
    logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s/app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":1,"color":true}`, "logs"))
}

func main() {
	r := setting.GetGinDefault()
	host := IniConf.String("Server::Host")
	port := IniConf.String("Server::Port")
    logs.Info("monitor is run, server: %s:%s", host, port)
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
