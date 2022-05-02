package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	//func (c *Context) Get(key string) (value interface{}, exists bool) {}
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	//类型断言
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
