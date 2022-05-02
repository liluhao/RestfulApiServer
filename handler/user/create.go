package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"github.com/zxmrlc/log/lager"
)

/*
创建用户逻辑：
1.从HTTP消息体获取参数（用户名和密码）
2.参数校验
3.加密密码
4.在数据库中添加数据记录
5.返回结果（这里是用户名）


*/
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)}) //lager是一个type Data map[string]interface{}
	var r CreateRequest
	//自动解析，并且失败的话会写一个400状态码
	if err := c.Bind(&r); err != nil {
		//即发起请求时既不传入用户名，也不出传入密码，返回ErrBind错误
		//Errno是实现了error接口的，所以可以传入，并且以指针行形式赋值给error接口
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//对密码进行加密
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// 插入这个用户到数据库
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	handler.SendResponse(c, nil, rsp)
}
