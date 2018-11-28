package main

import (
    "os"
    "fmt"
    "flag"
    "monitor_server/router"
    . "monitor_server/tools"
)

func init() {
    // 切换执行目录
    pwd, _ := os.Getwd();
    exeDir := flag.String("d", pwd, "execute directory")
    os.Chdir(*exeDir)

    // 初始化配置
    IniConf = LoadConfig()
}

func main() {
    host := IniConf.String("Server::Host")
    port := IniConf.String("Server::Port")
    r := router.GetGinDefault()
    r.Run(fmt.Sprintf("%s:%s", host, port))
}
