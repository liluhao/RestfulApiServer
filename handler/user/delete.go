package user

import (
	"strconv"

	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary Delete an user by the user identifier
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	/*比如删除时，DELETE http://127.0.0.1/v1/user/1解析出id的值1，该id实际上就是数据库中的id索引，
	调用model.DeleteUser()函数删除*/
	userId, _ := strconv.Atoi(c.Param("id")) //将字符串转化为int型ID,
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
