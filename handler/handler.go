package handler

import (
	"net/http"

	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//供所有的服务模块返回时调用，所以直接放在handler目录下
/*在返回结构体中，固定有Code和Message参数，这两个参数是通过函数DecodeErr()解析error类型的变量而来，并填充在Response结构体中。
Data域为interface{}类型，可以根据业务自己的需求来返▣，可以是map、int、string、struct、.array等Go语言变量类型。
*/
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	//返回JSON数据
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
