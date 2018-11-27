package main

import (
    "os"
    "fmt"
    "flag"
    "monitor_server/router"
    "monitor_server/util"
)

func init() {
    fmt.Println("start Run")
    // 切换执行目录
    pwd, _ := os.Getwd();
    exeDir := flag.String("d", pwd, "execute directory")
    os.Chdir(*exeDir)

    // 初始化配置
    util.IniConf = util.LoadConfig()
}

func main() {
    host := util.IniConf.String("Server::Host")
    port := util.IniConf.String("Server::Port")
    r := router.GetGinDefault()
    r.Run(fmt.Sprintf("%s:%s", host, port))
}
