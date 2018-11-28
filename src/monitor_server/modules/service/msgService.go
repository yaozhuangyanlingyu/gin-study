package service

import(
    "fmt"
    "github.com/gin-gonic/gin"
)

// 消息控制器
type MsgService struct{
    BaseService
}

/**
 * 发送消息
 */
func (this MsgService) SendMsg(c *gin.Context) {
    fmt.Println("")
    msg := c.Query("msg")

    if len(msg) == 0 {
        c.JSON(200, gin.H{"success": false, "message":"msg param cannot empty!"})
    }

    c.JSON(200, gin.H{"success": true, "message":""})
}


























