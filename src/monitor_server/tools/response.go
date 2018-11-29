package tools

import (
	"github.com/gin-gonic/gin"
)

/**
 * 返回请求成功数据
 * @param code  int     // code码
 * @param msg   string  // 消息内容
 * @return json
 */
func ApiReturnSuccess(c *gin.Context, code int, msg string) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": msg,
	})
}

/**
 * 返回请求失败数据
 * @param httpCode  int     // http状态码
 * @param code      int     // code码
 * @param msg       string  // 消息内容
 * @return json
 */
func ApiReturnError(c *gin.Context, httpCode int, code int, msg string) {
	c.JSON(httpCode, gin.H{
		"code":    code,
		"message": msg,
	})
}
