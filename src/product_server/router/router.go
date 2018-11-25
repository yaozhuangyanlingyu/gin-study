package router

import(
    "github.com/gin-gonic/gin"
    "product_server/modules/service"
)

/**
 * 获取gin实例 
 */
func GetGinDefault() *gin.Engine {
    r := gin.Default()

    // 获取商品信息
    r.GET("/product/info", service.GetProductInfo)

    return r
} 



















