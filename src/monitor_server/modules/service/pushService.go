package service

import (
	"github.com/gin-gonic/gin"
	"monitor_server/components"
	. "monitor_server/tools"
)

// 发送消息控制器
type PushService struct {
	BaseService
}

/**
 * 发送微信消息
 */
func (this PushService) WxMsg(c *gin.Context) {
	msg := c.Query("msg")
	toUser := c.Query("toUser")

	// 验证参数
	if len(msg) == 0 {
		ApiReturnError(c, 200, 1001, "msg param cannot empty")
		return
	}
	if len(toUser) == 0 {
		ApiReturnError(c, 200, 1001, "toUser param cannot empty")
		return
	}

	// 发送推送消息
	_, err := components.SendWxMsg(toUser, msg)
	if err != nil {
		ApiReturnError(c, 200, 1001, err.Error())
	}

	ApiReturnSuccess(c, 0, "success")
}
