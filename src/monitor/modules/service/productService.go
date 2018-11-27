package service

import(
    "fmt"
    "github.com/gin-gonic/gin"
)

// 商品信息控制器
type ProductService struct{}

/**
 * 获取商品信息
 */
func (this ProductService) GetProductInfo(c *gin.Context) {
    fmt.Println("Hello")

    c.JSON(200, gin.H{
        "message": "pong",
    })
}


























