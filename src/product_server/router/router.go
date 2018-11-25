package router

import(
    "fmt"
    "github.com/gin-gonic/gin"
)

func init() {
    fmt.Println("日死你党静姣")
}

/**
 * 获取gin实例 
 */
func GetGinDefault() *gin.Engine {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })  
    }) 
    return r
} 



















