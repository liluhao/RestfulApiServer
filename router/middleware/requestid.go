package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

/*为了演示中间件的功能，这里给apiserver新增功能：
在请求和返回的Header中插入X-Request-Id(X-Request-Id值为32位的UUID,用于唯一标识次HTTP请求)*/

//
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")

		// 该中间件调用github.com/satori/go.uuid包生成个32位的UUID
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
