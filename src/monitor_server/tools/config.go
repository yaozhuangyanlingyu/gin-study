package tools

import(
    "github.com/astaxie/beego/config"
)

// 配置变量
var IniConf config.Configer

/**
 * 初始化配置
 */
func LoadConfig() config.Configer {
    iniconf, err := config.NewConfig("ini", "config/local.conf")
    if err != nil {
        panic("file config/local.conf not exists.")
    }
    return iniconf
}

