package main

import (
    "os"
    "fmt"
    "flag"
    "monitor/router"
)

func init() {
    fmt.Println("start Run")
    // 切换执行目录
    pwd, _ := os.Getwd();
    exeDir := flag.String("d", pwd, "execute directory")
    os.Chdir(*exeDir)
}

func main() {
    r := router.GetGinDefault()
    r.Run(":8110")
}
