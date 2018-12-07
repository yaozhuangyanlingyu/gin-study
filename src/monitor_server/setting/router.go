package setting

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
	// gin.DisableConsoleColor()
    f, _ := os.Create(fmt.Sprintf("%s/%s.log", IniConf.String("Log::LogDir"), IniConf.String("Log::LogFile")))
    gin.DefaultWriter = io.MultiWriter(f)

    // 初始化组件
    components.RedisHandleObj = components.NewRedisHandle()
	r := gin.Default()

    // 发送报警接口
	pushService := new(service.PushService)
	r.GET("/push/wxmsg", pushService.WxMsg)

    // 监听sentry报警
	sentryService := new(service.SentryService)
	r.GET("/sentry/event/check", sentryService.EventCheck)

	return r
}

