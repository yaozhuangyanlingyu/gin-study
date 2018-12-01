package tools

import (
	"github.com/astaxie/beego/config"
)

// 配置变量
// -- 配置文件变量
var IniConf config.Configer
// -- 运行环境
var RunEnv string

/**
 * 初始化配置
 */
func LoadConfig() config.Configer {
    if IniConf != nil {
        return IniConf
    }
	iniconf, err := config.NewConfig("ini", "config/local.conf")
	if err != nil {
		panic("file config/local.conf not exists.")
	}
	return iniconf
}
