package service

import(
    "fmt"
    "github.com/gin-gonic/gin"
    . "monitor_server/tools"
)

// 发送消息控制器
type PushService struct{
    BaseService
}

/**
 * 发送微信消息
 */
func (this PushService) WxMsg(c *gin.Context) {
    fmt.Println("")
    msg := c.Query("msg")

    // 验证参数
    if len(msg) == 0 {
        ApiReturnError(c, 200, 1001, "msg param cannot empty")
        return
    }

    // 发送推送消息
    err := SendWxMsg(msg)
    if err != nil {
        ApiReturnError(c, 200, 1001, err.Error())
    }

    ApiReturnSuccess(c, 0, "success");
}









