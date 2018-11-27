package service

import(
    "fmt"
    "github.com/gin-gonic/gin"
)

// 消息控制器
type MsgService struct{}

/**
 * 发送消息
 */
func (this MsgService) SendMsg(c *gin.Context) {
    fmt.Println("Hello")

    c.JSON(200, gin.H{
        "message": "pong",
    })
}


























