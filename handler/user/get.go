package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

/*比如GET http://127.0.0.1/v1/user/admin，
会解析出username的值是admin,然后调用model.GetUser()函数查询该用户的数据库记录并返回
*/
func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
