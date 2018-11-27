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
    msgService := new(service.MsgService)
    r.GET("/sendmsg", msgService.SendMsg)

    return r
} 



















