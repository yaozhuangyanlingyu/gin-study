package router

import (
    "os"
    "io"
    "fmt"
    . "monitor_server/tools"
	"github.com/gin-gonic/gin"
	"monitor_server/modules/service"
	"monitor_server/components"
)

/**
 * 获取gin实例
 */
func GetGinDefault() *gin.Engine {
    // 运行环境
    gin.SetMode(gin.DebugMode)

    // 日志配置
	gin.DisableConsoleColor()
    f, _ := os.Create(fmt.Sprintf("%s/%s.log", IniConf.String("Log::LogDir"), IniConf.String("Log::LogFile")))
    gin.DefaultWriter = io.MultiWriter(f)

    // 初始化组件
    components.RedisHandleObj = components.NewRedisHandle()

	// 获取商品信息
	r := gin.Default()
	pushService := new(service.PushService)
	r.GET("/push/wxmsg", pushService.WxMsg)

	return r
}

