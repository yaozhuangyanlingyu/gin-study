package router

import(
    "github.com/gin-gonic/gin"
    "monitor_server/modules/service"
)

/**
 * 获取gin实例 
 */
func GetGinDefault() *gin.Engine {
    r := gin.Default()

    // 获取商品信息
    pushService := new(service.PushService)
    r.GET("/push/wxmsg", pushService.WxMsg)

    return r
} 



















