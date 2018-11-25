package service

import(
    "fmt"
    "github.com/gin-gonic/gin"
)

/**
 * 获取商品信息
 */
func GetProductInfo(c *gin.Context) {
    fmt.Println("党静姣想操你")

    c.JSON(200, gin.H{
        "message": "pong",
    })
}


























